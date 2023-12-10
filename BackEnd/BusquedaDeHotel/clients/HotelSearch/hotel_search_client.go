package clients

import (
	e "busquedadehotel/Utils"
	dto "busquedadehotel/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func UpdateHotel(hotelDto dto.HotelDto) e.ApiError {
	fmt.Println("UpdateHotel Client running\n")

	// Crear el documento a enviar a Solr
	document := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"hotel_id":         hotelDto.Id,
				"name":             hotelDto.Name,
				"ciudad":           hotelDto.Ciudad,
				"cantHabitaciones": hotelDto.CantHabitaciones,
				"descripcion":      hotelDto.Desc,
				"amenities":        hotelDto.Amenities,
			},
		},
	}
	fmt.Printf("%v\n", document)

	// Convertir el documento a formato JSON
	jsonDocument, err := json.Marshal(document)
	if err != nil {
		fmt.Println("Error json Marshal: ", err)
		return e.NewBadRequestApiError("Error al convertir documento a JSON")
	}

	// Establecer la URL de Solr donde se enviarán los datos
	solrURL := "http://solr:8983/solr/hotels/update?commit=true" // Reemplaza con la URL correcta de tu colección

	// Crear una solicitud HTTP POST para enviar el documento a Solr
	req, err := http.NewRequest("POST", solrURL, bytes.NewBuffer(jsonDocument))
	if err != nil {
		fmt.Println("Error al crear la request request: ", err)
		return e.NewBadRequestApiError("Error al crear la solicitud HTTP")
	}

	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error solr request: ", err)
		return e.NewBadRequestApiError("Error al realizar la solicitud HTTP")
	}
	defer resp.Body.Close()

	// Verificar el código de respuesta de Solr
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error http status code: ", resp)
		return e.NewBadRequestApiError("Solr respondió con un código de estado no válido")
	}

	// Leer la respuesta de Solr
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decode resultado solr: ", err)
		return e.NewBadRequestApiError("Error al decodificar la respuesta de Solr")
	}

	// Imprimir la respuesta de Solr
	fmt.Println("Respuesta de Solr:", result)

	return nil
}

func GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError) {
	// Crear la consulta a Solr con parámetros de búsqueda
	solrURL := fmt.Sprintf("http://solr:8983/solr/hotels/select?q=ciudad:%s&facet.field=hotel_id&facet.field=name&facet.field=cantHabitaciones&facet.field=descripcion&facet.field=amenities&facet=true&rows=100000", searchDto.Ciudad)

	// Realizar la solicitud GET a Solr
	resp, err := http.Get(solrURL)
	if err != nil {
		fmt.Println("Error request solr search: ", err)
		return nil, e.NewBadRequestApiError("Error al realizar la solicitud HTTP a Solr")
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta HTTP
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response solr search: ", resp)
		return nil, e.NewBadRequestApiError("Solr respondió con un código de estado no válido")
	}

	// Decodificar la respuesta JSON de Solr
	var solrResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&solrResponse)
	if err != nil {
		fmt.Println("Error decode json response: ", err)
		return nil, e.NewBadRequestApiError("Error al decodificar la respuesta JSON de Solr")
	}

	// Procesar los resultados de Solr
	var hotelsByCity dto.HotelsDto

	response := solrResponse["response"].(map[string]interface{})
	docs := response["docs"].([]interface{})

	for _, doc := range docs {
		hotel := doc.(map[string]interface{})

		var hotelDto dto.HotelDto
		hotelDto.Ciudad = searchDto.Ciudad

		if val, ok := hotel["hotel_id"].([]interface{}); ok {
			if str, ok := val[0].(string); ok {
				hotelDto.Id = str
			}
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("ID is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}

		if val, ok := hotel["name"].([]interface{}); ok {
			if str, ok := val[0].(string); ok {
				hotelDto.Name = str
			}
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Name is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}

		if val, ok := hotel["descripcion"].([]interface{}); ok {
			if str, ok := val[0].(string); ok {
				hotelDto.Desc = str
			}
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Descripcion is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}

		if val, ok := hotel["amenities"].([]interface{}); ok {
			var amenitiesSlice []string
			for _, v := range val {
				if subVal, ok := v.(string); ok {
					amenitiesSlice = append(amenitiesSlice, subVal)
				}
			}
			hotelDto.Amenities = amenitiesSlice
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Amenitis is not a string array!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}

		// Agregar hotelDto a hotelsByCity
		hotelsByCity = append(hotelsByCity, hotelDto)
	}

	return hotelsByCity, nil
}
