package database

import (
	"CRUD-REDIS/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func RedisConnect(config *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	fmt.Println("Successfully connected to Redis!")
	return client
}
