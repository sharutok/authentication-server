package main

import "log"

func AddField(hashKey string, key string, value string) {
	rdb := ConnectToRedis()
	err := rdb.HSet(ctx, hashKey, key, value).Err()
	if err != nil {
		log.Fatalf("Could not set field in Redis hash: %v", err)
	}
}

func DeleteField(hashKey string, key string) {
	rdb := ConnectToRedis()
	err := rdb.HDel(ctx, hashKey, key).Err()
	if err != nil {
		log.Fatalf("Could not delete field from Redis hash: %v", err)
	}
}
