package cache

import (
	"testing"
)

func TestMemoryCache_SetGet(t *testing.T) {
	cache := NewMemoryCache(nil)

	err := cache.Set("key1", "value1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	value, err := cache.Get("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if value != "value1" {
		t.Fatalf("expected value1, got %v", value)
	}
}

func TestMemoryCache_GetNonExistentKey(t *testing.T) {
	cache := NewMemoryCache(nil)

	_, err := cache.Get("nonexistent")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache(nil)

	err := cache.Set("key1", "value1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = cache.Delete("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = cache.Get("key1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}