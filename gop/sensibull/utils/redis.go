// utils/redis.go
package utils

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient *redis.Client
)

// InitRedis initializes the Redis client
func InitRedis(addr, password string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// GetFromCache retrieves data from Redis cache
func GetFromCache(key string) (string, error) {
	//ctx := context.Background()
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// SetToCache sets data in Redis cache with an expiration time
func SetToCache(key, value string, expiration time.Duration) error {
	//ctx := context.Background()
	return redisClient.Set(key, value, expiration).Err()
}
