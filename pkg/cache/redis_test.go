package cache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache_Set(t *testing.T) {
	cache := NewRedisCache(map[string]interface{}{
		"hostname": "localhost",
		"port":     6379,
		"password": "",
		"database": 0,
	})

	err := cache.Set("test-key", "test-value")
	assert.NoError(t, err)
}

func TestRedisCache_Get(t *testing.T) {
	cache := NewRedisCache(map[string]interface{}{
		"hostname": "localhost",
		"port":     6379,
		"password": "",
		"database": 0,
	})

	err := cache.Set("test-key", "test-value")
	assert.NoError(t, err)

	value, err := cache.Get("test-key")
	assert.NoError(t, err)
	assert.Equal(t, "test-value", value)
}

func TestRedisCache_Get_NotFound(t *testing.T) {
	cache := NewRedisCache(map[string]interface{}{
		"hostname": "localhost",
		"port":     6379,
		"password": "",
		"database": 0,
	})

	_, err := cache.Get("non-existent-key")
	assert.Error(t, err)
	assert.IsType(t, &CacheError{}, err)
	assert.Equal(t, CacheErrorNotFound, err.(*CacheError).Code)
}

func TestRedisCache_Delete(t *testing.T) {
	cache := NewRedisCache(map[string]interface{}{
		"hostname": "localhost",
		"port":     6379,
		"password": "",
		"database": 0,
	})

	err := cache.Set("test-key", "test-value")
	assert.NoError(t, err)

	err = cache.Delete("test-key")
	assert.NoError(t, err)

	_, err = cache.Get("test-key")
	assert.Error(t, err)
	assert.IsType(t, &CacheError{}, err)
	assert.Equal(t, CacheErrorNotFound, err.(*CacheError).Code)
}