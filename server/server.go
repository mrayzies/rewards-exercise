package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Address string `env:"SERVER_ADDRESS"`
	Port    string `env:"SERVER_PORT"`
}

type Server struct {
	*gin.Engine
	Address string
	Port    string
}

func New(c *Config) (*Server, error) {
	s := Server{}

	// save configuration results
	if c.Address != "" {
		s.Address = c.Address
	} else {
		s.Address = "0.0.0.0"
	}

	if c.Port != "" {
		s.Port = c.Port
	} else {
		s.Port = "8080"
	}

	// create a Gin router with default logger and recovery middleware
	s.Engine = gin.Default()

	// define routes
	s.setRoutes()

	return &s, nil
}

func (s Server) setRoutes() {
	s.Engine.GET("/transacion")
	s.Engine.POST("/points")
	s.Engine.GET("/points")
}

func (s Server) Run() error {
	addr := fmt.Sprintf("%s:%s", s.Address, s.Port)
	return s.Engine.Run(addr)
}
