package services

import (
	e "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/Utils"
	reservaClient "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/clients/reserva"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/dto"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/model"
)

type reservaService struct{}

type reservaServiceInterface interface {
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError)
	DeleteReserva(userId int) e.ApiError
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.ReservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("reserva not found")
	}

	reservaDto.FechaIngreso = reserva.FechaIn
	reservaDto.FechaEgreso = reserva.FechaOut
	reservaDto.HotelId = reserva.HotelId
	reservaDto.UserId = reserva.UserId
	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var searchDto dto.SearchDto

	searchDto.HotelId = reservaDto.HotelId
	searchDto.FechaIngreso = reservaDto.FechaIngreso
	searchDto.FechaEgreso = reservaDto.FechaEgreso

	available, err := AmadeusMappingService.CheckAvailability(searchDto)

	if available == false {
		var errorDto dto.ReservaDto
		return errorDto, e.NewBadRequestApiError("El hotel no tenia disponibilidad")
	}
	if err != nil {
		var errorDto dto.ReservaDto
		return errorDto, e.NewBadRequestApiError("Hubo un error al obtener la disponibilidad del hotel")
	}

	var reserva model.Reserva

	reserva.FechaIn = reservaDto.FechaIngreso
	reserva.FechaOut = reservaDto.FechaEgreso
	reserva.HotelId = reservaDto.HotelId
	reserva.UserId = reservaDto.UserId

	reserva = reservaClient.ReservaClient.InsertReserva(reserva)

	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByUser(userId)
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) DeleteReserva(reservaId int) e.ApiError {
	err := reservaClient.ReservaClient.DeleteReserva(reservaId)
	if err != nil {
		return e.NewBadRequestApiError("Error al eliminar la reserva")
	}

	return nil
}
