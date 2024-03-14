package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sentryeye/handler"
)

func main() {
	// Config
	port := ":1666"
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(handler.Logger()))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/livez"))

	r.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	slog.Info("Starting SentryEye", "port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
