package main

import (
	"crudinmemory/api"
	"log/slog"
	"net/http"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {

	handler := api.NewHandler()
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
