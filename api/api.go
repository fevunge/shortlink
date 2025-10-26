// Package api
package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func NewHandler(db *sql.DB) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Post("/api/short", HandlePost(db))
	r.Get("/{code}", HandleGet(db))
	return r
}

func HandlePost(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJSON(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
		}
		if _, err := url.Parse(body.URL); err != nil {
			SendJSON(w, Response{Error: "Invalid URL"}, http.StatusBadRequest)
		}
	})
}

func HandleGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func SendJSON(wr http.ResponseWriter, res Response, status int) {
	data, err := json.Marshal(res)
	if err != nil {
		logger.Error("JSON Marchal ERROR", zap.Any("error", err))
	}
	wr.WriteHeader(status)
	if _, err := wr.Write(data); err != nil {
		logger.Error("Writing DATA Buffer ERROR", zap.Any("error", err))
	}
}
