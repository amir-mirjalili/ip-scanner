package api

import (
	"github.com/amir-mirjalili/ip-scanner/internal/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	App      *echo.Echo
	Database *db.Database
}

func NewServer(database *db.Database) *Server {
	e := echo.New()

	s := &Server{
		App:      e,
		Database: database,
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.App.GET("/", s.healthCheck)
	// add more: s.App.GET("/assets", s.listAssets) etc.
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "API is running 🚀",
	})
}
