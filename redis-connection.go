package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "|*15kVpmmw`[3^}^:O8B&&v/L>QE+7OY7c9?%`VRsAq{,Cyxw{",
		DB:       0,
	})

	return rdb
}
