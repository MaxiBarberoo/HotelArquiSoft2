package main

import (
	app "HotelArquiSoft2/BackEnd/BusquedaDeHotel/app"
	"HotelArquiSoft2/BackEnd/BusquedaDeHotel/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
