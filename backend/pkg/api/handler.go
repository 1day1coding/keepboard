package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	getComment func() ([]string, error)
}

func New(getComment func() ([]string, error)) *Handler {
	return &Handler{getComment: getComment}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetComments(c echo.Context) error {
	res, err := h.getComment()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}
