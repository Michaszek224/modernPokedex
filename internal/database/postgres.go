package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Pokemon struct {
	ID    int
	Name  string
	Type1 string
	Type2 string
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (db *Postgres) GetPokemonByID(id int) (*Pokemon, error) {
	var pokemon Pokemon
	err := db.db.QueryRow("SELECT * FROM pokemons WHERE id = $1", id).Scan(&pokemon.ID, &pokemon.Name, &pokemon.Type1, &pokemon.Type2)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (db *Postgres) GetPokemonByName(name string) (*Pokemon, error) {
	var pokemon Pokemon
	err := db.db.QueryRow("SELECT * FROM pokemons WHERE name = $1", name).Scan(&pokemon.ID, &pokemon.Name, &pokemon.Type1, &pokemon.Type2)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (db *Postgres) GetAllPokemon() ([]*Pokemon, error) {
	var pokemons []*Pokemon
	rows, err := db.db.Query("SELECT * FROM pokemons")
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var pokemon Pokemon
		err := rows.Scan(&pokemon.ID, &pokemon.Name, &pokemon.Type1, &pokemon.Type2)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, &pokemon)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pokemons, nil
}
