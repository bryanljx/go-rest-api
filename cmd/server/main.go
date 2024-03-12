package main

import (
	"fmt"
	"log"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/bryanljx/go-rest-api/internal/monitoring"
	"github.com/bryanljx/go-rest-api/internal/server"
)

//	@title			Go Rest Api
//	@version		1.0
//	@description	A simple and limited Golang backend API created for OneCV take home assignment
//	@BasePath		/api
func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config - %v", err)
	}

	l := monitoring.NewLogger(config.Env)

	server, err := server.Init(l, config)
	if err != nil {
		l.Error(fmt.Sprintf("Error initialising server - %v", err))
	}

	server.Start()
}
