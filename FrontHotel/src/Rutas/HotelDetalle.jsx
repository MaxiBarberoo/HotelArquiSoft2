import React, { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import '../Stylesheet/HotelDetalle.css'
import Header from '../Componentes/Header'
import { useParams } from 'react-router-dom';
import { parse, formatISO } from 'date-fns';

function convertirAFechaISO(fechaString) {
  try {
    // Eliminar la parte de la zona horaria en paréntesis para facilitar el parseo
    const fechaSinZona = fechaString.replace(/\s\(.*\)$/, "");

    // Convertir la cadena a un objeto Date
    const fecha = new Date(fechaSinZona);

    // Formatear la fecha manualmente a ISO 8601
    const año = fecha.getFullYear();
    const mes = (fecha.getMonth() + 1).toString().padStart(2, '0');
    const dia = fecha.getDate().toString().padStart(2, '0');
    const horas = fecha.getHours().toString().padStart(2, '0');
    const minutos = fecha.getMinutes().toString().padStart(2, '0');
    const segundos = fecha.getSeconds().toString().padStart(2, '0');
    const milisegundos = fecha.getMilliseconds().toString().padStart(3, '0');

    return `${año}-${mes}-${dia}T${horas}:${minutos}:${segundos}.${milisegundos}Z`;
  } catch (error) {
    console.error("Error al convertir la fecha:", error);
    return '';
  }
}

function HotelDetalle() {
    const { hotelId } = useParams();
    const { fechaDesde } = useParams();
    const { fechaHasta } = useParams();
    const { userId } = useParams();
    const navigate = useNavigate();
    const [hotel, setHotel] = useState(null);
    const [amenities, setAmenities] = useState([]);
    const usuarioValidado = localStorage.getItem('usuarioValidado');

    useEffect(() => {
        if (!usuarioValidado) {
            navigate(`/`);
        }
        // Define la URL y realiza la solicitud
        const url = `http://localhost:8021/hotels/${hotelId}`;

        fetch(url)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Error al obtener el hotel.');
                }
                return response.json();
            })
            .then(data => {
                setHotel(data);
                console.log(data.imagen);
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

    const handleReserva = () => {
        const url = `http://localhost:8020/reservas`;

        const fechaDesdeISO = convertirAFechaISO(fechaDesde);
        const fechaHastaISO = convertirAFechaISO(fechaHasta);

        // Creamos el cuerpo de la solicitud
        const reserva = {
            user_id: parseInt(userId),
            hotel_id: hotelId,
            fecha_ingreso: fechaDesdeISO,
            fecha_egreso: fechaHastaISO
        };

        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(reserva)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Error al crear la reserva.');
                }
                return response.json();
            })
            .then(data => {
                alert('Reserva creada con éxito.');
                // Aquí puedes manejar acciones post-reserva, como redireccionar a una página de confirmación, mostrar un mensaje, etc.
            })
            .catch(error => {
                alert('Ha habido un error al realizar su reserva. Inténtelo nuevamente.');
                // Maneja el error, mostrando un mensaje al usuario o lo que consideres necesario.
            });
    };

    const handleHome = () => {
        navigate("/");  // Esto redirigirá al usuario a la página principal
    }

    return (
        <div>
            <Header />
            <button className="boton-volver" onClick={handleHome}>Volver</button>
            <div className="contenedor-hoteles">
                <p className="nombre-hotel1">
                    <strong>{hotel.name}</strong>
                </p>
                <img src={`../../../../imagenes/${hotel.imagen}`} />
                <p className="cantidad-piezas">
                    Habitaciones: {hotel.cantHabitaciones}
                </p>
                <p className="descripcion-hotel">
                    Descripción: {hotel.descripcion}
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