// Copyright 2024 0x1115 Inc

package messagesystem

type FactoryFunction func(map[string]interface{}) IMessageSystem

var registry = make(map[string]FactoryFunction)

func Register(name string, factory FactoryFunction) {
	registry[name] = factory
}

func GetMessageSystem(name string, args map[string]interface{}) IMessageSystem {
	if factory, existed := registry[name]; existed {
		return factory(args)
	}
	return nil
}