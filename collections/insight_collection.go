package collections

import (
    "fmt"
	"log"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"database-manager/configuration"
)

type insight struct {
	dataID string
	insightID primitive.ObjectID
}

collection_name string = "insights"
collection := client.Database(configuration.Config.DatabaseName).Collection(collection_name)
func InsertOne(dataID string)  {
	insight = Insight{dataID} 
	insertResult, err := collection.InsertOne(context.TODO(), insight)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult
}

func FindOne(insightID string) Insight{
	filter := bson.D{{insightID }}
	var insight Insight 
	err := collection.FindOne(context.Background(), filter).Decode(&insight)
	if err != nil { return err }
	return insight
}

