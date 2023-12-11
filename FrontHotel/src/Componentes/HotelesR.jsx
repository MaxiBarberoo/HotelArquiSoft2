import '../Stylesheet/HotelesR.css'
import React, { useState, useEffect } from "react";

function HotelesR(props) {
    const rutaImagen = props.imagen;
    
      return (
          <div className="contenedor-hoteles">
              <img src={`../../imagenes/${rutaImagen}`} alt={props.nombreHotel} />
              <h3 className="nombre-hotel1">
                  {props.nombreHotel}
              </h3>
              <p className="descripcion-hotel">
                  Descripción: {props.descripcion}
              </p>
          </div>
      );  
}

export default HotelesR;
