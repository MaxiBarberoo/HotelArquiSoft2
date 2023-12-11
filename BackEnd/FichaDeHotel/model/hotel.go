package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nombre      string             `bson:"nombre"`
	CantHab     int                `bson:"cantHab"`
	Descripcion string             `bson:"descripcion"`
	Ciudad      string             `bson:"ciudad"`
	Amenities   []string           `bson:"amenities"`
	Imagen      string             `bson:"imagen"`
}

type Hotels []Hotel
