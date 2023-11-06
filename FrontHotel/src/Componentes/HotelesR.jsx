import '../Stylesheet/HotelesR.css'

function HotelesR(props) {
    return (
        <div className="contenedor-hoteles">
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
