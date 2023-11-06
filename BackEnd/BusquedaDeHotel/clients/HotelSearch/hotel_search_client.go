package clients

import (
	e "HotelArquiSoft2/BackEnd/BusquedaDeHotel/Utils"
	dto "HotelArquiSoft2/BackEnd/BusquedaDeHotel/dto"
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
			"facet.field": []string{"id", "name", "cantHabitaciones", "descripcion", "amenities"},
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

	var hotelsByCity dto.HotelsDto
	for i := 0; i < results.Len(); i++ {

		hotelIdInterface := results.Get(i).Field("id")
		hotelId, ok := hotelIdInterface.(string)
		if !ok {
			// handle the error; this means that the "id" field is not a string
			fmt.Println("ID is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelsByCity[i].Id = hotelId

		hotelNameInterface := results.Get(i).Field("name")
		hotelName, ok := hotelNameInterface.(string)
		if !ok {
			// handle the error; this means that the "id" field is not a string
			fmt.Println("Name is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelsByCity[i].Name = hotelName

		hotelCantHabitacionesInterface := results.Get(i).Field("cantHabitaciones")
		hotelCantHabitaciones, ok := hotelCantHabitacionesInterface.(int)
		if !ok {
			// handle the error; this means that the "id" field is not a string
			fmt.Println("CantHabitaciones is not an int!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelsByCity[i].CantHabitaciones = hotelCantHabitaciones

		hotelDescInterface := results.Get(i).Field("descripcion")
		hotelDesc, ok := hotelDescInterface.(string)
		if !ok {
			// handle the error; this means that the "id" field is not a string
			fmt.Println("Descripcion is not a string!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelsByCity[i].Desc = hotelDesc

		hotelAmenitiesInterface := results.Get(i).Field("amenities")
		hotelAmenities, ok := hotelAmenitiesInterface.([]string)
		if !ok {
			// handle the error; this means that the "id" field is not a string
			fmt.Println("Descripcion is not a string array!")
			return nil, e.NewBadRequestApiError("Error con un tipo de dato de solr")
		}
		hotelsByCity[i].Amenities = hotelAmenities

		hotelsByCity[i].Ciudad = searchDto.Ciudad
	}

	return hotelsByCity, nil
}
