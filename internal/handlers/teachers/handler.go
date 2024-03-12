package teachers

import (
	"log/slog"

	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/repositories"
	"github.com/bryanljx/go-rest-api/internal/service"
)

type TeacherHandler struct {
	repo      repositories.TeacherRepository
	logger    *slog.Logger
	Validator validator.Validator
}

func New(repo repositories.TeacherRepository, s *service.Service) TeacherHandler {
	return TeacherHandler{
		repo:      repo,
		logger:    s.Logger,
		Validator: s.Validator,
	}
}
