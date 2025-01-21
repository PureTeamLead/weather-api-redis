package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient = redisInit()

func hostEnv() string {
	_ = godotenv.Load()

	host := os.Getenv("REDIS_HOST")
	return host
}

func redisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     hostEnv(),
		Password: "",
		DB:       0,
		Protocol: 3,
	})

	return client
}

func CacheResponse(weatherLocation string, resp *Response) error {

	resp.CurrentConditions.Temp = math.Round(celsiusConverter(resp.CurrentConditions.Temp))

	respJSON, err := json.Marshal(*resp)
	if err != nil {
		return fmt.Errorf("error marshalling to json for redis db: %w", err)
	}

	if err = redisClient.Set(ctx, weatherLocation, respJSON, time.Hour).Err(); err != nil {
		return fmt.Errorf("error caching response: %v", err)
	}

	log.Println("Cached with key", weatherLocation)

	return nil
}

func ExistsInCache(key string) bool {
	_, err := redisClient.Get(ctx, key).Result()
	log.Println("checking if key exists in RedisDB")
	return err != redis.Nil
}

func GetCachedResponse(key string) (*Response, error) {
	var result *Response

	log.Println("Trying to get response from Redis")

	resultJSON, _ := redisClient.Get(ctx, key).Result()

	if err := json.Unmarshal([]byte(resultJSON), &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling from redis db: %w", err)
	}

	log.Println("Result for the key:", key)
	return result, nil
}
