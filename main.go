package main

import (
	"fmt"
	"github.com/agh-eduweb/webAuth"
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
