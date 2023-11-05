import React, { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import '../Stylesheet/HotelDetalle.css'
import Header from '../Componentes/Header'
import { useParams, useLocation } from 'react-router-dom';

function HotelDetalle() {
    const { hotelId } = useParams();
    const navigate = useNavigate();
    const [hotel, setHotel] = useState(null);
    const [amenities, setAmenities] = useState([]);
    const location = useLocation();
    const fechaDesde = location.state?.fechaDesde;
    const fechaHasta = location.state?.fechaHasta;

    useEffect(() => {
        // Define la URL y realiza la solicitud
        const url = `http://localhost:8090/hotels/${hotelId}`; 

        fetch(url)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Error al obtener el hotel.');
                }
                return response.json();
            })
            .then(data => {
                setHotel(data);
                // Suponiendo que los amenities vienen en el campo 'amenities' de la respuesta:
                setAmenities(data.amenities);
            })
            .catch(error => {
                console.error('Hubo un problema con la solicitud fetch:', error);
            });
    }, [hotelId]); // El useEffect se ejecutará cada vez que cambie 'hotelId'

    if (!hotel) {
        return <div>Cargando...</div>; // Mostrar un loader o algún feedback mientras los datos se están cargando
    }

    const handleReserva = () =>{
        
    }

    return (
        <div>
            <Header />
            <div className="contenedor-hoteles">
                <p className="nombre-hotel1">
                    <strong>{props.nombreHotel}</strong>
                </p>
                <p className="cantidad-piezas">
                    Habitaciones: {props.piezas}
                </p>
                <p className="descripcion-hotel">
                    Descripción: {props.descripcion}
                </p>

                <h2>Amenities del hotel:</h2>
                <ul>
                    {amenities.map((amenitie, index) => (
                        <li key={index}>{amenitie}</li>
                    ))}
                </ul>
                <button className="boton-reservar" onClick={handleReserva}>Reservar</button>
            </div>
        </div>
    );
}

export default HotelDetalle;