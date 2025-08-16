package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/service"
	"github.com/teryble09/subscription_service/storage/postgres"
	"github.com/teryble09/subscription_service/xmiddleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := postgres.NewStorage(databaseURL, logger)
	if err != nil {
		logger.Error("Can not connect to the database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	srv := &service.SubscriptionService{
		Logger:  logger,
		Storage: db,
	}

	server, err := api.NewServer(srv)
	if err != nil {
		logger.Error("Failed to create server", "error", err.Error())
		os.Exit(1)
	}

	loggingMiddleware := xmiddleware.NewLoggingMiddleware(logger)

	logger.Info("Starting server")

	err = http.ListenAndServe(":8080", loggingMiddleware(server))

	if err != nil {
		logger.Error("Unable to start server",
			slog.String("error", err.Error()),
		)
	}
}
