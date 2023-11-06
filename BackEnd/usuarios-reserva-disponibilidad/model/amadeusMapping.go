
package model


type AmadeusMapping struct {
	ID              int          `gorm:"primaryKey"`
  HotelId         string       `gorm:"type:varchar(250);not null"`
  AmadeusHotelId  string       `gorm:"type:varchar(250);not null"`
}

type AmadeusMappings []AmadeusMapping
