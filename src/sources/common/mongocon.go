package common

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongoconnect function
func Mongoconnect() (client *mongo.Client) {
	Mohost := GetMoHost()
	Moport := GetMoPort()
	
	// Set client options
	clientOptions := options.Client().ApplyURI(Mohost + Moport)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	return client

}
