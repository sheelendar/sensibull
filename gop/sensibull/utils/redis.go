// utils/redis.go
package utils

import (
	"encoding/json"
	"fmt"
	"gop/sensibull/consts"
	"gop/sensibull/dao"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

// InitRedis initializes the Redis client
// docker run -d --name my-redis -p 6379:6379 redis:latest
func InitRedis(addr, password string, db int) *redis.Client {
	fmt.Println("redis connection init..")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	res, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis err when ping ", err)
		panic("not able to connect with db")
	}
	fmt.Println("redis ping result:", res)
	return redisClient
}

// GetStringFromCache retrieves string data from Redis cache
func GetStringFromCache(key string) (string, error) {
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
		return nil, err
	}
	var symbolToTokenCache map[string]int
	err = json.Unmarshal(val, &symbolToTokenCache)

	if err != nil {
		return nil, err
	}
	return symbolToTokenCache, err
}

// SaveObjectInRedis is used to save a map in Redis
func SaveObjectInRedis(key string, data interface{}, expiryTime time.Duration) error {
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

// SaveDerivativeResponseInDB is used to save derivative obj into db with token as key.
func SaveDerivativeResponseInDB(token int, res *dao.UnderlyingAsset) error {
	if res == nil || res.Payload == nil {
		return fmt.Errorf("response is nil nothing to same db")
	}
	dbUnderlyingAsset := dao.DBUnderlyingAssetObject{DerivativeToDataMap: map[int]dao.Data{}}
	dbUnderlyingAsset.Success = res.Success
	for i := 0; i < len(res.Payload); i++ {
		dbUnderlyingAsset.DerivativeToDataMap[res.Payload[i].Token] = res.Payload[i]
	}
	return SaveObjectInRedis(fmt.Sprintf(consts.DerivativeKey, token), dbUnderlyingAsset, time.Minute)
}

// GetDerivativeResponseFromDB is used to get derivative obj from db with token as key.
func GetDerivativeResponseFromDB(token int) (*dao.UnderlyingAsset, error) {
	derivativeRedisKey := fmt.Sprintf(consts.DerivativeKey, token)
	dbUnderlyingAsset := dao.DBUnderlyingAssetObject{}

	err := GetObjectFromRedis(derivativeRedisKey, &dbUnderlyingAsset)
	underlyingAsset := dao.UnderlyingAsset{Payload: make([]dao.Data, len(dbUnderlyingAsset.DerivativeToDataMap))}
	if err != nil {
		return nil, fmt.Errorf("getting error while fetching derivative response from db")
	}
	i := 0
	for _, data := range dbUnderlyingAsset.DerivativeToDataMap {
		underlyingAsset.Payload[i] = data
		i++
	}
	return &underlyingAsset, nil
}
