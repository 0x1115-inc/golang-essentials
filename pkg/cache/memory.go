package cache

const (
	MemoryCacheType = "memory"
)

func init() {
	Register(MemoryCacheType, NewMemoryCache)
}

type MemoryCache struct {
	cache map[string]interface{}
}

func (m *MemoryCache) Set(key string, value interface{}) error {
	m.cache[key] = value
	return nil
}

func (m *MemoryCache) Get(key string) (interface{}, error) {
	if value, existed := m.cache[key]; existed {
		return value, nil
	}
	return nil, &CacheError{
		Code:    CacheErrorNotFound,
		Message: "Key not found",
	}
}

func (m *MemoryCache) Delete(key string) error {
	delete(m.cache, key)
	return nil
}

func NewMemoryCache(args map[string]interface{}) Cache {
	return &MemoryCache{
		cache: make(map[string]interface{}),
	}
}