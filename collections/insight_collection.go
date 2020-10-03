package collections

import (
	"context"
	"database-manager/configuration"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insight struct holds data id and insight id
type Insight struct {
	DataID    string `json:"dataID" bson:"dataID"`
	InsightID string `json:"isbn" bson:"_id"`
}

var collectionName string = "insights"

//InsertOne : inserts new insight
func InsertOne(dataID string) (string, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(collectionName)
	insight := bson.M{"dataID": dataID}
	fmt.Println("insight: ", insight)
	insertResult, err := collection.InsertOne(context.TODO(), insight)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := insertResult.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return oid, nil
}

// FindOne : finds and returns one insight with the id given
func FindOne(insightID string) (*Insight, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(collectionName)
	objID, err1 := primitive.ObjectIDFromHex(insightID)
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	if err1 != nil {
		return nil, err1
	}
	var insight *Insight
	err := collection.FindOne(context.Background(), filter).Decode(&insight)
	if err != nil {
		return nil, err
	}
	return insight, nil
}
