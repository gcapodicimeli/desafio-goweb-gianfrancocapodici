package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, dest string) ([]domain.Ticket, error)
	AverageDestination() (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, dest string) (tickets []domain.Ticket, err error) {
	tick, err := s.repository.GetAll(ctx)
	if err != nil {
		return
	}

	for _, t := range tick {
		if t.Country == dest {
			tickets = append(tickets, t)
		}
	}

	return
}

func (s *service) AverageDestination() (avg float64, err error) {

	return
}
