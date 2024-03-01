package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"upstash-redis-local/internal"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Version = "development"

func main() {
	redisAddr := os.Getenv("REDIS_URL")
	apiToken := os.Getenv("API_TOKEN")
	port := os.Getenv("PORT")

	addr := fmt.Sprintf("0.0.0.0:%s", port)

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := config.Build()

	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer logger.Sync()

	server := internal.Server{Address: addr, APIToken: apiToken, RedisConn: connectToRedis(redisAddr), Logger: logger}

	defer server.Serve()
}

func connectToRedis(addr string) *redis.Client {
	opts, err := redis.ParseURL(addr)

	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return redis.NewClient(opts)
}
