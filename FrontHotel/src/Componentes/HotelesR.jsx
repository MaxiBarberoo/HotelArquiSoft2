import React, { useEffect, useRef, useState } from 'react';
import '../Stylesheet/HotelesR.css'
import { useNavigate } from 'react-router-dom';

function HotelesR(props) {
    const navigate = useNavigate();

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
            <h3 className="nombre-hotel1">
                {props.nombreHotel}
            </h3>
            <p className="descripcion-hotel">
                Descripción: {props.descripcion}
            </p>
            <button onClick={handleVerDetallesClick} className="boton-detalles">Ver detalles</button>
        </div>
    );    
}

export default HotelesR;
