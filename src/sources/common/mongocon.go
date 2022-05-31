package common

import (
	"context"
	"log"
	"os"

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

func GetMoHost() (moHost string) {

	moHost, exists := os.LookupEnv("MO_HOST")
	if !exists {
		log.Fatal("MO_HOST not defined in .env file")
	}

	return moHost
}

func GetMoPort() (moPort string) {

	moPort, exists := os.LookupEnv("MO_PORT")
	if !exists {
		log.Fatal("MO_PORT not defined in .env file")
	}

	return moPort
}

func GetMoDb() (moDb string) {

	moDb, exists := os.LookupEnv("MO_DATABASE")
	if !exists {
		log.Fatal("MO_DATABASE not defined in .env file")
	}

	return moDb
}

func GetMoUser() (moUser string) {

	moUser, exists := os.LookupEnv("MO_USER")
	if !exists {
		log.Fatal("MO_USER not defined in .env file")
	}

	return moUser
}

func GetMoPass() (moPass string) {

	moPass, exists := os.LookupEnv("MO_PASS")
	if !exists {
		log.Fatal("MO_PASS not defined in .env file")
	}

	return moPass
}