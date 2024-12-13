package message

// IPacket is the interface that defines the methods that a message packet should implement
// Structure of the packet is as follows:
// - Header: A map of key-value pairs that can be used to store metadata about the packet
// - Attributes: A map of key-value pairs that can be used to store context and unique information
// - Data: The actual data that the packet is carrying
type IPacket interface {
	SetHeader(key string, value IValue)
	GetHeader(key string) IValue
	GetHeaders() map[string]IValue
	SetAttribute(key string, value IValue)
	GetAttribute(key string) IValue
	GetAttributes() map[string]IValue
	SetData(data IData)
	GetData() IData

	String() string
}

// IData is the interface that defines the methods that a data object should implement
type IData interface {	
	String() string
}

type IValue interface{}
