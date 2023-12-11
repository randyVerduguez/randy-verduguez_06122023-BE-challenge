package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	toDoRepo "github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/repository"
	toDoService "github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/service"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	toDoService toDoService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:      lg,
		toDoService: toDoService.NewService(toDoRepo.NewRepository(db)),
	}
}
