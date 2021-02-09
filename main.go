package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/resideo/testing-service/api"
	"github.com/resideo/testing-service/internal/config"
)

// Generating OpenAPI3 wrappers
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o api/openapi-server.generated.go openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o api/openapi-types.generated.go openapi.yaml

func configure() *echo.Echo {
	log.Output(1, "resideo - testing-service")
	config.ReadConfig()
	flag.Parse()
	swagger := api.MakeSwagger()
	apiImpl := new(api.API)

	// This is how you set up a basic Echo router
	e := echo.New()

	// Log all requests
	e.Use(echomiddleware.Logger())

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our implementation above as the handler for the interface
	api.RegisterHandlers(e, apiImpl)

	return e
}

func main() {
	// Set up echo and load files
	e := configure()

	var port = flag.Int("port", 8080, "Port for HTTP server")

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
