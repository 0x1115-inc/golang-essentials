package cache

import (
	"fmt"
	"strconv"
	"strings"

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

	for _, node := range strings.Split(args["nodes"].(string), "|") {
		nodeParts := strings.Split(node, ":")

		// Convert the port to int
		port, err := strconv.Atoi(nodeParts[1])
		if err != nil {
			return nil
		}

		nodes = append(nodes, MemcachedCacheNode{
			Host: nodeParts[0],
			Port: port,
		})
	}
	
	return &MemcachedCache{
		Connections: nodes,
	}
}
