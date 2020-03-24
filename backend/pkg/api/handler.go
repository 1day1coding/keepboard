package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	getComments func() ([]string, error)
	addComment  func(string) error
}

func New(getComments func() ([]string, error), addComment func(string) error) *Handler {
	return &Handler{
		getComments: getComments,
		addComment:  addComment,
	}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetComment(c echo.Context) error {
	res, err := h.getComments()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) AddComment(c echo.Context) error {
	type aux struct {
		Comment string `json:"comment"`
	}
	var input aux
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.addComment(input.Comment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
