package services

import (
	e "HotelArquiSoft2/BackEnd/BusquedaDeHotel/Utils"
	hotelClient "HotelArquiSoft2/BackEnd/BusquedaDeHotel/clients/HotelSearch"
	"HotelArquiSoft2/BackEnd/BusquedaDeHotel/dto"
	"HotelArquiSoft2/BackEnd/BusquedaDeHotel/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel
	var err error
	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel.Descripcion = hotelDto.Desc
	hotel, err = hotelClient.InsertHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel could not be inserted")
	}
	hotelDto.Id = hotel.ID.Hex()

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
	hotel, err = hotelClient.UpdateHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}
	hotelDto.Id = hotel.ID.Hex()

	return hotelDto, nil
}