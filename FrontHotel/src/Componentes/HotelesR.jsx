import React, { useEffect, useRef, useState } from 'react';
import '../Stylesheet/HotelesR.css'
import { useNavigate } from 'react-router-dom';

function HotelesR(props) {
    navigate = useNavigate();

    const handleVerDetallesClick = () => {
        navigate(
            `/detalle/${props.hotelId}`,
            {
                state: {
                    fechaDesde: props.fechaDesde,
                    fechaHasta: props.fechaHasta
                }
            }
        );
    }

    return (
        <div className="contenedor-hoteles">
            <p className="nombre-hotel1">
                <strong>{props.nombreHotel}</strong>
            </p>
            <p className="cantidad-piezas">Habitaciones: {props.piezas}</p>
            <p className="descripcion-hotel">
                Descripción: {props.descripcion}
            </p>
            <form onSubmit={checkDisponibilidad} className="boton-reservar">
                <button onClick={handleVerDetallesClick} className="boton-detalles">Ver detalles</button>
            </form>
        </div>
    );
}

export default HotelesR;
