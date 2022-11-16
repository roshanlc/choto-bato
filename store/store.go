package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

// StorageService is a wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// Time duration to keep links in memory
const CacheDuration = 6 * time.Hour

// InitializeStore setups a new redis client and returns a pointer
func InitializeStore() *StorageService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password is set
		DB:       0,  // use default DB
	})

	// Create a redis client instance
	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error init redis: %v", err))
	}

	fmt.Printf("Redis started successfully: pong msg = {%s}\n", pong)
	storeService.redisClient = rdb
	return storeService
}

// SaveURL saves a pair of short and original URLs in redis
func SaveURL(shortURL, originalURL string) error {

	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()

	return err
}

// RetrieveURL fetches longURL from redis storage
func RetrieveURL(shortURL string) (string, error) {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
