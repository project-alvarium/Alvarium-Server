// Package collections This is the package comment,
// a top-level piece of documentation
// used to explain things about the package (see json or exp/template)
// All godoc comments are in this form
// with no whitespace between them and what they accompany
package collections

import (
	"context"
	"fmt"
	"github.com/project-alvarium/Alvarium-Server/configuration"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseClient provide interface to mongodb
var DatabaseClient mongo.Client

// Database connect to mongodb
func Database() {
	// Set client options
	clientOptions := options.Client().ApplyURI(configuration.Config.DatabaseURL)
	// Connect to MongoDB
	DatabaseClient1, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DatabaseClient = *DatabaseClient1
	// Check the connection
	err = DatabaseClient1.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}
