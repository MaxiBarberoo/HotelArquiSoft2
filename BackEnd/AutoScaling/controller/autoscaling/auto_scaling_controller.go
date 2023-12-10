package controller

import (
	"AutoScaling/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type servicio struct {
	Servicio string `json:"servicio"`
}

func GetServicesAndStats(c *gin.Context) {
	dtoEstadisticas, err := service.AutoScalingService.GetServicesAndStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, dtoEstadisticas)
}

func GetStatsByService(c *gin.Context) {
	var serv servicio
	err := c.BindJSON(&serv)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dtoEstadisticas, err := service.AutoScalingService.GetStatsByService(serv.Servicio)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, dtoEstadisticas)
}

func ScaleService(c *gin.Context) {

	var serv servicio
	err := c.BindJSON(&serv)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numInstancias, err := service.AutoScalingService.ScaleService(serv.Servicio)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"servicio": serv.Servicio,
		"instancias": numInstancias,
	})
}

func DeleteContainer(c *gin.Context) {

	id := c.Param("id")

	err := service.AutoScalingService.DeleteContainer(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"contenedor_eliminado": id})
}

func GetScalableServices(c *gin.Context) {
	servicios := service.AutoScalingService.GetServiciosEscalables()
	c.JSON(http.StatusOK, servicios)
}
