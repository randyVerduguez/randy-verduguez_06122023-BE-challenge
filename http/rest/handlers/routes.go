package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)

	r.Use(handler.MiddlewareLogger())
	r.HandleFunc("/weather/current", handler.GetCurrentWeather()).Methods(http.MethodGet)
	r.HandleFunc("/weather/welcome", handler.Test()).Methods(http.MethodGet)
}
