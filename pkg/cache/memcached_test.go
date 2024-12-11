package cache

import (
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

// Test multiple connections
func TestMemcachedCache_Connections(t *testing.T) {
	cache := NewMemcachedCache(map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{"host": "localhost", "port": 11211},
			map[string]interface{}{"host": "localhost", "port": 11212},
		},
	})

	if len(cache.(*MemcachedCache).Connections) != 2 {
		t.Fatalf("expected 2 connections, got %d", len(cache.(*MemcachedCache).Connections))
	}
}

func TestMemcachedCache_SetGet(t *testing.T) {
	cache := NewMemcachedCache(map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{"host": "localhost", "port": 11211},
		},
	})

	key := "test-key"
	value := "test-value"

	err := cache.Set(key, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, err := cache.Get(key)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.(*memcache.Item).Value == nil || string(result.(*memcache.Item).Value) != value {
		t.Fatalf("expected %v, got %v", value, result)
	}
}

func TestMemcachedCache_GetNotFound(t *testing.T) {
	cache := NewMemcachedCache(map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{"host": "localhost", "port": 11211},
		},
	})

	_, err := cache.Get("non-existent-key")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if cacheErr, ok := err.(*CacheError); !ok || cacheErr.Code != CacheErrorNotFound {
		t.Fatalf("expected CacheErrorNotFound, got %v", err)
	}
}

func TestMemcachedCache_Delete(t *testing.T) {
	cache := NewMemcachedCache(map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{"host": "localhost", "port": 11211},
		},
	})

	key := "test-key"
	value := "test-value"

	err := cache.Set(key, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = cache.Delete(key)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = cache.Get(key)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if cacheErr, ok := err.(*CacheError); !ok || cacheErr.Code != CacheErrorNotFound {
		t.Fatalf("expected CacheErrorNotFound, got %v", err)
	}
}