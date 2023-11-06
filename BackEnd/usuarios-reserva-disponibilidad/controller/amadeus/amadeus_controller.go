package controller

import (
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/dto"
	service "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateMapping(c *gin.Context) {
	var amadeusMappingDto dto.AmadeusMappingDto

	err := c.BindJSON(&amadeusMappingDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = service.AmadeusMappingService.CreateMapping(amadeusMappingDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "Mapping Created Successfully",
	})

}

func CheckAvailability(c *gin.Context) {
	var searchDto dto.SearchDto

	err := c.BindJSON(&searchDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	availability, err2 := service.AmadeusMappingService.CheckAvailability(searchDto)

	if err2 != nil {
    fmt.Println(err2)
		c.JSON(http.StatusBadRequest, err2)
		return
	}

  fmt.Println(availability)
	c.JSON(http.StatusAccepted, gin.H{
		"availability": availability,
	})

}
