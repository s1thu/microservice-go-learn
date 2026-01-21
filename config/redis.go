package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func ExampleClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     AppConfig.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		panic("❌ Failed to connect to Redis")
	} else {
		println("✅ Connected to Redis successfully")
	}

}
