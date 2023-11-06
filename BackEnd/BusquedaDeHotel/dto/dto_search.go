package dto

import "time"

type SearchDto struct {
	HotelId      string    `json:"hotel_id"`
	Ciudad       string    `json:"ciudad"`
	FechaIngreso time.Time `json:"fecha_ingreso"`
	FechaEgreso  time.Time `json:"fecha_egreso"`
}

type SearchDtos []SearchDto
