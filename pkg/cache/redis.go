package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	RedisCacheType = "redis"
)

func init() {
	Register(RedisCacheType, NewRedisCache)
}

type RedisCache struct {
	Hostname string
	Port     int
	Password string
	Database int
}

func (r *RedisCache) Set(key string, value interface{}) error {
	var (
		ctx = context.Background()
	)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Hostname, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	defer rdb.Close()

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	var (
		ctx = context.Background()
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Hostname, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	defer rdb.Close()

	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, &CacheError{
			Code:    CacheErrorNotFound,
			Message: "Key not found",
		}
	}

	return value, err
}

func (r *RedisCache) Delete(key string) error {
	var (
		ctx = context.Background()
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Hostname, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	defer rdb.Close()

	return rdb.Del(ctx, key).Err()
}

func NewRedisCache(args map[string]interface{}) Cache {
	// Parameter conversion
	port, err := strconv.ParseInt(fmt.Sprintf("%s", args["port"]), 10, 32)
	if err != nil {
		return nil
	}

	database, err := strconv.ParseInt(fmt.Sprintf("%s", args["db"]), 10, 32)
	if err != nil {
		return nil
	}

	return &RedisCache{
		Hostname: args["host"].(string),
		Port:     int(port),
		Password: args["password"].(string),
		Database: int(database),
	}
}
