package services_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	clients "urd/clients/reserva"
	"urd/model"
	services "urd/services"
)

type mockReservaClient struct {
	mock.Mock
}

func (m *mockReservaClient) GetReservaById(id int) model.Reserva {
	args := m.Called(id)
	return args.Get(0).(model.Reserva)
}

func (m *mockReservaClient) GetReservas() model.Reservas {
	args := m.Called()
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) InsertReserva(reserva model.Reserva) model.Reserva {
	args := m.Called(reserva)
	return args.Get(0).(model.Reserva)
}

func (m *mockReservaClient) GetReservasByUser(userId int) model.Reservas {
	args := m.Called(userId)
	return args.Get(0).(model.Reservas)
}

func TestGetReservaById(t *testing.T) {
	mockClient := new(mockReservaClient)

	expectedReserva := model.Reserva{
		ID:       1,
		FechaIn:  time.Now(),
		FechaOut: time.Now().Add(time.Hour * 24),
		UserId:   123,
		HotelId:  "Pele",
	}

	mockClient.On("GetReservaById", 1).Return(expectedReserva)

	clients.ReservaClient = mockClient

	reservaDto, err := services.ReservaService.GetReservaById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedReserva.ID, reservaDto.Id)
	assert.Equal(t, expectedReserva.FechaIn, reservaDto.FechaIngreso)
	assert.Equal(t, expectedReserva.FechaOut, reservaDto.FechaEgreso)
	assert.Equal(t, expectedReserva.UserId, reservaDto.UserId)
	assert.Equal(t, expectedReserva.HotelId, reservaDto.HotelId)
	mockClient.AssertExpectations(t)
}

func TestGetReservas(t *testing.T) {
	mockClient := new(mockReservaClient)

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  time.Now(),
			FechaOut: time.Now().Add(time.Hour * 24),
			UserId:   123,
			HotelId:  "Pele",
		},
		{
			ID:       2,
			FechaIn:  time.Now().Add(time.Hour * 24),
			FechaOut: time.Now().Add(time.Hour * 48),
			UserId:   789,
			HotelId:  "RonaldoGordo",
		},
	}

	mockClient.On("GetReservas").Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservas()

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)

}

func TestGetReservasByUser(t *testing.T) {
	mockClient := new(mockReservaClient)
	userId := 123

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  time.Now(),
			FechaOut: time.Now().Add(time.Hour * 24),
			UserId:   123,
			HotelId:  "Pele",
		},
		{
			ID:       2,
			FechaIn:  time.Now().Add(time.Hour * 24),
			FechaOut: time.Now().Add(time.Hour * 48),
			UserId:   123,
			HotelId:  "RonaldoGordo",
		},
	}

	mockClient.On("GetReservasByUser", userId).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByUser(userId)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)
}
