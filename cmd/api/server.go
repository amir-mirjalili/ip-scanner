package api

import (
	"github.com/amir-mirjalili/ip-scanner/internal/adapter"
	"github.com/amir-mirjalili/ip-scanner/internal/db"
	"github.com/amir-mirjalili/ip-scanner/internal/handlers"
	"github.com/amir-mirjalili/ip-scanner/internal/repositories"
	"github.com/amir-mirjalili/ip-scanner/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	App         *echo.Echo
	ScanHandler *handlers.ScanHandler
}

func NewServer(database *db.Database) *Server {
	e := echo.New()

	validate := validator.New()

	scanRepo := repositories.NewGormScanRepository(database.DB)
	assetRepo := repositories.NewAssetGormRepository(database.DB)
	//assetService := services.NewAssetService(assetRepo)
	assetAdapter := adapter.NewAssetAdapter(assetRepo)
	scanService := services.NewScanService(scanRepo, assetAdapter)

	scanHandler := &handlers.ScanHandler{ScanService: scanService, Validator: validate}

	s := &Server{
		App:         e,
		ScanHandler: scanHandler,
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.App.GET("/", s.healthCheck)
	s.App.POST("/scan", s.ScanHandler.StartScan)
	s.App.GET("/scan/:id", s.ScanHandler.GetScan)
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "API is running 🚀",
	})
}
