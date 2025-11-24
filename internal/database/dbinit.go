package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func PostgresInit() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	psqlLogin := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlLogin)
	if err != nil {
		return nil, fmt.Errorf("error opening sql: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging sql: %v", err)
	}

	err = createTable(db)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	} else {
		fmt.Println("Table exists")
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS pokemons (
		id INT PRIMARY KEY,
		name TEXT NOT NULL,
		types TEXT[] NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func redisInit() {
	// TODO: implement
}
