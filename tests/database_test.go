package tests

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"modernPokedex/internal/database"

	_ "github.com/lib/pq"
)

func TestPostgresTableExists(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Log("No .env file found")
	}

	if err := os.Setenv("DB_HOST", "localhost"); err != nil {
		t.Fatalf("Error setting DB_HOST environment variable: %v", err)
	}
	var db *sql.DB

	for range 5 {
		db, err = database.PostgresInit()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Waiting for database to start...")
	}

	if err != nil {
		t.Fatal("Error connecting to database")
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("Error closing database: %v", err)
		}
	}()

	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'pokemons'
		);`
	var exists bool
	err = db.QueryRowContext(context.Background(), query).Scan(&exists)
	if err != nil {
		t.Fatal("Error checking if table exists")
	}

	if !exists {
		t.Fatal("Table does not exist")
	} else {
		t.Log("Table exists")
	}
}
