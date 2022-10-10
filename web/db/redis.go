package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zrwaite/github-graphs/config"
)

var Cache *redis.Client

func ConnectToRedis() {
	addr := "codegraphs_redis:6379"
	if config.CONFIG.Dev {
		addr = "localhost:6379"
	}
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.CONFIG.RedisPassword, // no password set
		DB:       0,                           // use default DB
	})
	err := SetCache("test", "test")
	if err != nil {
		log.Fatal(err)
	}
}

func ClearCache() {
	ctx := context.Background()
	Cache.FlushDB(ctx)
}

func GetCache(key string) (string, bool) {
	ctx := context.Background()
	val, err := Cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		fmt.Println(err)
		return "", false
	} else {
		return val, true
	}
}

func SetCache(key string, value string) error {
	ctx := context.Background()
	duration, err := time.ParseDuration("1h")
	if err != nil {
		return err
	}
	err = Cache.Set(ctx, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetCacheNoExpire(key string, value string) error {
	ctx := context.Background()
	err := Cache.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetJsonCache(key string, target any) bool {
	cacheJson, found := GetCache(key)
	if found {
		err := json.Unmarshal([]byte(cacheJson), target)
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false
}

func SetJsonCache(key string, value any) error {
	newJson, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = SetCache(key, string(newJson))
	if err != nil {
		return err
	}
	return nil
}

func SetJsonCacheNoExpire(key string, value any) error {
	newJson, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = SetCacheNoExpire(key, string(newJson))
	if err != nil {
		return err
	}
	return nil
}
