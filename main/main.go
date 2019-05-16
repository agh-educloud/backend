package main

import (
	"context"
	"fmt"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id           int
	OAuthService string
}

func init() {
	err := gotenv.Load("./../.env")
	if err != nil {
		panic("Error loading .env!")
	}
}

func main() {

	uri := BuildUri();
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_DATABASE_COLLECTION"))

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	// Insert Example user
	FirstUser := User{1, "None"}

	_, _ = collection.InsertOne(ctx, FirstUser)

	filter := bson.M{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	// Find inserted user
	result := User{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Id of insterted user: %d \n", FirstUser.Id)
	})

	_ = http.ListenAndServe(":"+PORT, nil)

}

func BuildUri() string {
	return "mongodb://" + os.Getenv("MONGO_ROOT_USERNAME") + ":" + os.Getenv("MONGO_ROOT_PASSWORD") + "@mongo:27017"
}
