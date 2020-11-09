package main

import (
	"github.com/joshiaj7/go-jwt/internal/controller"

	"net/http"
	"log"
)

func main() {
	controller.SetupRoutes()

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}