// Copyright 2024 0x1115 Inc

package messages

const (
	ParameterSubscriptionHandler = "subscription_handler"
)

// MessageSystem is the interface that wraps the basic methods for a message system
// The message system can be used to publish and receive messages from a channel
type MessageSystem interface {
	Publish(channel string, message Packet) error
	Receive(channel string) error	
	SetParameter(key string, value interface{})
}

