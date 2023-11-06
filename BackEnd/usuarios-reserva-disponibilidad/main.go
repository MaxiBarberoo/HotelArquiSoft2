package main

import (
	app "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/app"
	cache "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/cache"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
	cache.InitCache()
}
