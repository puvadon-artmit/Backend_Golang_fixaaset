package database

import "github.com/redis/go-redis/v9"

func Redis_cache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}
