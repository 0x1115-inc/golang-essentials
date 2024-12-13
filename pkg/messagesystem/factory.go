// Copyright 2024 0x1115 Inc

package messagesystem

import (
	"github.com/0x1115-inc/golang-essentials/pkg/messagesystem/message"
)

const (
	ParameterSubscriptionHandler = "subscription_handler"
)

// MessageSystem is the interface that wraps the basic methods for a message system
// The message system can be used to publish and receive messages from a channel
type IMessageSystem interface {
	Publish(channel string, message message.IPacket) error
	Receive(channel string) error

	SetParameter(key string, value interface{})
}
