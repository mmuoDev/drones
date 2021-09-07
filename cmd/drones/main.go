package main

import (
	"drones/internal/app"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	a := app.New()
	log.Println(fmt.Sprintf("Starting server on port:%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), a.Handler()))
}
