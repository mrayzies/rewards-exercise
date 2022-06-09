package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/mrayzies/reward-exercise/server"
)

func main() {
	c := &server.Config{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("failed to parse server configuration: %v", err)
		return
	}

	s, err := server.New(c)
	if err != nil {
		log.Fatalf("failed to create HTTP server: %v", err)
		return
	}

	if err := s.Run(); err != nil {
		log.Fatalf("fatal error running HTTP server: %v", err)
		return
	}
}
