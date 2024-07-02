package api

import (
	"github.com/majorchork/book-keeper/internal/ports"
)

type HTTPHandler struct {
	Service ports.ComputerService
}

func NewHTTPHandler(service ports.ComputerService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}
