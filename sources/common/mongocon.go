package common

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongoconnect function
func Mongoconnect() (status bool, msg string, client *mongo.Client) {
	status = true
	msg = ""

	Mohost := GetMoHost()
	Moport := GetMoPort()
	Mouser := GetMoUser()
	Mopass := GetMoPass()
	MoAuthMethod := GetMoAuthMethod()
	MoAuthDb := GetMoAuthDb()

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
		msg = err.Error()
	} else {
		msg = "Success"
		status = true
	}

	return status, msg, client

}
