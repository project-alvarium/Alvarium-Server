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
var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(collectionName)

//InsertOne : inserts new insight
func InsertOne(dataID string) (string, error) {
	insight := Insight{DataID: dataID}
	insertResult, err := collection.InsertOne(context.TODO(), insight)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := insertResult.InsertedID.(primitive.ObjectID)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return oid.String(), nil
}

// FindOne : finds and returns one insight with the id given
func FindOne(insightID string) (*Insight, error) {
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
