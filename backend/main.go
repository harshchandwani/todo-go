package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo-app/config"
	"todo-app/routes"
	// Adjust the import paths according to your project structure
)

func main() {
	client := config.ConnectDB()
	defer client.Disconnect(nil)

	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
