package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/comment/:str", addComment)

	e.Logger.Fatal(e.Start(":8888")) // localhost:1323
}

func addComment(c echo.Context) error {
	str := c.Param("str")
	return c.String(http.StatusOK, str)
}
