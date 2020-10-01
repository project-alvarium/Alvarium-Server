package collections

import(
	"context"
	"fmt"
	"log"
	"database-manager/configuration"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseClient mongo.Client

func Database() {
	// Set client options
	clientOptions := options.Client().ApplyURI(configuration.Config.DatabaseURL)
	// Connect to MongoDB
	DatabaseClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = DatabaseClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}