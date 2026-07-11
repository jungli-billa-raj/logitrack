package logitrack

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	// These underscores force Go to register the drivers behind the scenes
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	connStr := "postgres://postgres:MySecretPassword123@localhost:5432/logitrack?sslmode=disable"
	var err error
	testPool, err = ConnectDB(ctx, connStr)
	if err != nil {
		log.Fatalf("message: %v", err) // If DB is not connected then stop the test execution
	}
	defer testPool.Close()

	// --- AUTOMATED MIGRATION RUNNER ---
	migrator, err := migrate.New(
		"file://migrations/", // Update this to match your actual migration folder path!
		connStr,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migration engine: %v", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run up migrations: %v", err)
	}

	// --- AUTOMATED seed RUNNER ---
	seed, err := os.ReadFile("migrations/seed.sql")
	if err != nil {
		log.Fatalf("Trouble reading the seed file: %v", err)
	}

	_, err = testPool.Exec(ctx, string(seed))
	if err != nil {
		log.Fatalf("Trouble seeding: %v", err)
	}

	errorCode := m.Run()

	if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run up migrations: %v", err)
	}

	os.Exit(errorCode)
}
