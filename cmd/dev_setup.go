package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func devServerSetup(connStr string, pool *pgxpool.Pool) error {
	ctx := context.Background()

	// --- AUTOMATED MIGRATION RUNNER ---
	migrator, err := migrate.New(
		"file://migrations/", // Update this to match your actual migration folder path!
		connStr,
	)
	if err != nil {
		return fmt.Errorf("Failed to initialize migration engine: %v", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to run up migrations: %v", err)
	}

	// --- AUTOMATED seed RUNNER ---
	seed, err := os.ReadFile("migrations/seed.sql")
	if err != nil {
		return fmt.Errorf("Trouble reading the seed file: %v", err)
	}

	_, err = pool.Exec(ctx, string(seed))
	if err != nil {
		return fmt.Errorf("Trouble seeding: %v", err)
	}

	return nil

}
