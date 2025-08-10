package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/service"
	"github.com/teryble09/subscription_service/xmiddleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	srv := &service.SubscriptionService{
		Logger: logger,
	}

	server, err := api.NewServer(srv)
	if err != nil {
		logger.Error("Failed to create server", "error", err.Error())
		os.Exit(1)
	}

	loggingMiddleware := xmiddleware.NewLoggingMiddleware(logger)

	http.ListenAndServe(":8080", loggingMiddleware(server))
}
