package dto

type SolrDto struct {
	Id               []string `json:"id"`
	Name             []string `json:"name"`
	Ciudad           []string `json:"ciudad"`
	CantHabitaciones []int    `json:"cantHabitaciones"`
	Desc             []string `json:"descripcion"`
	Amenities        []string `json:"amenities"`
	Imagen           []string `json:"imagen"`
	Availability     []bool   `json:"availability"`
}
type SolrDtos []SolrDto
