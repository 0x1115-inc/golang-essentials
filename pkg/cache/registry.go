package cache

type FactoryFunction func(map[string]interface{}) Cache

var registry = make(map[string]FactoryFunction)

func Register(name string, factory FactoryFunction) {
	registry[name] = factory
}

func GetCacheInstance(name string, args map[string]interface{}) Cache {
	if factory, existed := registry[name]; existed {
		return factory(args)
	}
	return nil
}