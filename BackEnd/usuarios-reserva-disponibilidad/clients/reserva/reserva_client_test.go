package clients_test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	clients "urd/clients/reserva"
	"urd/model"
)

func TestReservaClient(t *testing.T) {
	Db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}

	Db.AutoMigrate(&model.Reserva{})
	clients.Db = Db

	fechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	fechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	reserva := model.Reserva{
		FechaIn:  fechaI,
		FechaOut: fechaE,
		UserId:   1,
		HotelId:  "RonaldoGordo",
	}
	insertedReserva := clients.ReservaClient.InsertReserva(reserva)

	foundReserva := clients.ReservaClient.GetReservaById(insertedReserva.ID)
	assert.Equal(t, insertedReserva.ID, foundReserva.ID, "Las Reservas deben coincidir")

	reservas := clients.ReservaClient.GetReservas()
	assert.NotEmpty(t, reservas, "Debería haber al menos una Reserva en la base de datos")
	assert.Equal(t, reserva.UserId, reservas[0].UserId)
	assert.Equal(t, reserva.HotelId, reservas[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservas[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservas[0].FechaOut)

	reservasByUser := clients.ReservaClient.GetReservasByUser(reserva.UserId)
	assert.NotEmpty(t, reservasByUser, "Debería haber al menos una Reserva para el usuario especificado")
	assert.Equal(t, reserva.UserId, reservasByUser[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByUser[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByUser[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByUser[0].FechaOut)

	err = Db.Close()
	if err != nil {
		t.Fatalf("Error al cerrar la base de datos: %v", err)
	}
}
