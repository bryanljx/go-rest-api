package database

import (
	"fmt"
	"log"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connect initialises a raw connection to the database specified by the given parameters.
func Connect(cfg *config.Config) (*sqlx.DB, error) {
	dsn := BuildDSN(cfg)
	db, err := sqlx.Connect("postgres", dsn)
	return db, err
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

	log.Print(dsn)
	return dsn
}
