package main

import (
	app "busquedadehotel/app"
	service "busquedadehotel/services"
)

func main() {
	go app.StartRoute()

	go service.Consumer()

	select {}
}
