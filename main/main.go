package main

import (
	"../webAuth"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id           int
	OAuthService string
}

func main() {

	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: webAuth.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

}

func BuildUri() string {
	return "mongodb://" + os.Getenv("MONGO_ROOT_USERNAME") + ":" + os.Getenv("MONGO_ROOT_PASSWORD") + "@mongo:27017"
}
