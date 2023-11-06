package main

import (
	app "HotelArquiSoft2/BackEnd/BusquedaDeHotel/app"
	controller "HotelArquiSoft2/BackEnd/BusquedaDeHotel/controller/HotelSearch"
)

func main() {
	go app.StartRoute()

	go controller.Consumer()

  select{}
}
