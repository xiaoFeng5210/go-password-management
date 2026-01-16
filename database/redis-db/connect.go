package redisDB

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client = nil

func Connect() {
	if Client != nil {
		return
	}

	// 加载 .env 文件（如果存在）
	_ = godotenv.Load()

	// 从环境变量读取配置，如果没有则使用默认值
	addr := getEnvOrDefault("REDIS_ADDR", "localhost:6379")
	username := getEnvOrDefault("REDIS_USERNAME", "")
	password := getEnvOrDefault("REDIS_PASSWORD", "")
	dbStr := getEnvOrDefault("REDIS_DB", "0")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Username: username,
		Password: password,
	})

	// 能ping成功才说明连接成功
	if err := Client.Ping(context.Background()).Err(); err != nil {
		fmt.Println("connect to redis failed", err)
		os.Exit(1)
	}
}

// getEnvOrDefault 获取环境变量，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
