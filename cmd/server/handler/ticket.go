package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetTicketsByCountry(ctx *gin.Context) {

	destination := ctx.Param("dest")

	tickets, err := s.service.GetTotalTickets(ctx, destination)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(tickets) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontro tickets para ese destino"})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (s *Service) AverageDestination(c *gin.Context) {

	//destination := c.Param("dest")

	avg, err := s.service.AverageDestination()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.JSON(200, avg)
}
