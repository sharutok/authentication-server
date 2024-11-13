package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectToRedis() *redis.Client {
	i, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Println(err, "error in db redis")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       i,
	})

	return rdb
}
