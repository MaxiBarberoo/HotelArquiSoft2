package controller

import (
	"AutoScaling/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServicesAndStats(c *gin.Context) {
	dtoEstadisticas, err := service.AutoScalingService.GetServicesAndStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, dtoEstadisticas)
}

func GetStatsByService(c *gin.Context) {
	servicio := c.Param("servicio")

	dtoEstadisticas, err := service.AutoScalingService.GetStatsByService(servicio)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, dtoEstadisticas)
}

func ScaleService(c *gin.Context) {

	servicio := c.Param("servicio")

	numInstancias, err := service.AutoScalingService.ScaleService(servicio)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"servicio": servicio,
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
