package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongoDB(collection string) (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found !")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set !")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	coll := client.Database("portfolio").Collection(collection)

	return coll, client, ctx, cancel
}

func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}

func CreateCollections(collections []string) {
	var cli *mongo.Client
	var ctx context.Context
	var cancel context.CancelFunc

	for _, item := range collections {
		_, cli, ctx, cancel = SetupMongoDB(item)
		log.Printf("Collection %s created", item)
	}

	CloseConnection(cli, ctx, cancel)
}
