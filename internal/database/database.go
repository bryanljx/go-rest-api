package database

import (
	"fmt"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Connect initialises a raw connection to the database specified by the given parameters.
func Connect(cfg *config.Config) (*sqlx.DB, error) {
	dsn := BuildDSN(cfg)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error connecting to database %s", cfg.DBName))
	}
	return db, nil
}

// BuildDSN builds the data source name that is used to connect to a database.
func BuildDSN(cfg *config.Config) string {
	dsn := ""
	if cfg.DBName != "" {
		dsn += fmt.Sprintf("dbname=%v", cfg.DBName)
	}
	if cfg.DBHostname != "" {
		dsn += fmt.Sprintf(" host=%v", cfg.DBHostname)
	}
	if cfg.DBPort != "" {
		dsn += fmt.Sprintf(" port=%v", cfg.DBPort)
	}
	if cfg.DBUser != "" {
		dsn += fmt.Sprintf(" user=%v", cfg.DBUser)
	}
	if cfg.DBPassword != "" {
		dsn += fmt.Sprintf(" password=%v", cfg.DBPassword)
	}
	if cfg.DBSSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%v", cfg.DBSSLMode)
	}
	return dsn
}
