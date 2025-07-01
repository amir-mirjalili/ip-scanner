package handlers

import (
	"net/http"
	"strconv"

	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"github.com/amir-mirjalili/ip-scanner/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AssetHandler struct {
	AssetService *services.AssetService
	Validator    *validator.Validate
}

func (h *AssetHandler) CreateAsset(c echo.Context) error {
	var asset models.Asset
	if err := c.Bind(&asset); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := h.Validator.Struct(asset); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "validation failed", "details": err.Error()})
	}

	if err := h.AssetService.Create(&asset); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, asset)
}

func (h *AssetHandler) GetAsset(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid asset id"})
	}

	asset, err := h.AssetService.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if asset == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "asset not found"})
	}
	return c.JSON(http.StatusOK, asset)
}

func (h *AssetHandler) UpdateAsset(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid asset id"})
	}

	var asset models.Asset
	if err := c.Bind(&asset); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	asset.ID = uint(id)

	if err := h.Validator.Struct(asset); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "validation failed", "details": err.Error()})
	}

	if err := h.AssetService.Update(&asset); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, asset)
}

func (h *AssetHandler) DeleteAsset(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid asset id"})
	}

	if err := h.AssetService.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *AssetHandler) ListAssets(c echo.Context) error {
	assets, err := h.AssetService.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, assets)
}
