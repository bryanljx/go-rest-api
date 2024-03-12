package mocks

import (
	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/monitoring"
	"github.com/bryanljx/go-rest-api/internal/service"
)

func MockService() service.Service {
	return service.Service{
		Logger:    monitoring.NewLogger("dev"),
		Validator: validator.New(),
	}
}
