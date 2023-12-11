package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetReservaById(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservaDto := dto.ReservaDto{
		Id:           1,
		UserId:       1,
		HotelId:      1,
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
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:           2,
			UserId:       2,
			HotelId:      2,
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
			HotelId:      1,
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

func TestReservaInsert(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservaDto := dto.ReservaDto{

		UserId:       1,
		HotelId:      1,
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("InsertReserva", mockReservaDto).Return(mockReservaDto, nil)
	services.ReservaService = mockReservaService

	router.POST("/reservas", controllerReserva.ReservaInsert)

	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	var mockUserDto dto.UserDto
	mockUserDto.Id = 1
	mockUserDto.UserEmail = "mock@mock.com"
	mockUserDto.Tipo = 0
	mockUserDto.FirstName = "Mock"
	mockUserDto.LastName = "Mock"

	tokenString, err := jwtG.GenerateUserToken(mockUserDto)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", tokenString)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	mockReservaService.AssertCalled(t, "InsertReserva", mockReservaDto)
}
