package common

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongoconnect function
func Mongoconnect() (status bool, msg string, client *mongo.Client) {
	status = true
	msg = ""

	Mohost := GetMoHost()
	Mouser := GetMoUser()
	Mopass := GetMoPass()


	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	// Set client options
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + Mouser + ":" + Mopass + "@" + Mohost + "?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	if err != nil {
		msg = err.Error()
	} else {
		msg = "Success"
		status = true
	}

	return status, msg, client

}
