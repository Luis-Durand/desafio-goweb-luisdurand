package tickets

import (
	"desafio-goweb-luisdurand/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll(destination string) ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverage(destination string) (string, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(destination string) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAverage(destination string) (string, error) {

	var ticketsDest []domain.Ticket
	var result float64

	if len(r.db) == 0 {
		return "", fmt.Errorf("empty list of tickets")
	}
	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}
	result = (float64(len(ticketsDest)) * float64(100)) / float64(len(r.db))
	return fmt.Sprintf("El promedio es %.1f", result), nil
}
