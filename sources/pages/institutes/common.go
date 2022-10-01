package institutes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"techpro.club/sources/common"
)

// Fetch notifications list
func SaveInstitute(newInstituteStruct common.SaveInstitutetruct) (status bool, msg string) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveInstitute := client.Database(dbName).Collection(common.CONST_MO_INSTITUTE)

	_, err := saveInstitute.InsertOne(context.TODO(), newInstituteStruct)

	if err != nil {
		msg = err.Error()

	} else {
		status = true
		msg = "Success"
	}

	return status, msg
}

func GetUnregisteredInstitute(userID primitive.ObjectID) (status bool, msg string, fetchInstituteStruct common.FetchInstitutetruct) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	getInstitute := client.Database(dbName).Collection(common.CONST_MO_INSTITUTE)

	err := getInstitute.FindOne(context.TODO(), bson.M{"userid": userID}).Decode((&fetchInstituteStruct))

	//err := fetchSocials.FindOne(context.TODO(), bson.M{"userid": userID}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&socials)

	if err != nil {
		msg = err.Error()

	} else {
		status = true
		msg = "Success"
	}

	return status, msg, fetchInstituteStruct
}

func UpdateInstitute(newInstituteStruct common.SaveInstitutetruct, instituteID primitive.ObjectID) (status bool, msg string) {
	status = false
	msg = ""

	_, _, client := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	saveInstitute := client.Database(dbName).Collection(common.CONST_MO_INSTITUTE)

	err := saveInstitute.FindOneAndUpdate(context.TODO(), bson.M{"_id": instituteID}, newInstituteStruct)

	if err != nil {
		msg = "Error"

	} else {
		status = true
		msg = "Success"
	}

	return status, msg
}
