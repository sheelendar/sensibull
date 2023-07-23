// utils/redis.go
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient *redis.Client
)

// InitRedis initializes the Redis client
// docker run -d --name my-redis -p 6379:6379 redis:latest
func InitRedis(addr, password string, db int) {
	fmt.Println("redis connection init..")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	res, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis err when ping ", err)
	}
	fmt.Println("redis ping result:", res)
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

// GetSymbolToTokenMapFromCache retrieves SymbolToTokenMap from Redis cache
func GetSymbolToTokenMapFromCache(key string) (map[string]int, error) {
	val, err := redisClient.Get(key).Bytes()
	if err != nil {
		// return error
	}
	var symbolToTokenCache map[string]int
	err = json.Unmarshal(val, &symbolToTokenCache)

	if err != nil {
		return nil, err
	}
	return symbolToTokenCache, err
}

// SetStringToCache sets data in Redis cache with an expiration time
func SetStringToCache(key, value string, expiration time.Duration) error {
	return redisClient.Set(key, value, expiration).Err()
}

// SaveObjectsInRedis is used to save a map in Redis
func SaveObjectsInRedis(key string, data interface{}, expiryTime time.Duration) error {
	// Convert the map to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Save the JSON data in Redis with the specified key
	err = redisClient.Set(key, jsonData, expiryTime).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetObjectFromRedis Function fetch value from redis and unmarshal into obj.
func GetObjectFromRedis(key string, data interface{}) error {
	// Get the JSON data from Redis
	jsonData, err := redisClient.Get(key).Bytes()
	if err != nil {
		return err
	}

	// Convert the JSON data back to a obj
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}
	return nil
}
