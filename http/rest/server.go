package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/configs"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/http/rest/handlers"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/db"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	config configs.Config
}

func NewServer() (*Server, error) {
	config, err := configs.NewParsedConfig()

	if err != nil {
		return nil, err
	}

	database, err := db.Connect(db.ConfigDB{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		Name:     config.Database.Name,
	})

	if err != nil {
		return nil, err
	}

	log, err := NewLogger()

	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	handlers.Register(router, log, database)

	server := Server{
		logger: log,
		config: config,
		router: router,
	}

	return &server, nil
}

func (s *Server) Run(ctx context.Context) error {
	cors := cors.New(cors.Options{
		AllowedMethods: []string{"GET"},
		AllowedOrigins: []string{"*"},
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.ServerPort),
		Handler: cors.Handler(s.router),
	}

	stopServer := make(chan os.Signal, 1)

	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	defer signal.Stop(stopServer)

	serverErrors := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s.logger.Printf("REST API listening on  %d", s.config.ServerPort)
		serverErrors <- server.ListenAndServe()
	}(&wg)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting REST API server %w", err)
	case <-stopServer:
		s.logger.Warn("server recieved STOP signal")

		err := server.Shutdown(ctx)

		if err != nil {
			return fmt.Errorf("graceful shutdown did not complete: %w", err)
		}

		wg.Wait()
		s.logger.Info("server was shutdown gracefully")
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
