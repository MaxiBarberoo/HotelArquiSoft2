package hotel

import (
	"fichadehotel/dto"
	service "fichadehotel/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id := c.Param("id")
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

func HotelInsert(c *gin.Context) {
	var hotelDto dto.HotelDto

	err := c.BindJSON(&hotelDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	service.SendToQueue(hotelDto.Id)
}

func UpdateHotel(c *gin.Context) {
	var hotelDto dto.HotelDto

	err := c.BindJSON(&hotelDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.UpdateHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	service.SendToQueue(hotelDto.Id)
}
