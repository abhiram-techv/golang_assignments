package controllers

import (
	"encoding/json"
	"flight-route-calculator/models"
	"flight-route-calculator/services"
	"net/http"
)

func CalculateFlightPathHandler(w http.ResponseWriter, r *http.Request) {

	// Decode JSON request into a slice of Flight structs
	var flights []models.Flight
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate flight path
	flightPath, err := services.CalculateFlightPath(flights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return flight path as JSON response
	response := map[string]models.Flight{
		"flightPath": flightPath,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
