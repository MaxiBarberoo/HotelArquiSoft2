import React, { useState, useEffect } from "react";
import { json, useNavigate } from "react-router-dom";
import Header from "../Componentes/Header";
import HotelesR from '../Componentes/HotelesR'
import DatePicker from "react-datepicker"
import 'react-datepicker/dist/react-datepicker.css'
import "../Stylesheet/Home.css";

function Home() {
  const [hotelesDisponibles, setHotelesDisponibles] = useState([]);
  const [fechaDesde, setFechaDesde] = useState(null);
  const [fechaHasta, setFechaHasta] = useState(null);
  const [ciudad, setCiudad] = useState(null);
  const [busquedaRealizada, setBusquedaRealizada] = useState(false);

  const handleFechaDesdeChange = (date) => {
    setFechaDesde(date);
  };

  const handleFechaHastaChange = (date) => {
    setFechaHasta(date);
  };

  const handleCiudadChange = (event) => {
    setCiudad(event.target.value);
  }


  const buscarHotelesDisponibles = (event) => {
    event.preventDefault();
    if (!fechaDesde || !fechaHasta || !ciudad) {
      alert("Debes completar los campos de fecha y de ciudad.");
    } else {
      // Define la URL y los parámetros de la solicitud
      const url = 'http://localhost:8090/hotels'; 
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          ciudad: ciudad,
          fecha_ingreso: fechaDesde,
          fecha_egreso: fechaHasta
        })
      };

      // Realiza la solicitud
      fetch(url, requestOptions)
        .then(response => {
          if (!response.ok) {
            throw new Error('Respuesta no válida desde el servidor');
          }
          return response.json();
        })
        .then(data => {
          setHotelesDisponibles(data);
        })
        .catch(error => {
          console.error('Hubo un problema con la solicitud fetch:', error);
        });
    }
  };

  /*useEffect(() => {
    const fetchAmenitiesForHotels = async () => {
      const hotelsWithAmenities = await Promise.all(
        amenities.map(async (amenitie) => {
          const response = await fetch(`http://localhost:8090/amenities/${amenitie.id}`);
          if (response.ok) {
            const amenitiesData = await response.json();
            return { ...amenitie, amenities: amenitiesData };
          } else {
            console.error(`Error en la petición GET de amenities para el hotel ${hotel.id}`);
            return amenitie;
          }
        })
      );
      setHoteles(hotelsWithAmenities);
    };
  
    if (amenities.length > 0) {
      fetchAmenitiesForHotels();
    }
  }, [hotelesDisponibles]);*/

  return (
    <div>
      <Header />
      <div className="contenedor-criterios">
        <h2>Ingrese las fechas y ciudad para su estadia:</h2>
        <div className="contenedor-fechas">
          <div className="fecha-desde">
            <p>Desde: </p>
            <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
          </div>
          <div className="fecha-hasta">
            <p>Hasta: </p>
            <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
          </div>
        </div>
        <div className="nombre-ciudad">
          <p>Ciudad: </p>
          <input type="text" placeholder="Ciudad..." onChange={handleCiudadChange}></input>
        </div>
      </div>
      <form className="contenedor-buscar" onSubmit={buscarHotelesDisponibles}>
        <button className="boton-buscar">Buscar</button>
      </form>
      <div className="contenedor-hoteles-r">
        {busquedaRealizada ? (
          hotelesDisponibles.length ? (
            hotelesDisponibles.map((hotel) => (
              <div key={hotel.id}>
                <HotelesR
                  key={hotel.id}
                  hotelId={hotel.id}
                  descripcion={hotel.descripcion}
                  nombreHotel={hotel.name}
                  fechaDesde={fechaDesde}
                  fechaHasta={fechaHasta}
                />
              </div>
            ))
          ) : (
            <p>No hay hoteles disponibles en esas fechas.</p>
          )
        ) : null}
      </div>
    </div>
  );
}

export default Home;











