package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		type1 TEXT NOT NULL,
		type2 TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func RedisInit() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("error pinging redis: %v", err)
	}
	fmt.Println("Redis is up")
	return client, nil

}
