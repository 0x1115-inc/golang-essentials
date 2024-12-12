// Copyright 2024 0x1115 Inc

package messages

// Packet is the interface that defines the methods that a message packet should implement
// Structure of the packet is as follows:
// - Attributes: A map of key-value pairs that can be used to store additional information
// - Data: The actual data that the packet is carrying
type Packet interface {	
	String() string	
	AddAttribute(key string, value interface{})
	GetAttribute(key string) interface{}
	RemoveAttribute(key string)
	SetData(data interface{})
}