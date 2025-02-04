package Data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client

// ConnectToDB connects to MongoDB Atlas and sets the global Client variable
func ConnectToDB() {
	var err error
	// Replace <username>, <password>, <cluster-url>, and <dbname> with your actual details
	uri := "mongodb+srv://baddamtejasri:skillarcade@cluster0.cw3xm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	
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
