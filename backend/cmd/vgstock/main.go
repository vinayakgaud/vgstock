package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/vinayakgaud/vgstock/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("Failed to load config, using defaults", "error", err)
	}
	port := cfg.Port
	if port == "" {
		port = "8081"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthRouteHandler)

	slog.Info("vgstock server starting", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("Server exited", "error", err)
		os.Exit(1)
	}
}

func healthRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
