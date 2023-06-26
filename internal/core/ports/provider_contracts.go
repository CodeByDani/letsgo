package ports

import "github.com/cyneptic/letsgo/internal/core/entities"

type FlightProviderContract interface {
	RequestFlights(source, destination, departure string) ([]entities.Flight, error)
	RequestFlight(id string) (entities.Flight, error)
}
