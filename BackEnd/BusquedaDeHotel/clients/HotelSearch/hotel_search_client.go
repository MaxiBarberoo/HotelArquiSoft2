package clients

import (
	e "busquedadehotel/Utils"
	dto "busquedadehotel/dto"
	"fmt"
	solr "github.com/rtt/Go-Solr"
)

func UpdateHotel(hotelDto dto.HotelDto) e.ApiError {
	s, err := solr.Init("localhost", 8983, "Hotels")

	if err != nil {
		fmt.Println(err)
		return e.NewBadRequestApiError("Error al conectarse a Solr")
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"hotel_id": hotelDto.Id, "name": hotelDto.Name, "ciudad": hotelDto.Ciudad, "cantHabitaciones": hotelDto.CantHabitaciones, "descripcion": hotelDto.Desc, "amenities": hotelDto.Amenities},
		},
	}

	// send off the update (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(f, true)

	if err != nil {
		return e.NewBadRequestApiError("Error al guardar el hotel en Solr")
	}

	fmt.Println("Solr response: ", resp)
	return nil
}

func GetHotelsByDateAndCity(searchDto dto.SearchDto) (dto.HotelsDto, e.ApiError) {
	s, err := solr.Init("localhost", 8983, "Hotels")

	if err != nil {
		fmt.Println(err)
		return nil, e.NewBadRequestApiError("Error al conectarse a solr")
	}

	// Build a query object
	// Here we are specifying a 'q' param,
	// rows, faceting and facet.fields
	q := solr.Query{
		Params: solr.URLParamMap{
			"q":           []string{"ciudad:" + searchDto.Ciudad},
			"facet.field": []string{"hotel_id", "name", "cantHabitaciones", "descripcion", "amenities"},
			"facet":       []string{"true"},
		},
		Rows: 100000,
	}

	// perform the query, checking for errors
	res, err := s.Select(&q)

	if err != nil {
		fmt.Println(err)
		return nil, e.NewBadRequestApiError("Error al conectarse al buscar los hoteles")
	}

	// grab results for ease of use later on
	results := res.Results
	var solrDto dto.SolrDto

	var hotelsByCity dto.HotelsDto
	var hotelDto dto.HotelDto
	for i := 0; i < results.Len(); i++ {
		// Asigna el valor del campo "hotel_id" a una variable de tipo interface{}
		hotelIdInterface := results.Get(i).Field("hotel_id")

		// Verifica si el valor es un slice de interfaces
		if hotelIds, ok := hotelIdInterface.([]interface{}); ok {
			// Inicializa un nuevo slice de strings
			var hotelIdStrings []string

			// Convierte cada elemento a string
			for _, id := range hotelIds {
				if strId, ok := id.(string); ok {
					hotelIdStrings = append(hotelIdStrings, strId)
				}
			}

			// Asigna el nuevo slice de strings a tu variable solrDto.Id
			solrDto.Id = hotelIdStrings
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("ID is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelDto.Id = solrDto.Id[0]

		nameInterface := results.Get(i).Field("name")
		if names, ok := nameInterface.([]interface{}); ok {
			// Inicializa un nuevo slice de strings
			var nameStrings []string

			// Convierte cada elemento a string
			for _, name := range names {
				if strName, ok := name.(string); ok {
					nameStrings = append(nameStrings, strName)
				}
			}

			// Asigna el nuevo slice de strings a tu variable solrDto.Id
			solrDto.Name = nameStrings
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Name is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		fmt.Println(solrDto.Name[0])
		hotelDto.Name = solrDto.Name[0]

		hotelDescInterface := results.Get(i).Field("descripcion")
		if descs, ok := hotelDescInterface.([]interface{}); ok {
			// Inicializa un nuevo slice de strings
			var descStrings []string

			// Convierte cada elemento a string
			for _, desc := range descs {
				if strDesc, ok := desc.(string); ok {
					descStrings = append(descStrings, strDesc)
				}
			}

			// Asigna el nuevo slice de strings a tu variable solrDto.Id
			solrDto.Desc = descStrings
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Descripcion is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		fmt.Println(solrDto.Desc[0])
		hotelDto.Desc = solrDto.Desc[0]

		hotelAmenitiesInterface := results.Get(i).Field("amenities")
		if amenities, ok := hotelAmenitiesInterface.([]interface{}); ok {
			// Inicializa un nuevo slice de strings
			var amenitiesStrings []string

			// Convierte cada elemento a string
			for _, amenitie := range amenities {
				if strAmenitie, ok := amenitie.(string); ok {
					amenitiesStrings = append(amenitiesStrings, strAmenitie)
				}

			}

			// Asigna el nuevo slice de strings a tu variable solrDto.Id
			solrDto.Amenities = amenitiesStrings
		} else {
			// Maneja el caso en que el tipo no sea el esperado
			fmt.Println("Amenities is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		fmt.Println(solrDto.Amenities[0])
		hotelDto.Amenities = solrDto.Amenities
		hotelDto.Ciudad = searchDto.Ciudad

		hotelsByCity = append(hotelsByCity, hotelDto)

	}

	return hotelsByCity, nil
}
