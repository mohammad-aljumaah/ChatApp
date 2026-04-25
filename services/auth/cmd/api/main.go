package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mohammad-aljumaah/ChatApp/auth/internal/handlers"
)

const webPort = "8080"

type Config struct {
	Handlers *handlers.Handler
}

func main() {
	log.Println("Starting auth service")

	// TODO: connect to database

	app := Config{
		Handlers: handlers.NewHandler(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Auth service listening on port: ", webPort, "")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error start listening. auth service: ", err)
	}

}
