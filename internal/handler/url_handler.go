package handler

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/service"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
)

type Handler struct {
	port    string
	service *service.Service
}

func NewHandler(service *service.Service, port string) *Handler {
	return &Handler{service: service, port: port}
}

func (h *Handler) CreateShortUrl(c echo.Context) error {
	input := new(models.UrlObject)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	match, err := regexp.MatchString("^https?://", input.LongUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "not a URL")
	}

	shortUrl, err := h.service.CreateShortUrl(*input)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"generated_url": "http://localhost" + h.port + "/" + shortUrl,
	})
}

func (h *Handler) GetLongUrl(c echo.Context) error {
	shortUrl := c.Param("shortUrl")

	longUrl, err := h.service.GetLongUrlByShortUrl(shortUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if longUrl != "" {
		return c.JSON(http.StatusOK, echo.Map{
			"original_url": longUrl,
		})
	} else {
		return c.JSON(http.StatusBadRequest, "No short url for this long url")
	}
}

func (h *Handler) RedirectToLongUrl(c echo.Context) error {
	shortUrl := c.Param("shortUrl")

	longUrl, err := h.service.GetLongUrlByShortUrl(shortUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if longUrl != "" {
		return c.Redirect(http.StatusPermanentRedirect, longUrl)
	} else {
		return c.JSON(http.StatusBadRequest, "No short url for this long url")
	}
}
