package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"net/http"
	"time"
)


type User struct {
	Id int
	OAuthService string
}


func main() {


	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	collection := client.Database("test").Collection("users")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)


	FirstUser := User{11, "None"}

	_, _ = collection.InsertOne(ctx, FirstUser)


	filter := bson.M{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

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
