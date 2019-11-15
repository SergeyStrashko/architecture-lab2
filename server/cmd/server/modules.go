package main

import (
	"github.com/SergeyStrashko/architecture-lab2/server/plants"
	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*PlantApiServer, error) {
	wire.Build(
		NewDbConnection,
		plants.Providers,
		wire.Struct(new(PlantApiServer), "Port", "PlantsHandler"),
	return nil, nil
}