package dto

type HotelDto struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	CantHabitaciones int      `json:"cantHabitaciones"`
	Ciudad           string   `json:"ciudad"`
	Desc             string   `json:"descripcion"`
	Amenities        []string `json:"amenities"`
	AmadeusId        string   `json:"amadeus_id"`
}
type HotelsDto []HotelDto
