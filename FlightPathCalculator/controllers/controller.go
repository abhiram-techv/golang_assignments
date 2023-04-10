package controllers

import (
	"encoding/json"
	"flight-route-calculator/services"
	"net/http"
)

func CalculateFlightPathHandler(w http.ResponseWriter, r *http.Request) {

	// Decode JSON request
	var flightRouteRequests [][]string
	err := json.NewDecoder(r.Body).Decode(&flightRouteRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate flight path
	response, err := services.CalculateFlightPath(flightRouteRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return flight path as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
