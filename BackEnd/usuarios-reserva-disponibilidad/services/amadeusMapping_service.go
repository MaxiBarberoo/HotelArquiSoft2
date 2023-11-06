package services

import (
	e "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/Utils"
	cache "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/cache"
	amadeusMappingClient "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/clients/amadeus"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/dto"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	// Agrega otros campos del JSON si es necesario
}

type amadeusMappingService struct{}

type amadeusMappingServiceInterface interface {
	CreateMapping(amadeusMappingDto dto.AmadeusMappingDto) e.ApiError
	CheckAvailability(searchDto dto.SearchDto) (bool, e.ApiError)
	GetMappingByHotelId(hotelId string) (dto.AmadeusMappingDto, e.ApiError)
}

var (
	AmadeusMappingService amadeusMappingServiceInterface
)

func init() {
	AmadeusMappingService = &amadeusMappingService{}
}

func (s *amadeusMappingService) CreateMapping(amadeusMappingDto dto.AmadeusMappingDto) e.ApiError {

	var amadeusMapping model.AmadeusMapping

	amadeusMapping.HotelId = amadeusMappingDto.HotelId
	amadeusMapping.AmadeusHotelId = amadeusMappingDto.AmadeusHotelId

	err := amadeusMappingClient.CreateMapping(amadeusMapping)

	if err != nil {
		return err
	}

	return nil

}

func (s *amadeusMappingService) CheckAvailability(searchDto dto.SearchDto) (bool, e.ApiError) {

	cacheKey := fmt.Sprintf("availability:%s:%s:%s", searchDto.HotelId, searchDto.FechaIngreso, searchDto.FechaEgreso)

	// 2. Verificar el caché
	cachedValue := cache.Get(cacheKey)
	if cachedValue != nil {
		available, _ := strconv.ParseBool(string(cachedValue))
		return available, nil
	}

	amadeusMappingDto, errMapping := s.GetMappingByHotelId(searchDto.HotelId)
	if errMapping != nil {
		return false, e.NewBadRequestApiError("Error al obtener el hotel")
	}

	url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	data := "grant_type=client_credentials&client_id=[KEY]&client_secret=[SECRET]"
	//AGREGAR API KEY Y API SECRET EN CLIENT ID Y CLIENT SECRET RESPECTIVAMENTE

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return false, e.NewBadRequestApiError("Error al pedir el token")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error haciendo la solicitud:", err)
		return false, e.NewBadRequestApiError("Error al pedir el token")
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON
	var tokenResponse AccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println("Error decodificando la respuesta JSON:", err)
		return false, e.NewBadRequestApiError("Error al obtener el token")
	}

	// Acceder al access_token
	accessToken := tokenResponse.AccessToken

	url = fmt.Sprintf("https://test.api.amadeus.com/v3/shopping/hotel-offers?hotelIds=%s&checkInDate=%s&checkOutDate=%s", amadeusMappingDto.AmadeusHotelId, searchDto.FechaIngreso, searchDto.FechaEgreso)
	// Crear la solicitud HTTP

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return false, e.NewBadRequestApiError("Error al crear la solicitud de amadeus")
	}

	// Agregar el encabezado de autorización
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Realizar la solicitud
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error haciendo la solicitud:", err)
		return false, e.NewBadRequestApiError("Error al realizar la solicitud de amadeus")
	}
	defer resp.Body.Close()

	var responseMap map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		fmt.Println("Error decodificando la respuesta JSON:", err)
		return false, e.NewBadRequestApiError("Error decodificando la respuesta JSON de Amadeus")
	}

	// Acceder al atributo "available"
	if data, ok := responseMap["data"].([]interface{}); ok && len(data) > 0 {
		if offer, ok := data[0].(map[string]interface{}); ok {
			if available, ok := offer["available"].(bool); ok {
				fmt.Println("Disponibilidad del hotel:", available)

				cache.Set(cacheKey, []byte(strconv.FormatBool(available)))
				return available, nil

			}
		}
	}

	return false, e.NewBadRequestApiError("Hubo un problema al extraer el atributo disponibilidad de la respuesta de Amadeus")

}

func (s *amadeusMappingService) GetMappingByHotelId(hotelId string) (dto.AmadeusMappingDto, e.ApiError) {

	amadeusMapping, err := amadeusMappingClient.GetMappingByHotelId(hotelId)

	if err != nil {
		var errorDto dto.AmadeusMappingDto
		return errorDto, err
	}

	var amadeusMappingDto dto.AmadeusMappingDto

	amadeusMappingDto.HotelId = amadeusMapping.HotelId
	amadeusMappingDto.AmadeusHotelId = amadeusMapping.AmadeusHotelId

	return amadeusMappingDto, nil
}
