package services

import (
	e "HotelArquiSoft2/BackEnd/BusquedaDeHotel/Utils"
	hotelSearchClient "HotelArquiSoft2/BackEnd/BusquedaDeHotel/clients/HotelSearch"
	"HotelArquiSoft2/BackEnd/BusquedaDeHotel/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type hotelSearchService struct{}

type hotelSearchServiceInterface interface {
	GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError)
	UpdateHotel(hotelId string) e.ApiError
}

var (
	HotelSearchService hotelSearchServiceInterface
)

func init() {
	HotelSearchService = &hotelSearchService{}
}

func (s *hotelSearchService) GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError) {

	var hotelsByCity, err = hotelSearchClient.GetHotelsByDateAndCity(searchDto)
	if err != nil {
		var errorHotel dto.HotelsDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}

	for i := 0; i < len(hotelsByCity); i++ {
		url := "http://localhost:8098/amadeus/availability"

		jsonData, err := json.Marshal(searchDto)
		if err != nil {
			return nil, e.NewBadRequestApiError("Error al convertir searchDto a JSON")
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, e.NewBadRequestApiError("Error al llamar al microservicio de disponibilidad de hotel")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, e.NewBadRequestApiError("Error con la respuesta obtenida")
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, e.NewBadRequestApiError("Error al extraer el json")
		}

		availability, ok := result["availability"].(bool)
		if !ok {
			return nil, e.NewBadRequestApiError("Availability no es una cadena o no existe en el JSON")
		}

		hotelsByCity[i].Availability = availability

	}

	var availableHotels dto.HotelsDto
	for i := 0; i < len(hotelsByCity); i++ {
		if hotelsByCity[i].Availability == true {
			availableHotels = append(availableHotels, hotelsByCity[i])
		}
	}

	return availableHotels, nil
}

func (s *hotelSearchService) UpdateHotel(hotelId string) e.ApiError {

	url := fmt.Sprintf("http://localhost:8090/hotels/%s", hotelId)

	resp, err := http.Get(url)
	if err != nil {
		return e.NewBadRequestApiError("Error al llamar al microservicio de ficha de hotel")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return e.NewBadRequestApiError("Error con la respuesta obtenida")
	}

	var hotel dto.HotelDto
	err = json.Unmarshal(body, &hotel)
	if err != nil {
		return e.NewBadRequestApiError("Error al extraer el json")
	}

	errSolr := hotelSearchClient.UpdateHotel(hotel)
	if err != nil {
		return errSolr
	}

	return nil
}
