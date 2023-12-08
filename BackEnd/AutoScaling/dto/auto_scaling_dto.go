package dto

type EstadisticasDto struct {
	Id       string `json:"ID"`
	Name     string `json:"Name"`
	CPUPerc  string `json:"CPUPerc"`
	MemPerc  string `json:"MemPerc"`
	MemUsage string `json:"MemUsage"`
}

type EstadisticasDtos []EstadisticasDto
