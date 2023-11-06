package model

import "time"

type Reserva struct {
	ID       int       `gorm:"primaryKey"`
	FechaIn  time.Time `gorm:"type:DATE"`
	FechaOut time.Time `gorm:"type:DATE"`
	UserId   int       `gorm:"foreignKey"`
	HotelId  string    `gorm:"foreignKey"`
}

type Reservas []Reserva
