package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Ctx = context.Background()

const ViewCount = "viewcount"

func ConnectRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		// Addr: "0.0.0.0:7001",
		// docker-compose 裡面有自己的 dns，api 如果在 docker 裡面  不能用127.0.0.1
		Addr: "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check if redis db is connected
	_, err := RedisDB.Ping(Ctx).Result()
	if err == nil {
		log.Println("Redis connected.")
	} else {
		panic("Redis connection failed: " + err.Error())
	}

	// set viewcount as initial 0
	err = RedisDB.Set(Ctx, ViewCount, 0, 0).Err()
	if err != nil {
		panic("Redis cannot initialize viewcount: " + err.Error())
	}
}
