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
	echo := echo.New()

	api := echo.Group("/api")
	api.POST("/login", nil)
	api.POST("/users", nil)

	return echo
}
