package common

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongoconnect function
func Mongoconnect() (client *mongo.Client, status bool) {
	Mohost := GetMoHost()
	Moport := GetMoPort()
	Mouser := GetMoUser()
	Mopass := GetMoPass()
	MoAuthMethod := GetMoAuthMethod()
	MoAuthDb := GetMoAuthDb()

	status = true

	credentials := options.Credential{
		Username: Mouser,
		Password: Mopass,
		AuthMechanism:MoAuthMethod,
		AuthSource: MoAuthDb,
	}
	
	// Set client options
	clientOptions := options.Client().ApplyURI(Mohost + Moport).SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		status = false
		log.Fatal(err.Error())
	}

	return client, status

}
