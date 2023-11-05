package hotel

import (
	"HotelArquiSoft2/BackEnd/FichaDeHotel/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

var Db *mongo.Database

func GetHotelById(id string) (model.Hotel, error) {
	var hotel model.Hotel
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return hotel, err
	}
	err = Db.Collection("hotels").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		fmt.Println(err)
		return hotel, nil
	} else {
		return hotel, err
	}
}

func InsertHotel(hotel model.Hotel) (model.Hotel, error) {
	hotel.ID = primitive.NewObjectID()
	_, err := Db.Collection("hotels").InsertOne(context.TODO(), &hotel)

	if err != nil {
		fmt.Println(err)
		return hotel, nil
	} else {
		return hotel, err
	}
}

func UpdateHotel(hotel model.Hotel) (model.Hotel, error) {

	filter := bson.M{"_id": hotel.ID}

	// Create an empty update operation
	update := bson.M{}

	// Check each field in updatedHotel and add it to the update operation if it's not null or empty
	if hotel.Nombre != "" {
		update["nombre"] = hotel.Nombre
	}

	if hotel.CantHab > 0 {
		update["cantHab"] = hotel.CantHab
	}

	if hotel.Descripcion != "" {
		update["descripcion"] = hotel.Descripcion
	}

	if len(hotel.Amenities) > 0 {
		update["amenities"] = hotel.Amenities
	}

	if hotel.Ciudad != "" {
		update["ciudad"] = hotel.Ciudad
	}

	_, err := Db.Collection("hotels").UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return hotel, err
	} else {
		return hotel, nil
	}
}
