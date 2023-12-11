package controller_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	e "urd/Utils"
	controllerReserva "urd/controller/reserva"
	"urd/dto"
	services "urd/services"
)

func TestGetReservaById(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservaDto := dto.ReservaDto{
		Id:           1,
		UserId:       1,
		HotelId:      "Pele",
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("GetReservaById", 1).Return(mockReservaDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas/:id", controllerReserva.GetReservaById)

	req, err := http.NewRequest("GET", "/reservas/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservaDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservaDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservaById", 1)
}

func TestGetReservas(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasDto := dto.ReservasDto{
		{Id: 1,
			UserId:       1,
			HotelId:      "Pele",
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:           2,
			UserId:       2,
			HotelId:      "RonaldoGordo",
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaService.On("GetReservas").Return(mockReservasDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas", controllerReserva.GetReservas)

	req, err := http.NewRequest(http.MethodGet, "/reservas", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservas")
}

func TestGetReservasByUser(t *testing.T) {
	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasByUserDto := dto.ReservasDto{
		dto.ReservaDto{
			Id:           1,
			UserId:       1,
			HotelId:      "Pele",
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaService.On("GetReservasByUser", 1).Return(mockReservasByUserDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas/:user_id", controllerReserva.GetReservasByUser)

	req, err := http.NewRequest(http.MethodGet, "/reservas/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasByUserDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservasByUser", 1)
}

type MockReservaService struct {
	mock.Mock
}

func (m *MockReservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.ReservaDto), nil
}

func (m *MockReservaService) GetReservas() (dto.ReservasDto, e.ApiError) {
	args := m.Called()
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservaDto), nil
}

func (m *MockReservaService) GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError) {
	args := m.Called(userId)
	return args.Get(0).(dto.ReservasDto), nil
}
