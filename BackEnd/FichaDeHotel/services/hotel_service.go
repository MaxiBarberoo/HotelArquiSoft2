package services

import (
	"bytes"
	"context"
	"encoding/json"
	e "fichadehotel/Utils"
	hotelClient "fichadehotel/clients/hotel"
	"fichadehotel/dto"
	"fichadehotel/model"
	"fmt"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
	"time"
)

var numImagen = 1

func generateImagenURL() string {
	var urlImagen string

	numero := strconv.Itoa(numImagen)
	urlImagen = "FrontHotel/imagenes/" + numero + ".jpg"
	numImagen++

	return urlImagen
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendToQueue(hotelId string) {
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	print(ctx)
	defer cancel()

	body := hotelId
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id string) (dto.HotelDto, e.ApiError) {

	var hotel, err = hotelClient.GetHotelById(id)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}
	var hotelDto dto.HotelDto

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDto.Name = hotel.Nombre
	hotelDto.CantHabitaciones = hotel.CantHab
	hotelDto.Id = hotel.ID.Hex()
	hotelDto.Desc = hotel.Descripcion
	hotelDto.Amenities = hotel.Amenities
	hotelDto.Ciudad = hotel.Ciudad
	hotelDto.Imagen = hotel.Imagen

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel
	var err error
	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel.Descripcion = hotelDto.Desc
	hotel.Amenities = hotelDto.Amenities
	hotel.Ciudad = hotelDto.Ciudad

	hotel, err = hotelClient.InsertHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel could not be inserted")
	}
	hotelDto.Id = hotel.ID.Hex()

	var dtoAmadeusMapping dto.AmadeusMappingDto
	dtoAmadeusMapping.HotelId = hotelDto.Id
	dtoAmadeusMapping.AmadeusHotelId = hotelDto.AmadeusId

	url := "http://urdnginx:8020/amadeus/mapping"

	jsonData, err := json.Marshal(dtoAmadeusMapping)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error al convertir el json para llamar al servicio de mapeo ")
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return hotelDto, e.NewBadRequestApiError("Error al llamar al servicio de mapeo de ids")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return hotelDto, e.NewBadRequestApiError("Error al llamar al servicio de mapeo de ids")
	}

	var updateImagen dto.HotelDto
	updateImagen.Id = hotelDto.Id
	updateImagen.Imagen = generateImagenURL()
	s.UpdateHotel(updateImagen)

	hotelDto.Imagen = updateImagen.Imagen

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	var ID, err = primitive.ObjectIDFromHex(hotelDto.Id)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel.Descripcion = hotelDto.Desc
	hotel.ID = ID
	hotel.Amenities = hotelDto.Amenities
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Imagen = hotelDto.Imagen

	hotel, err = hotelClient.UpdateHotel(hotel)
	if err != nil {
		var errorHotel dto.HotelDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}
	hotelDto.Id = hotel.ID.Hex()

	return hotelDto, nil
}
