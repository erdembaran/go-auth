package database

import (
	"context"
	"fmt"
	"log"

	"github.com/erdembaran/go-auth/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectDB() {
	MONGODB_URI	:= config.GetEnv("MONGO_URI", "")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Collection = client.Database("go-auth").Collection("user")
}