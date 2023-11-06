package dto

type AmadeusMappingDto struct {
	Id              int       `json:"id"`
	HotelId         string    `json:"hotel_id,omitempty"`
	AmadeusHotelId  string    `json:"amadeus_hotel_id,omitempty"`
}

type AmadeusMappingsDto []AmadeusMappingDto
