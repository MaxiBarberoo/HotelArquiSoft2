package dto

type HotelDto struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	CantHabitaciones int    `json:"cantHabitaciones"`
	Desc             string `json:"descripcion"`
}
type HotelsDto []HotelDto
