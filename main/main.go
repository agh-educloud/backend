package main

import (
	"../chat"
	"../class"
	"../feedback"
	"../homework"
	"sync"
)

type User struct {
	Id           int
	OAuthService string
}

//func init() {
//	err := gotenv.Load("./../.env")
//	if err != nil {
//		panic("Error loading .env!")
//	}
//}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go homework.StartServer()

	wg.Add(2)
	go chat.StartServer()

	wg.Add(3)
	go class.StartServer()

	wg.Add(4)
	go feedback.StartServer()

	wg.Wait()

	//uri := BuildUri()
	//client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	//
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//_ = client.Connect(ctx)
	//
	//collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_DATABASE_COLLECTION"))
	//
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//
	//// Insert Example user
	//FirstUser := User{1, "None"}
	//
	//_, _ = collection.InsertOne(ctx, FirstUser)
	//
	//filter := bson.M{}
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//
	//// Find inserted user
	//result := User{}
	//err := collection.FindOne(ctx, filter).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var PORT string
	//if PORT = os.Getenv("PORT"); PORT == "" {
	//	PORT = "3001"
	//}
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	_, _ = fmt.Fprintf(w, "Id of insterted user: %d \n", FirstUser.Id)
	//})
	//
	//_ = http.ListenAndServe(":"+PORT, nil)

}
