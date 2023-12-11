import '../Stylesheet/HotelesR.css'

function HotelesR(props) {
    return (
        <div className="contenedor-hoteles">
            <img src={`../Imagenes/${props.imagen}.jpg`} alt={props.nombreHotel} />
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
