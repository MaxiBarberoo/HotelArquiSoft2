package services

import (
	e "HotelArquiSoft2/BackEnd/FichaDeHotel/Utils"
	hotelClient "HotelArquiSoft2/BackEnd/FichaDeHotel/clients/hotel"
	"HotelArquiSoft2/BackEnd/FichaDeHotel/dto"
	"HotelArquiSoft2/BackEnd/FichaDeHotel/model"
	"bytes"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id string) (dto.HotelDto, e.ApiError) {

	var hotel, err = hotelClient.GetHotelById(id)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}
	var hotelDto dto.HotelDto

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDto.Name = hotel.Nombre
	hotelDto.CantHabitaciones = hotel.CantHab
	hotelDto.Id = hotel.ID.Hex()
	hotelDto.Desc = hotel.Descripcion
	hotelDto.Amenities = hotel.Amenities
	hotelDto.Ciudad = hotel.Ciudad

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel
	var err error
	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel.Descripcion = hotelDto.Desc
	hotel.Amenities = hotelDto.Amenities
	hotel.Ciudad = hotelDto.Ciudad

	hotel, err = hotelClient.InsertHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel could not be inserted")
	}
	hotelDto.Id = hotel.ID.Hex()

	var dtoAmadeusMapping dto.AmadeusMappingDto
	dtoAmadeusMapping.HotelId = hotelDto.Id
	dtoAmadeusMapping.AmadeusHotelId = hotelDto.AmadeusId

	url := "http://localhost:8098/amadeus/mapping"

	jsonData, err := json.Marshal(dtoAmadeusMapping)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error al convertir el json para llamar al servicio de mapeo ")
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error al llamar al servicio de mapeo de ids")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return hotelDto, e.NewBadRequestApiError("Error al llamar al servicio de mapeo de ids")
	}

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	var ID, err = primitive.ObjectIDFromHex(hotelDto.Id)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel.Descripcion = hotelDto.Desc
	hotel.ID = ID
	hotel.Amenities = hotelDto.Amenities
	hotel.Ciudad = hotelDto.Ciudad

	hotel, err = hotelClient.UpdateHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}
	hotelDto.Id = hotel.ID.Hex()

	return hotelDto, nil
}
