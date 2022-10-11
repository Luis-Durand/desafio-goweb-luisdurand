package tickets

import (
	"desafio-goweb-luisdurand/internal/domain"
)

type Service interface {
	GetAll(destination string) ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverage(destination string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}

func (s *service) GetAll(destination string) ([]domain.Ticket, error) {

	tickets, err := s.repository.GetAll(destination)
	if err != nil {
		return nil, err
	}

	return tickets, nil

}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	ticket, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}

	return ticket, nil

}

func (s *service) GetAverage(destination string) (string, error) {

	ticket, err := s.repository.GetAverage(destination)
	if err != nil {
		return "", err
	}

	return ticket, nil

}
