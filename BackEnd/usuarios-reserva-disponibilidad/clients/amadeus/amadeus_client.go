package clients

import (
	e "urd/Utils"
	"urd/model"


	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func CreateMapping(amadeusMapping model.AmadeusMapping) e.ApiError {

	result := Db.Create(&amadeusMapping)

	if result.Error != nil {
		log.Error("")
		return e.NewBadRequestApiError("Error al insertar el mapeo en la BD")
	}

	log.Debug("User Created: ", amadeusMapping.ID)

	return nil
}
func GetMappingByHotelId(hotelId string) (model.AmadeusMapping, e.ApiError) {
	var amadeusMapping model.AmadeusMapping
	result := Db.Where("hotel_id = ?", hotelId).Find(&amadeusMapping) // Chequear nombre del atributo hotelId en la tabla
	if result.Error != nil {
		var errorModel model.AmadeusMapping
		return errorModel, e.NewBadRequestApiError("Error al obtener el mapping del hotel")
	}
	log.Debug("User: ", amadeusMapping)

	return amadeusMapping, nil
}
