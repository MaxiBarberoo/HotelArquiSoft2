package app

import (
	hotelController "fichadehotel/controller/hotel"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/hotels/:id", hotelController.GetHotelById)
	router.POST("/hotels", hotelController.HotelInsert)
	router.POST("/hotels/update", hotelController.UpdateHotel)

	log.Info("Finishing mappings configurations")
}
