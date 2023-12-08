package app

import (
	hotelController "busquedadehotel/controller/HotelSearch"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.POST("/hotels", hotelController.GetHotelsByDateAndCity)

	log.Info("Finishing mappings configurations")
}
