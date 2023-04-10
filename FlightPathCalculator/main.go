package main

import (
	"flight-route-calculator/controllers"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	http.HandleFunc("/calculate", controllers.CalculateFlightPathHandler)
	log.Printf("Server is running and listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
