package main

import (
	"AutoScaling/app"
	"AutoScaling/service"
)

func main() {
	servicios := service.AutoScalingService.GetServiciosEscalables()

	for _, servicio := range servicios {
		go service.AutoScalingService.AutoScaleContinuously(servicio)
	}

	app.StartRoute()

}
