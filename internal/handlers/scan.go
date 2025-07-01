package handlers

import (
	"github.com/amir-mirjalili/ip-scanner/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScanHandler struct {
	ScanService *services.ScanService
	Validator   *validator.Validate
}

type ScanRequest struct {
	CIDR string `json:"cidr" validate:"required,cidr"`
}

func (h *ScanHandler) StartScan(c echo.Context) error {
	var req ScanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "validation failed", "details": err.Error()})
	}

	scan, err := h.ScanService.RunAndSaveScan(req.CIDR)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, scan)
}
