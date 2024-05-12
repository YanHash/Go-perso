package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Go-perso/types"
)

var client *mongo.Client

func InitMongoDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found !")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set !")
	}

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)

	}
	defer func() {
		if err := cli.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	client = cli
}

func GetTools(w http.ResponseWriter, r *http.Request) {
	var tools []types.Tool
	collection := client.Database("portfolio").Collection("Tools")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error while retrieving Tools documents...")
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tool types.Tool
		if err := cursor.Decode(&tool); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error during iteration on Tools documents...")
			return
		}
		tools = append(tools, tool)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error post-iteration on Tools documents...")
		return
	}
	json.NewEncoder(w).Encode(tools)
}

func GetExperience(w http.ResponseWriter, r *http.Request) {
	var xp []types.Experience
	collection := client.Database("portfolio").Collection("Experience")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error while retrieving Experience documents...")
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var job types.Experience
		if err := cursor.Decode(&job); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error during iteration on Experience documents...")
			return
		}
		xp = append(xp, job)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error post-iteration on Experience documents...")
		return
	}
	json.NewEncoder(w).Encode(xp)
}

func GetEducation(w http.ResponseWriter, r *http.Request) {
	var edu []types.Education
	collection := client.Database("portfolio").Collection("Education")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error while retrieving Education documents...")
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var study types.Education
		if err := cursor.Decode(&study); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error during iteration on Education documents...")
			return
		}
		edu = append(edu, study)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error post-iteration on Education documents...")
		return
	}
	json.NewEncoder(w).Encode(edu)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	var contacts []types.Contact
	collection := client.Database("portfolio").Collection("Contacts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error while retrieving Contacts documents...")
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var contact types.Contact
		if err := cursor.Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error during iteration on Contacts documents...")
			return
		}
		contacts = append(contacts, contact)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error post-iteration on Contacts documents...")
		return
	}
	json.NewEncoder(w).Encode(contacts)
}
