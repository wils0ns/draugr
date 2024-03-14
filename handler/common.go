package handler

import (
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"time"
)

func MakeHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error(err.Error(), "path", r.URL.Path)
		}
	}
}

func Logger() *httplog.Logger {
	return httplog.NewLogger("default", httplog.Options{
		// TODO customize if in prod
		JSON:           false,
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		QuietDownRoutes: []string{
			"/livez",
			"/readyz",
		},
		QuietDownPeriod: 10 * time.Second,
	})
}
