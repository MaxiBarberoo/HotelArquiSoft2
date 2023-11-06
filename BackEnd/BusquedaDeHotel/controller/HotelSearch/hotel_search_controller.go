package HotelSearch

import (
	"HotelArquiSoft2/BackEnd/BusquedaDeHotel/dto"
	service "HotelArquiSoft2/BackEnd/BusquedaDeHotel/services"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Consumer() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		hotelId := string(d.Body)

		// Update hotel
		errorSolr := service.HotelSearchService.UpdateHotel(hotelId)
		if errorSolr != nil {
			fmt.Println(errorSolr)
		}

		log.Printf("Received a message with ID: %d", hotelId)
		// Llamado a actualizar hotel en solr
		// si llaman metodo gethotelbyId o getHotels ahi llamo
		// a API de disponibilidad de hoteles para agregar atributo de disponibilidad
	}
}

func GetHotelsByDateAndCity(c *gin.Context) {
	var hotelsDto dto.HotelsDto

	var searchDto dto.SearchDto

	err := c.BindJSON(&searchDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hotelsDto, err = service.HotelSearchService.GetHotelsByDateAndCity(searchDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, hotelsDto)
}
