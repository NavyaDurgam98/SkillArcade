package Data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var Client *mongo.Client
var DBName string

// ConnectToDB connects to MongoDB Atlas and sets the global Client variable
func ConnectToDB() {
	var err error

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB credentials from environment variables
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoCluster := os.Getenv("MONGO_CLUSTER")

	// Replace <username>, <password>, <cluster-url>, and <dbname> with your actual details
	//uri := "mongodb+srv://baddamtejasri:skillarcade@cluster0.cw3xm.mongodb.net/?retryWrites=true&w=majority"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		mongoUser, mongoPassword, mongoCluster)

	// Create a new MongoDB client
	Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")
}

// GetDatabase returns a MongoDB database
func GetDatabase(dbName string) *mongo.Database {
	return Client.Database(dbName)
}

// GetCollection returns a collection from the specified database
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return GetDatabase(dbName).Collection(collectionName)
}
