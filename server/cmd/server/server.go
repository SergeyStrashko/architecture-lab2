package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SergeyStrashko/architecture-lab2/server/plants"
)

type HttpPortNumber int

type PlantApiServer struct {
	Port HttpPortNumber

	PlantsHandler plants.HttpHandlerFunc

	server *http.Server
}

func (s *PlantApiServer) Start() error {
	if s.PlantsHandler == nil {
		return fmt.Errorf("Plants HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/sendData", s.PlantsHandler)
	handler.HandleFunc("/getData", s.PlantsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *PlantsApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
