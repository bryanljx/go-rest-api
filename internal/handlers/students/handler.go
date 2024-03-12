package students

import (
	"log/slog"

	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/repositories"
	"github.com/bryanljx/go-rest-api/internal/service"
)

type StudentHandler struct {
	repo      repositories.StudentRepository
	logger    *slog.Logger
	Validator validator.Validator
}

func New(repo repositories.StudentRepository, s *service.Service) StudentHandler {
	return StudentHandler{
		repo:      repo,
		logger:    s.Logger,
		Validator: s.Validator,
	}
}
