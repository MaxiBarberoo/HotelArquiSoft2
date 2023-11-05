import React, { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import '../Stylesheet/HotelDetalle.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker'

import { useParams } from 'react-router-dom';

function HotelDetalle() {
    const { hotelId } = useParams();
    const navigate = useNavigate();
    const [amenities, setAmenities] = useState([]);
    const hotelIdRef = useRef(props.hotelId);
    const location = useLocation();
    const fechaDesde = location.state?.fechaDesde;
    const fechaHasta = location.state?.fechaHasta;

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
                    {amenities.map(amenitie => (
                        <li key={amenitie.id}>{amenitie.tipo}</li>
                    ))}
                </ul>
                <button className="boton-reservar">Reservar</button>
            </div>
        </div>
    );
}

export default HotelDetalle;