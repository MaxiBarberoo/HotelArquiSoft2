package service

import (
	"AutoScaling/dto"
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
	"strings"
	"time"
)
import e "AutoScaling/Utils"

type autoScalingService struct{}

type autoScalingServiceInterface interface {
	GetServicesAndStats() (dto.EstadisticasDtos, e.ApiError)
	GetStatsByService(servicio string) (dto.EstadisticasDtos, e.ApiError)
	ScaleService(servicio string) (int, e.ApiError)
	DeleteContainer(id string) e.ApiError
	getContainersIdsByService(servicio string) ([]string, e.ApiError)
	getContainerServiceById(id string) (string, e.ApiError)
	checkServiceExistenceAndScalability(servicio string) bool
	GetServiciosEscalables() []string
	AutoScaleContinuously(servicio string)
}

var (
	AutoScalingService autoScalingServiceInterface
)

func init() {
	AutoScalingService = &autoScalingService{}
}

var serviciosEscalables = []string{"hotel", "search", "user-reservation"}

func (a autoScalingService) GetServicesAndStats() (dto.EstadisticasDtos, e.ApiError) {
	var dtosEstadisticas dto.EstadisticasDtos

	command := exec.Command("docker", "stats", "--no-stream", "--format", "{{json .}}")

	stdout, err := command.StdoutPipe()
	if err != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(err.Error())
	}

	err = command.Start()
	if err != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(err.Error())
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var containerStats dto.EstadisticasDto

		err = json.Unmarshal(scanner.Bytes(), &containerStats)
		if err != nil {
			return dto.EstadisticasDtos{}, e.NewBadRequestApiError(err.Error())
		}

		dtosEstadisticas = append(dtosEstadisticas, containerStats)
	}

	err = scanner.Err()
	if err != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(err.Error())
	}

	err = command.Wait()
	if err != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(err.Error())
	}

	return dtosEstadisticas, nil
}

func (a autoScalingService) GetStatsByService(servicio string) (dto.EstadisticasDtos, e.ApiError) {
	if !a.checkServiceExistenceAndScalability(servicio) {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError("El servicio no existe")
	}

	var dtosEstadisticas dto.EstadisticasDtos

	containers, err := a.getContainersIdsByService(servicio)
	if err != nil {
		return dto.EstadisticasDtos{}, err
	}

	cmdArgs := append([]string{"stats", "--no-stream", "--format", "{{json .}}"}, containers...)

	command := exec.Command("docker", cmdArgs...)

	stdout, er := command.StdoutPipe()
	if er != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(er.Error())
	}

	er = command.Start()
	if er != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(er.Error())
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var containerStats dto.EstadisticasDto

		er = json.Unmarshal(scanner.Bytes(), &containerStats)
		if err != nil {
			return dto.EstadisticasDtos{}, e.NewBadRequestApiError(er.Error())
		}

		dtosEstadisticas = append(dtosEstadisticas, containerStats)
	}

	er = scanner.Err()
	if er != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(er.Error())
	}

	er = command.Wait()
	if er != nil {
		return dto.EstadisticasDtos{}, e.NewBadRequestApiError(er.Error())
	}

	return dtosEstadisticas, nil
}

func (a autoScalingService) ScaleService(servicio string) (int, e.ApiError) {
	if !a.checkServiceExistenceAndScalability(servicio) {
		return 0, e.NewBadRequestApiError("El servicio no existe o no es escalable")
	}

	ids, err := a.getContainersIdsByService(servicio)
	if err != nil {
		return 0, err
	}

	currQty := len(ids)

	scaleCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "up", "-d", "--scale", fmt.Sprintf("%s=%d", servicio, currQty+1))

	er := scaleCommand.Run()
	if er != nil {
		return 0, e.NewBadRequestApiError(er.Error())
	}

	restartCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "restart", fmt.Sprintf("%s%s", servicio, "nginx"))
	er = restartCommand.Run()
	if er != nil {
		return 0, e.NewBadRequestApiError(er.Error())
	}

	return currQty + 1, err
}

func (a autoScalingService) DeleteContainer(id string) e.ApiError {
	service, err := a.getContainerServiceById(id)
	if err != nil {
		return err
	}

	if !a.checkServiceExistenceAndScalability(service) {
		return e.NewBadRequestApiError("you cant delete this service's containers")
	}

	containers, err := a.getContainersIdsByService(service)
	if len(containers) == 1 {
		return e.NewBadRequestApiError("there has to be at least one container running for each service")
	}

	deleteCommand := exec.Command("docker", "rm", "-f", id)
	er := deleteCommand.Run()
	if er != nil {
		return e.NewBadRequestApiError(er.Error())
	}

	restartCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "restart", fmt.Sprintf("%s%s", service, "nginx"))
	er = restartCommand.Run()
	if er != nil {
		return e.NewBadRequestApiError(er.Error())
	}

	return nil
}

func (a autoScalingService) getContainersIdsByService(servicio string) ([]string, e.ApiError) {
	command := exec.Command("docker-compose", "-f", "../docker-compose.yml", "ps", "-q", servicio)
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return nil, e.NewBadRequestApiError(err.Error())
	}

	ids := strings.TrimSpace(string(output))

	idsArray := strings.Split(ids, "\n")

	return idsArray, nil
}

func (a autoScalingService) getContainerServiceById(id string) (string, e.ApiError) {
	command := exec.Command("docker", "inspect", "--format", `{ "service": "{{ index .Config.Labels "com.docker.compose.service" }}" }`, id)
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return "", e.NewBadRequestApiError("container not found")
	}

	var containerService struct {
		Service string `json:"Service"`
	}

	err = json.Unmarshal(output, &containerService)
	if err != nil {
		log.Error(err)
		return "", e.NewBadRequestApiError(err.Error())
	}

	return containerService.Service, nil
}

func (a autoScalingService) checkServiceExistenceAndScalability(servicio string) bool {
	command := exec.Command("docker-compose", "-f", "../docker-compose.yml", "config", "--services")
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return false
	}

	services := strings.TrimSpace(string(output))
	servicesArray := strings.Split(services, "\n")

	for _, serv := range servicesArray {

		if serv == servicio {
			return true
		}
	}

	for _, serv := range serviciosEscalables {
		if serv == servicio {
			return true
		}
	}
	return false
}

func (a autoScalingService) GetServiciosEscalables() []string {
	return serviciosEscalables
}

func (a autoScalingService) AutoScaleContinuously(servicio string) {
	log.Infof("Autoscaling %s", servicio)

	for {
		var avgCpuUsage float64

		stats, err := a.GetStatsByService(servicio)
		if err != nil {
			log.Errorf("Error getting %s stats: %v", servicio, err)
			continue
		}

		containersAmount := len(stats)

		for _, container := range stats {

			stringCPU := strings.Trim(container.CPUPerc, "%")
			intCPU, err := strconv.ParseFloat(stringCPU, 64)
			if err != nil {
				log.Errorf("Error parsing string: %v", err)
				continue
			}

			avgCpuUsage += intCPU
		}

		avgCpuUsage = avgCpuUsage / float64(containersAmount)

		if avgCpuUsage >= 60 || containersAmount < 2 {
			instances, err := a.ScaleService(servicio)
			if err != nil {
				log.Errorf("Error creating %s container: %s", servicio, err)
				continue
			}

			log.Infof("Scaling up %s to %d instances", servicio, instances)

		} else if avgCpuUsage < 20 && containersAmount > 2 {

			err = a.DeleteContainer(stats[containersAmount-1].Id)
			if err != nil {
				log.Errorf("Error deleting %s container: %s", servicio, err)
				continue
			}

			log.Infof("Scaling down %s to %d instances", servicio, containersAmount-1)
		}

		time.Sleep(20 * time.Second)
	}
}
