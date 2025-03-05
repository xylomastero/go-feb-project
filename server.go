package main

import (
	"fmt"
	"log"
	"net/http"

	"main/api/db"
	"main/api/routes"
)

func Start() {
	if err := db.Initialize(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := routes.SetupRoutes()

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
