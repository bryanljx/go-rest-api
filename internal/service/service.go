package service

import (
	"log/slog"

	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB        *sqlx.DB
	Logger    *slog.Logger
	Validator validator.Validator
}
