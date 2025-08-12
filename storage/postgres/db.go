package postgres

import (
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB

	// Подготовленные запросы
	stmtInsertSubsription *sqlx.NamedStmt
	stmtListSubsriptions  *sqlx.NamedStmt
}

func NewStorage(databaseURL string, logger *slog.Logger) (*Storage, error) {
	s := &Storage{}

	db, err := sqlx.Connect("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}

	logger.Info("Connected to the database")
	s.db = db

	if err = s.runMigrations(); err != nil {
		logger.Error("Could not run migrations",
			slog.String("error", err.Error()),
		)
		return nil, fmt.Errorf("storage: %w", err)
	}
	logger.Info("Succesfully aplied migrations")

	if err = s.prepareStatements(); err != nil {
		logger.Error("Could not prepare",
			slog.String("error", err.Error()),
		)
	}
	logger.Info("Succesfully prepared statements")

	return s, nil
}
