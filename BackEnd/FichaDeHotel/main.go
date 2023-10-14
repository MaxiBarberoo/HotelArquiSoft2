package main

import (
	app "HotelArquiSoft2/BackEnd/FichaDeHotel/app"
	"HotelArquiSoft2/BackEnd/FichaDeHotel/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
