package main

import (
	"github.com/SergeyStrashko/architecture-lab2/server/plants"
)

func ComposeApiServer(port HttpPortNumber) (*PlantApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := plants.NewStore(db)
	httpHandlerFunc := plants.HttpHandler(store)
	plantApiServer := &PlantApiServer{
		Port:          port,
		PlantsHandler: httpHandlerFunc,
	}
	return plantApiServer, nil
}
