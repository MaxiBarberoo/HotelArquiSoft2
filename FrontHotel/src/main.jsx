import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Home from './Rutas/Home'
import HotelDetalle from './Rutas/HotelDetalle'
import LoginRegister from './Rutas/LoginRegister'
import Admin from './Rutas/Admin'
import './main.css'

createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<LoginRegister />} />
        <Route path="/home/:email" element={<Home />} />
        <Route path="/admin/:email" element={<Admin />} />
        <Route path="/detalle/:hotelId/:fechaDesde/:fechaHasta" element={<HotelDetalle />} />
      </Routes>
    </Router>
  </React.StrictMode>
);