package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/bryanljx/go-rest-api/internal/database"
	"github.com/bryanljx/go-rest-api/internal/lib/validator"
	"github.com/bryanljx/go-rest-api/internal/router"
	"github.com/bryanljx/go-rest-api/internal/service"

	"log/slog"
)

type Server struct {
	config  *config.Config
	service service.Service
}

func Init(logger *slog.Logger, config *config.Config) (*Server, error) {
	v := validator.New()
	if err := v.ValidateStruct(config); err != nil {
		logger.Error(fmt.Sprintf("Error: invalid config - %v", err))
		return nil, err
	}

	db, err := database.Connect(config)
	if err != nil {
		logger.Error(fmt.Sprintf("Error: unable to start db - %v", err))
		return nil, err
	}

	// Init DB first
	return &Server{
		config: config,
		service: service.Service{
			DB:        db,
			Logger:    logger,
			Validator: validator.New(),
		},
	}, nil
}

func (server *Server) Start() {
	l := server.service.Logger
	timeout := time.Minute
	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, os.Interrupt, os.Kill, syscall.SIGHUP)

	mux := router.SetUpRoutes(&server.service, server.config)
	s := http.Server{
		Handler:      *mux,
		Addr:         server.config.ServerPort,
		ReadTimeout:  time.Second * time.Duration(server.config.ServerTimeoutRead),
		WriteTimeout: time.Second * time.Duration(server.config.ServerTimeoutWrite),
		IdleTimeout:  time.Second * time.Duration(server.config.ServerTimeoutIdle),
	}

	go func() {
		l.Info(fmt.Sprintf("Server listening at %s", server.config.ServerPort))
		if err := s.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	// block until either OS signal, or server fatal error
	select {
	case err := <-errChan:
		l.Error(fmt.Sprintf("Server error: %v\n", err))
	case <-stopChan:
	}

	timeoutFunc := time.AfterFunc(timeout, func() {
		l.Error("server failed to shutdown within %d min, force quit")
		os.Exit(1)
	})
	defer timeoutFunc.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		l.Error(fmt.Sprintf("Error shutting down server: %v", err))
	}

	err = server.service.DB.Close()
	if err != nil {
		l.Error(fmt.Sprintf("Error shutting down server: %v", err))
	}

	l.Info("Server shutdown")
}
