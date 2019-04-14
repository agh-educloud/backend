package main

import (
 "context"
 "fmt"
 "go.mongodb.org/mongo-driver/mongo"
 "go.mongodb.org/mongo-driver/mongo/options"
 "net/http"
 "os"
 "time"
)

func main() {

 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "3001"
 }

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  _, _ = fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
 })

 _ = http.ListenAndServe(":"+PORT, nil)

 client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

 ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
 _ = client.Connect(ctx)


}
