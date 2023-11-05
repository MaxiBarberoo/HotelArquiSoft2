package dto

type HotelDto struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Ciudad           string `json:"ciudad"`
	CantHabitaciones int    `json:"cantHabitaciones"`
	Desc             string `json:"descripcion"`
	Availability     bool   `json:"availability"`
}
type HotelsDto []HotelDto
