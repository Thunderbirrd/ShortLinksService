package handler

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/service"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateShortUrl(c echo.Context) error {
	input := new(models.UrlObject)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	shortUrl, err := h.service.CreateShortUrl(*input)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"generated_url": shortUrl,
	})
}
