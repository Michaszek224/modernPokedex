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
		t.Fatal("Error loading .env file")
	}

	os.Setenv("DB_HOST", "localhost")
	var db *sql.DB

	for _ = range 5 {
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

	defer db.Close()

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
