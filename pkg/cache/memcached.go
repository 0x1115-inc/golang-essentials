package cache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

const MemcachedCacheType = "memcached"

func init() {
	Register(MemcachedCacheType, NewMemcachedCache)
}

type MemcachedCacheNode struct {
	Host string
	Port int
}

type MemcachedCache struct {
	Connections []MemcachedCacheNode
}

func (m MemcachedCache) Strings() []string {
	var nodes []string
	for _, node := range m.Connections {
		nodes = append(nodes, fmt.Sprintf("%s:%d", node.Host, node.Port))
	}
	return nodes
}

func (m *MemcachedCache) Set(key string, value interface{}) error {
	memcacheInstance := memcache.New(m.Strings()...)	
	defer memcacheInstance.Close()

	return memcacheInstance.Set(&memcache.Item{
		Key:   key,
		Value: []byte(fmt.Sprintf("%v", value)),
	})
}

func (m *MemcachedCache) Get(key string) (interface{}, error) {
	memcacheInstance := memcache.New(m.Strings()...)	
	defer memcacheInstance.Close()

	item, err := memcacheInstance.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, &CacheError{
			Code:    CacheErrorNotFound,
			Message: "Key not found",
		}
	}	

	return item, err
}

func (m *MemcachedCache) Delete(key string) error {
	memcacheInstance := memcache.New(m.Strings()...)	
	defer memcacheInstance.Close()

	return memcacheInstance.Delete(key)
}

func NewMemcachedCache(args map[string]interface{}) Cache {
	var nodes []MemcachedCacheNode

	for _, node := range args["nodes"].([]interface{}) {
		n := node.(map[string]interface{})
		nodes = append(nodes, MemcachedCacheNode{
			Host: n["host"].(string),
			Port: n["port"].(int),
		})
	}
	return &MemcachedCache{
		Connections: nodes,
	}
}
