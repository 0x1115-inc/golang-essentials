package messages

type FactoryFunction func(map[string]interface{}) MessageSystem

var registry = make(map[string]FactoryFunction)

func Register(name string, factory FactoryFunction) {
	registry[name] = factory
}

func GetMessageSystem(name string, args map[string]interface{}) MessageSystem {
	if factory, existed := registry[name]; existed {
		return factory(args)
	}
	return nil
}