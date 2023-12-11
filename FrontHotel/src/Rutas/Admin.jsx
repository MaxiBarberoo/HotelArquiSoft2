import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import '../Stylesheet/Admin.css';
import Header from '../Componentes/Header';

function Admin() {
    const [estadisticas, setEstadisticas] = useState([]);
    const [serviciosEscalables, setServiciosEscalables] = useState([]);
    const usuarioValidado = localStorage.getItem('usuarioValidado');
    const navigate = useNavigate();
    const [eliminando, setEliminando] = useState(false);
    const [escalando, setEscalando] = useState(false);

    useEffect(() => {
        if (!usuarioValidado) {
            navigate(`/home`);
        }

        cargarEstadisticas();

        fetch("http://localhost:8059/services")
            .then(response => response.json())
            .then(data => setServiciosEscalables(data))
            .catch(error => console.error("Error al obtener servicios escalables:", error));
    }, []);

    const handleCheckboxChange = (ID) => {
        setEstadisticas(prevStats =>
            prevStats.map(stat =>
                stat.ID === ID ? { ...stat, isSelected: !stat.isSelected } : stat
            )
        );
    };

    const cargarEstadisticas = () => {
        fetch("http://localhost:8059/stats")
            .then(response => response.json())
            .then(data => {
                const sortedData = data.map(stat => ({ ...stat, isSelected: false }))
                    .sort((a, b) => a.Name.localeCompare(b.Name));
                setEstadisticas(sortedData);
                setEliminando(false); // Cambiar estado eliminando a false después de cargar estadísticas
            })
            .catch(error => {
                console.error("Error al obtener estadísticas:", error);
                setEliminando(false); // Asegurarse de cambiar el estado incluso si hay un error
            });
    };

    const handleDeleteSelectedContainers = () => {
        setEliminando(true); // Iniciar eliminación
        const selectedIds = estadisticas
            .filter(stat => stat.isSelected)
            .map(stat => stat.ID);

        Promise.all(selectedIds.map(ID =>
            fetch(`http://localhost:8059/container/${ID}`, { method: 'DELETE' })
                .then(response => response.json())
                .then(data => {
                    console.log(`Contenedor ${ID} eliminado:`, data);
                    return ID;
                })
                .catch(error => {
                    console.error(`Error al eliminar contenedor ${ID}:`, error);
                    return null;
                })
        ))
            .then(() => {
                setTimeout(() => {
                    cargarEstadisticas(); // Llamar a cargarEstadisticas después de 5 segundos
                }, 30000);
            });
    };

    const handleScaleService = (servicio) => {
        setEscalando(true);
        fetch("http://localhost:8059/scale", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ servicio })
        })
        .then(response => response.json())
        .then(data => {
            console.log(`Servicio ${servicio} escalado:`, data);
            setTimeout(() => {
                cargarEstadisticas();
                setEscalando(false); // Finalizar el proceso de escalamiento
            }, 5000);
        })
        .catch(error => {
            console.error(`Error al escalar el servicio ${servicio}:`, error);
            setEscalando(false); // Asegurarse de cambiar el estado incluso si hay un error
        });
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
                        <tr key={estadistica.ID}>
                            <td>{estadistica.Name}</td>
                            <td>{estadistica.CPUPerc}</td>
                            <td>{estadistica.MemPerc}</td>
                            <td>{estadistica.MemUsage}</td>
                            <td>
                                <input
                                    type="checkbox"
                                    onChange={() => handleCheckboxChange(estadistica.ID)}
                                    checked={estadistica.isSelected}
                                />
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            <button className="boton-borrar-contenedores" onClick={handleDeleteSelectedContainers}>
                Borrar contenedores seleccionados
            </button>

            {eliminando && <div>Eliminando contenedor...</div>}

            <div className="scalable-services-section">
                <h3>Servicios Escalables</h3>
                <ul>
                    {serviciosEscalables.map(servicio => (
                        <li key={servicio}>
                            {servicio}
                            <button onClick={() => handleScaleService(servicio)} disabled={escalando}>Escalar</button>
                        </li>
                    ))}
                </ul>
            </div>

            {escalando && <div>Escalando servicio...</div>}
        </div>
    );
}

export default Admin;

