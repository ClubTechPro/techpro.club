package general

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongoconnect function
func Mongoconnect() *mongo.Client {
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

func GetMoHost() string {

	MoHost, exists := os.LookupEnv("MO_HOST")
	if !exists {
		log.Fatal("MO_HOST not defined in .env file")
	}

	return MoHost
}

func GetMoPort() string {

	MoPort, exists := os.LookupEnv("MO_PORT")
	if !exists {
		log.Fatal("MO_PORT not defined in .env file")
	}

	return MoPort
}

func GetMoDb() string {

	MoDb, exists := os.LookupEnv("MO_DATABASE")
	if !exists {
		log.Fatal("MO_DATABASE not defined in .env file")
	}

	return MoDb
}

func GetMoUser() string {

	MoUser, exists := os.LookupEnv("MO_USER")
	if !exists {
		log.Fatal("MO_USER not defined in .env file")
	}

	return MoUser
}

func GetMoPass() string {

	MoPass, exists := os.LookupEnv("MO_PASS")
	if !exists {
		log.Fatal("MO_PASS not defined in .env file")
	}

	return MoPass
}