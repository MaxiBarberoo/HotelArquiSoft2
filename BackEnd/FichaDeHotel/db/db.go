package db

import (
	hotelClient "HotelArquiSoft2/BackEnd/FichaDeHotel/clients/hotel"
	"HotelArquiSoft2/BackEnd/FichaDeHotel/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func generateFilename(number int) string {
	// Base directory and filename template
	baseDirectory := "imagenes/"
	baseFilename := "%d.jpg"

	// Format the filename with the given number
	filename := fmt.Sprintf(baseFilename, number)

	// Concatenate the directory and filename
	fullPath := baseDirectory + filename

	return fullPath
}

func readImageAsBlob(filepath string) ([]byte, error) {
	// Read the image file
	imageData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func init() {
	// DB Connections Paramters
	DBName := "pruebaHash"
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
	hotelClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.

	db.AutoMigrate(&model.Hotels{})

	log.Info("Finishing Migration Database Tables")
}
