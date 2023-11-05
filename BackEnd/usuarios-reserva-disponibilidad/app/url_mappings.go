package app

import (
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
	router.GET("/reservas/hotel/:hotel_id", reservaController.GetReservasByHotel)
	router.POST("/users/auth", userController.UserAuth)
	router.POST("/reservas/hotelsbyfecha", reservaController.GetHotelsByFecha)
	router.POST("/users", userController.UserInsert)
	router.POST("/reservas", reservaController.ReservaInsert)
	router.POST("/reservas/rooms", reservaController.GetRooms)
	router.POST("/reservas/byfecha", reservaController.GetReservasByFecha)
	router.POST("/reservas/:id", reservaController.GetReservaById)
	router.POST("/reservas/hoteluser", reservaController.GetReservasByHotelAndUser)
	router.POST("/reservas/fechauser", reservaController.GetReservasByFechaAndUser)
	router.POST("/reservas/fechahotel", reservaController.GetReservasByHotelAndFecha)
	router.POST("/reservas/hotelfechauser", reservaController.GetReservasByHotelFechaAndUser)
	router.GET("/reservas/reservauser/:user_id", reservaController.GetReservasByUser)
	log.Info("Finishing mappings configurations")
}
