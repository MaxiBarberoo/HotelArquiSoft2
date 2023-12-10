import React, { useState, useEffect } from "react"
import '../Stylesheet/Admin.css'
import Header from '../Componentes/Header'

function Admin() {
    const [estadisticas, setEstadisticas] = useState([]);
    const [serviciosEscalables, setServiciosEscalables] = useState([]);
    const [contenedoresSeleccionados, setContenedoresSeleccionados] = useState([]);
    const usuarioValidado = localStorage.getItem('usuarioValidado');

    useEffect(() => {
        if (!usuarioValidado) {
            navigate(`/home`);
        }
        // Llamar a la API para obtener estadísticas cuando el componente se monta
        fetch("http://localhost:8059/stats")  // Actualiza la URL con el puerto correcto
            .then(response => response.json())
            .then(data => setEstadisticas(data))
            .catch(error => console.error("Error al obtener estadísticas:", error));

        // Llamar a la API para obtener servicios escalables al cargar la página
        fetch("http://localhost:8059/services")  // Actualiza la URL con el puerto correcto
            .then(response => response.json())
            .then(data => setServiciosEscalables(data))
            .catch(error => console.error("Error al obtener servicios escalables:", error));
    }, []);

    const handleCheckboxChange = (id) => {
        // Actualizar la lista de contenedores seleccionados
        setContenedoresSeleccionados((prevSelected) => {
            if (prevSelected.has(id)) {
                // Desmarcar el checkbox si ya está seleccionado
                prevSelected.delete(id);
            } else {
                // Marcar el checkbox si no está seleccionado
                prevSelected.add(id);
            }
            return new Set(prevSelected); // Garantizar la inmutabilidad
        });
    };

    const handleDeleteSelectedContainers = () => {
        // Llamar a la API para eliminar contenedores seleccionados
        contenedoresSeleccionados.forEach((id) => {
            fetch(`http://localhost:8059/container/${id}`, {
                method: 'DELETE',
            })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                    alert("Se han borrado los contenedores exitosamente");
                    // Recargar la página después de borrar los contenedores
                    window.location.reload();
                })
                .catch(error => console.error(`Error al eliminar contenedor ${id}:`, error));
        });
        // Limpiar la lista de contenedores seleccionados después de la eliminación
        setContenedoresSeleccionados([]);
    };

    return (
        <div>
            <Header />
            <h2>Módulo de administrador</h2>
            <table className="admin-table">
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Uso de CPU</th>
                        <th>Uso de Memoria</th>
                        <th>Memoria Utilizada</th>
                        <th>Seleccionar</th>
                    </tr>
                </thead>
                <tbody>
                    {estadisticas.map(estadistica => (
                        <tr key={estadistica.Id}>
                            <td>{estadistica.Name}</td>
                            <td>{estadistica.CPUPerc}</td>
                            <td>{estadistica.MemPerc}</td>
                            <td>{estadistica.MemUsage}</td>
                            <td>
                                <input
                                    type="checkbox"
                                    onChange={() => handleCheckboxChange(estadistica.Id)}
                                    checked={contenedoresSeleccionados.has(estadistica.Id)}
                                />
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            <button onClick={handleDeleteSelectedContainers}>Borrar contenedores seleccionados</button>

            <div className="scalable-services-section">
                <h3>Servicios Escalables</h3>
                <ul>
                    {serviciosEscalables.map(servicio => (
                        <li key={servicio}>{servicio}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default Admin;