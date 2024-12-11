package cache

import (
	"fmt"
	"context"

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
	
	return rdb.Del(ctx, key).Err()
}

func NewRedisCache(args map[string]interface{}) Cache {
	return &RedisCache{
		Hostname: args["hostname"].(string),
		Port:     args["port"].(int),
		Password: args["password"].(string),
		Database: args["database"].(int),
	}
}
