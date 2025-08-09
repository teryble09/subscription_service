package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func (s *Storage) runMigrations() error {
	p, err := goose.NewProvider(goose.DialectPostgres, s.db.DB, migrations)
	if err != nil {
		return fmt.Errorf("migration provider: %w", err)
	}

	_, err = p.Up(context.Background())
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
