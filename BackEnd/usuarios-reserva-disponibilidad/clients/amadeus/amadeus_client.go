package clients

import (
	e "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/Utils"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/model"
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
	err := Db.Where("hotelid = ?", hotelId).First(&amadeusMapping) // Chequear nombre del atributo hotelId en la tabla
	if err != nil {
		var errorModel model.AmadeusMapping
		return errorModel, e.NewBadRequestApiError("Error al obtener el mapping del hotel")
	}
	log.Debug("User: ", amadeusMapping)

	return amadeusMapping, nil
}
