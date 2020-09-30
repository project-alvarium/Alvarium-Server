package collections

import(
    "fmt"
	"log"
	"database-manager/configuration"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

Client DatabaseClient

func Database() {
	// Set client options
	clientOptions := options.Client().ApplyURI(configuration.Config.DatabaseUrl)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	DatabaseClient := client
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}