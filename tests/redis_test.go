package tests

import (
	"context"
	"modernPokedex/internal/database"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func setupRedis(t *testing.T) *redis.Client {
	t.Helper()

	if err := os.Setenv("REDIS_ADDR", "localhost:6379"); err != nil {
		t.Fatalf("failed to set REDIS_ADDR: %v", err)
	}

	var (
		rdb *redis.Client
		err error
	)

	for range 5 {
		rdb, err = database.RedisInit()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		t.Fatalf("failed to connect to redis: %v", err)
	}

	return rdb
}

func TestRedisSetGet(t *testing.T) {
	rdb := setupRedis(t)
	defer func() {
		if err := rdb.Close(); err != nil {
			t.Fatalf("failed to close redis connection: %v", err)
		}
	}()

	ctx := context.Background()
	key := "test-key"
	value := "Nidorina"

	_, err := rdb.Get(ctx, key).Result()
	if err != redis.Nil {
		t.Fatalf("should be cache miss, got %v", err)
	}
	t.Log("cache miss- thats good")

	err = rdb.Set(ctx, key, value, 5*time.Minute).Err()
	if err != nil {
		t.Fatalf("failed to set cache: %v", err)
	}

	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("failed to get cache: %v", err)
	}

	if result != value {
		t.Fatalf("should be %v, got %v", value, result)
	}
	t.Log("cache hit- thats good")
	t.Logf("Cached value: %v", result)
}

func TestRedisMultipleKeys(t *testing.T) {
	rdb := setupRedis(t)
	defer func() {
		if err := rdb.Close(); err != nil {
			t.Fatalf("failed to close redis connection: %v", err)
		}
	}()

	ctx := context.Background()
	keys := map[string]string{
		"test-key-1": "Nidorina",
		"test-key-2": "Ekans",
		"test-key-3": "Mew",
	}

	for key, value := range keys {
		err := rdb.Set(ctx, key, value, 5*time.Minute).Err()
		if err != nil {
			t.Fatalf("failed to set cache: %v", err)
		}
	}

	for key, valueExpected := range keys {
		result, err := rdb.Get(ctx, key).Result()
		if err != nil {
			t.Fatalf("failed to get cache for key %v: %v", key, err)
		}

		if result != valueExpected {
			t.Fatalf("should be %v, got %v", valueExpected, result)
		}
		t.Logf("Good [key,value]: [%v,%v]", key, result)
	}
}
