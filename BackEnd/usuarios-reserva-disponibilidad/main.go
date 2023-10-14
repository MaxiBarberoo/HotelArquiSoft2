package main

import (
	app "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/app"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
