package tests

import (
	"database/sql"
	"modernPokedex/internal/database"
	"os"
	"testing"
	"time"
)

func setupPostgres(t *testing.T) *sql.DB {
	t.Helper()
	if err := os.Setenv("DB_HOST", "localhost"); err != nil {
		t.Fatalf("failed to set DB_HOST: %v", err)
	}

	var (
		db  *sql.DB
		err error
	)

	for range 5 {
		db, err = database.PostgresInit()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		t.Fatalf("failed to connect to postgres: %v", err)
	}

	return db
}

func TestGetPokemonByID(t *testing.T) {
	db := setupPostgres(t)

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("Error closing database: %v", err)
		}
	}()
	pg := database.NewPostgres(db)

	pokemon, err := pg.GetPokemonByID(1)
	if err != nil {
		t.Fatalf("failed to get pokemon: %v", err)
	}
	if pokemon.ID != 1 {
		t.Fatalf("expected pokemon id to be 1, got %d", pokemon.ID)
	}

	if pokemon.Name != "Bulbasaur" {
		t.Fatalf("expected pokemon name to be Bulbasaur, got %s", pokemon.Name)
	}
	t.Logf("got pokemon: %+v", pokemon)
}

func TestGetPokemonByName(t *testing.T) {
	db := setupPostgres(t)

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("Error closing database: %v", err)
		}
	}()

	pg := database.NewPostgres(db)

	pokemon, err := pg.GetPokemonByName("Mewtwo")
	if err != nil {
		t.Fatalf("failed to get pokemon: %v", err)
	}
	if pokemon.ID != 150 {
		t.Fatalf("expected pokemon id to be 150, got %d", pokemon.ID)
	}

	if pokemon.Name != "Mewtwo" {
		t.Fatalf("expected pokemon name to be Mewtwo, got %s", pokemon.Name)
	}
	t.Logf("got pokemon: %+v", pokemon)
}

func TestGetAllPokemon(t *testing.T) {
	db := setupPostgres(t)

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("Error closing database: %v", err)
		}
	}()

	pg := database.NewPostgres(db)

	pokemons, err := pg.GetAllPokemon()
	if err != nil {
		t.Fatalf("failed to get pokemons: %v", err)
	}

	if len(pokemons) == 0 {
		t.Fatalf("expected at least one pokemon, got %d", len(pokemons))
	}
}
