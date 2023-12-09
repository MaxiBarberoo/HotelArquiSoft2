package HotelSearch

import (
	"busquedadehotel/dto"
	service "busquedadehotel/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetHotelsByDateAndCity(c *gin.Context) {
	var hotelsDto dto.HotelsDto

	var searchDto dto.SearchDto

	err := c.BindJSON(&searchDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hotelsDto, err = service.HotelSearchService.GetHotelsByDateAndCity(searchDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, hotelsDto)
}
