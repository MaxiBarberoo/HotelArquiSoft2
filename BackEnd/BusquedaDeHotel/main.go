package main

import (
	app "busquedadehotel/app"
	controller "busquedadehotel/controller/HotelSearch"
)

func main() {
	go app.StartRoute()

	go controller.Consumer()

  select{}
}
