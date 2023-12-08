package clients

import (
	e "urd/Utils"
	"urd/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type reservaClient struct{}

type ReservaClientInterface interface {
	GetReservaById(id int) model.Reserva
	GetReservas() model.Reservas
	InsertReserva(reserva model.Reserva) model.Reserva
	GetReservasByUser(userId int) model.Reservas
	DeleteReserva(userId int) e.ApiError
}

var (
	ReservaClient ReservaClientInterface
)

func init() {
	ReservaClient = &reservaClient{}
}
func (c *reservaClient) GetReservaById(id int) model.Reserva {
	var reserva model.Reserva

	Db.Where("id = ?", id).First(&reserva)
	log.Debug("Reserva: ", reserva)
	return reserva
}

func (c *reservaClient) GetReservas() model.Reservas {
	var reservas model.Reservas
	Db.Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func (c *reservaClient) InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)
	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva Created: ", reserva.ID)

	return reserva
}

func (c *reservaClient) GetReservasByUser(userId int) model.Reservas {
	var reservas model.Reservas
	Db.Where("user_id = ?", userId).Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func (c *reservaClient) DeleteReserva(reservaId int) e.ApiError {
	result := Db.Table("reservas").Where("reserva_id = ?", reservaId).Delete(nil)
	if result.Error != nil {
		return e.NewBadRequestApiError("Error al eliminar la reserva")
	}
	return nil
}
