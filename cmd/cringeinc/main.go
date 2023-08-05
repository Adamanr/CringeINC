package main

import (
	"cringeinc_server/internal/config"
	"cringeinc_server/internal/database/postgres"
	"cringeinc_server/internal/http-server/middleware"
	"github.com/go-chi/chi"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	cfg := config.Load()

	db, err := postgres.New(&cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	middleware.SetRouter(router, db)
	srv := http.Server{
		Addr:         cfg.HTTPServer.Address,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("error listen server!")
		return
	}
}
