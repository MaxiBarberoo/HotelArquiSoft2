package services

import (
	e "busquedadehotel/Utils"
	hotelSearchClient "busquedadehotel/clients/HotelSearch"
	"busquedadehotel/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"github.com/streadway/amqp"
)
type hotelSearchService struct{}

type hotelSearchServiceInterface interface {
	GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError)
	UpdateHotel(hotelId string) e.ApiError
}

var (
	HotelSearchService hotelSearchServiceInterface
)

func init() {
	HotelSearchService = &hotelSearchService{}
}

func (s *hotelSearchService) GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError) {

	var wg sync.WaitGroup
	var er e.ApiError
	var hotelsByCity, err = hotelSearchClient.GetHotelsByDateAndCity(searchDto)
	if err != nil {
		var errorHotel dto.HotelsDto
		return errorHotel, e.NewBadRequestApiError("Hotel not found")
	}

	for i := 0; i < len(hotelsByCity); i++ {
		wg.Add(1) // Añade 1 al WaitGroup por cada goroutine

		go func(index int) {
			defer wg.Done() // Marca la goroutine como terminada al finalizar

			url := "http://urdnginx:8020/amadeus/availability"
			searchDto.HotelId = hotelsByCity[index].Id

			jsonData, err := json.Marshal(searchDto)
			if err != nil {
				er = e.NewBadRequestApiError("Error al convertir searchDto a JSON")
				return
			}

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				er = e.NewBadRequestApiError("Error al llamar al microservicio de disponibilidad de hotel")
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				er = e.NewBadRequestApiError("Error con la respuesta obtenida")
				return
			}

			var result map[string]interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				er = e.NewBadRequestApiError("Error al extraer el JSON")
				return
			}

			availability, ok := result["availability"].(bool)
			fmt.Println("Availability:", availability)
			if !ok {
				er = e.NewBadRequestApiError("Availability no es una cadena o no existe en el JSON")
				return
			}

			fmt.Println("Is Available:", availability)
			hotelsByCity[index].Availability = availability
		}(i)
		if er != nil {
			return nil, er
		}
	}

	wg.Wait()

	if er != nil {
		return nil, er
	}

	var availableHotels dto.HotelsDto
	for i := 0; i < len(hotelsByCity); i++ {
		fmt.Println("Hotel Availability: ", hotelsByCity[i].Availability)
		if hotelsByCity[i].Availability == true {
			availableHotels = append(availableHotels, hotelsByCity[i])
		}
	}

	return availableHotels, nil
}

func (s *hotelSearchService) UpdateHotel(hotelId string) e.ApiError {

	url := fmt.Sprintf("http://fichadehotelnginx:8021/hotels/%s", hotelId)

	resp, err := http.Get(url)
	if err != nil {
		return e.NewBadRequestApiError("Error al llamar al microservicio de ficha de hotel")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return e.NewBadRequestApiError("Error con la respuesta obtenida")
	}

	var hotel dto.HotelDto
	err = json.Unmarshal(body, &hotel)
	if err != nil {
		return e.NewBadRequestApiError("Error al extraer el json")
	}

	errSolr := hotelSearchClient.UpdateHotel(hotel)
	if err != nil {
		return errSolr
	}

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Consumer() {
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
		errorSolr := HotelSearchService.UpdateHotel(hotelId)
		if errorSolr != nil {
			fmt.Println(errorSolr)
		}

		log.Printf("Received a message with ID: %d", hotelId)
		// Llamado a actualizar hotel en solr
		// si llaman metodo gethotelbyId o getHotels ahi llamo
		// a API de disponibilidad de hoteles para agregar atributo de disponibilidad
	}
}
