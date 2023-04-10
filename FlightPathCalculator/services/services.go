package services

import (
	"errors"
)

var (
	Source      = "source"
	Destination = "destination"
)

func CalculateFlightPath(flightRoutes [][]string) ([]string, error) {

	AirportMap := make(map[string]string)

	//finding ports that appears only once in input request
	for _, flightRoute := range flightRoutes {
		source, destination := flightRoute[0], flightRoute[1]
		if _, ok := AirportMap[source]; ok {
			AirportMap[source] = ""
		} else {
			AirportMap[source] = Source
		}
		if _, ok := AirportMap[destination]; ok {
			AirportMap[destination] = ""
		} else {
			AirportMap[destination] = Destination
		}
	}

	var startPort string
	var endPort string
	// checking the ports that appears only once. if more than one source and destination port exists, then input is invalid
	for key := range AirportMap {
		switch AirportMap[key] {
		case Source:
			if startPort != "" {
				return nil, errors.New("invalid input, more than one starting ports found")
			}
			startPort = key
		case Destination:
			if endPort != "" {
				return nil, errors.New("invalid input, more than one ending ports found")
			}
			endPort = key
		}
	}

	if startPort == "" || endPort == "" {
		return nil, errors.New("invalid input, starting or ending ports not found")
	}

	// Return the starting and ending ports.
	return []string{startPort, endPort}, nil
}
