package app

import (
	hotelController "HotelArquiSoft2/BackEnd/BusquedaDeHotel/controller/HotelSearch"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.POST("/hotels", hotelController.GetHotelsByDateAndCity)

	log.Info("Finishing mappings configurations")
}
