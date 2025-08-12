package postgres

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func (s *Storage) runMigrations() error {
	migrationsFS, err := fs.Sub(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("embedding migratrions: %w", err)
	}
	p, err := goose.NewProvider(goose.DialectPostgres, s.db.DB, migrationsFS)
	if err != nil {
		return fmt.Errorf("migration provider: %w", err)
	}

	_, err = p.Up(context.Background())
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
