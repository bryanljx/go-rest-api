package notifications

import (
	"log/slog"

	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/repositories"
	"github.com/bryanljx/go-rest-api/internal/service"
)

type NotificationHandler struct {
	studentRepo repositories.StudentRepository
	teacherRepo repositories.TeacherRepository
	logger      *slog.Logger
	Validator   validator.Validator
}

func New(studentRepo repositories.StudentRepository, teacherRepo repositories.TeacherRepository, s *service.Service) NotificationHandler {
	return NotificationHandler{
		studentRepo: studentRepo,
		teacherRepo: teacherRepo,
		logger:      s.Logger,
		Validator:   s.Validator,
	}
}
