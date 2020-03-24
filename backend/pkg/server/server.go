package server

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"keepboard/backend/pkg/api"
)

type Option struct {
	Port string
}

type Server struct {
	*echo.Echo
	opt Option
}

func New(opt Option) *Server {
	if opt.Port == "" {
		opt.Port = "8080"
	}

	return &Server{
		Echo: echo.New(),
		opt:  opt,
	}
}

func (s *Server) Run() error {
	s.setRoute()

	// TODO - env local 환경에서만 CORS allowed
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	endpoint := fmt.Sprintf(":%s", s.opt.Port)
	return s.Start(endpoint)
}

func (s *Server) setRoute() {
	h := api.New(func() ([]string, error) {
		return []string{"hello world", "this is sparta", "test"}, nil
	}, func(msg string) error {
		log.Println("input", msg)
		return nil
	})

	s.GET("/", h.HealthCheck)

	apiGroup := s.Group("/api")
	{
		apiGroup.GET("/comment/", h.GetComment)
		apiGroup.POST("/comment/", h.AddComment)
	}
}
