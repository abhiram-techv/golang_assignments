package services

import (
	"errors"
	"flight-route-calculator/models"
)

var (
	Source      = "source"
	Destination = "destination"
)

func CalculateFlightPath(flights []models.Flight) (models.Flight, error) {
	AirportMap := make(map[string]string)

	//finding ports that appears only once in input slice
	for _, flight := range flights {
		if _, ok := AirportMap[flight.Source]; ok {
			AirportMap[flight.Source] = ""
		} else {
			AirportMap[flight.Source] = Source
		}
		if _, ok := AirportMap[flight.Destination]; ok {
			AirportMap[flight.Destination] = ""
		} else {
			AirportMap[flight.Destination] = Destination
		}

	}
	var startPort string
	var endPort string
	// checking the value of ports that appears only once. if more than one source and destination values exists , then input is invalid
	for key := range AirportMap {
		switch AirportMap[key] {
		case Source:
			if startPort != "" {
				return models.Flight{}, errors.New("invalid input, more than one starting ports found")
			}
			startPort = key
		case Destination:
			if endPort != "" {
				return models.Flight{}, errors.New("invalid input, more than one ending ports found")
			}
			endPort = key
		}
	}

	if startPort == "" || endPort == "" {
		return models.Flight{}, errors.New("invalid input, starting or ending ports not found")
	}

	// Return the starting and ending ports.
	return models.Flight{
		Source:      startPort,
		Destination: endPort,
	}, nil
}
