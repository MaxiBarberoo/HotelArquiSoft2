package app

import (
	amadeusController "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/controller/amadeus"
	reservaController "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/controller/reserva"
	userController "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/controller/user"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/reservas", reservaController.GetReservas)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/email", userController.GetUserByEmail)
	router.GET("/users/:id", userController.GetUserById)
	router.POST("/users/auth", userController.UserAuth)
	router.POST("/users", userController.UserInsert)
	router.POST("/reservas", reservaController.ReservaInsert)
	router.POST("/reservas/:id", reservaController.GetReservaById)
	router.GET("/reservas/reservauser/:user_id", reservaController.GetReservasByUser)
	router.POST("/amadeus/mapping", amadeusController.CreateMapping)
	router.POST("/amadeus/availability", amadeusController.CheckAvailability)
	router.DELETE("/reservas/:id", reservaController.DeleteReserva)
	log.Info("Finishing mappings configurations")
}
