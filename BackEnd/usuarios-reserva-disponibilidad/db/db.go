package db

import (
	reservaClient "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/clients/reserva"
	userClient "HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/clients/user"
	"HotelArquiSoft2/BackEnd/usuarios-reserva-disponibilidad/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "usuarios-reserva-disponibilidad"
	DBUser := "root"
	DBPass := "arquisoft1"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "db"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":8090)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}
	// We need to add all CLients that we build
	userClient.Db = db
	reservaClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reservas{})
  db.AutoMigrate(&model.AmadeusMapping{})


	log.Info("Finishing Migration Database Tables")
}
