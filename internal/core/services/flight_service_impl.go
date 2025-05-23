package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/sourcegraph/conc/iter"
)

type FlightService struct {
	repo secondary.FlightRepository
}

func NewFlightService(repo secondary.FlightRepository) primary.FlightService {
	return &FlightService{repo: repo}
}

func (s *FlightService) FindById(request *entities.GetFlightByIdRequest) (*entities.Flight, error) {
	return s.repo.FindById(request.ID)
}

func (s *FlightService) FindByFlightNumber(flightNumber string) (*entities.Flight, error) {
	return s.repo.FindByFlightNumber(flightNumber)
}

func (s *FlightService) FindAllActive(request *entities.GetAllFlightsRequest) ([]entities.Flight, error) {
	return s.repo.FindAllActive(request)
}

func (s *FlightService) FindAll(request *entities.GetAllFlightsRequest) ([]entities.Flight, error) {
	return s.repo.FindAll(request)
}

func (s *FlightService) Create(request *entities.CreateFlightRequest) error {
	flight := mapCreateFlightRequestToFlight(request)
	return s.repo.Create(&flight)
}

func (s *FlightService) CreateAll(requests []entities.CreateFlightRequest) error {
	flights := iter.Map(requests, mapCreateFlightRequestToFlight)
	return s.repo.CreateAll(flights)
}

func (s *FlightService) FindSeatsByFlightId(request *entities.GetSeatsByFlightIdRequest) ([]int, error) {
	return s.repo.FindSeatsByFlightId(request.ID)
}

func mapCreateFlightRequestToFlight(request *entities.CreateFlightRequest) entities.Flight {
	return entities.Flight{
		FlightNumber:          request.FlightNumber,
		DepartureAirport:      request.DepartureAirport,
		DestinationAirport:    request.DestinationAirport,
		DepartureDatetime:     request.DepartureDatetime,
		ArrivalDatetime:       request.ArrivalDatetime,
		DepartureGateNumber:   request.DepartureGateNumber,
		DestinationGateNumber: request.DestinationGateNumber,
		PlaneRegistration:     request.PlaneRegistration,
		Price:                 request.Price,
		Status:                "scheduled",
	}
}
