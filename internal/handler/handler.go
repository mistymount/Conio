package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mistymount/Conio/internal/service"
)

type handler struct {
	*service.Service
}

// New creates an http handler with predefined routing
func New(s *service.Service) http.Handler {
	h := &handler{s}
	e := echo.New()

	api := e.Group("/api")
	api.POST("/login", h.login)
	api.POST("/users", h.createUser)

	return e
}
