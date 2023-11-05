import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Home from './Rutas/Home'
import HotelDetalle from './Rutas/HotelDetalle'
import './main.css'

createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/detalle/:hotel_id" element={<HotelDetalle />} />
      </Routes>
    </Router>
  </React.StrictMode>
);