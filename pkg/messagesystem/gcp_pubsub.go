// Copyright 2024 0x1115 Inc

// Guide to setup the Google Cloud Pub/Sub message system
// Pre-requisites:
// 1. Create a Google Cloud Platform (GCP) account
// 2. Create a project in GCP
// 3. Enable the Pub/Sub API for the project

package messagesystem

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/0x1115-inc/golang-essentials/pkg/messagesystem/message"
)

// GCPProvider is the provider name for the Google Cloud Pub/Sub message system when register with the registry.
const (
	GCPPubSubProvider                      = "GCPPubSub"
	GCPPubSubParameterMaxSubscribeMessages = "max_subscribe_messages"
	GCPPubSubMessageHeaderPrefix           = "__"
)

func init() {
	Register(GCPPubSubProvider, NewGCPPubSub)
}

type pubsubPacketData struct {
	Message string `json:"m"`
}

func (p *pubsubPacketData) String() string {
	return p.Message
}

type pubsubPacket struct {
	Headers    map[string]message.IValue `json:"h"`
	Attributes map[string]message.IValue `json:"a"`
	Data       *pubsubPacketData         `json:"d"`
}

func (p *pubsubPacket) String() string {
	// Json marshal the struct
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func (p *pubsubPacket) SetHeader(key string, value message.IValue) {
	p.Headers[key] = value
}

func (p *pubsubPacket) GetHeader(key string) message.IValue {
	return p.Headers[key]
}

func (p *pubsubPacket) GetHeaders() map[string]message.IValue {
	return p.Headers
}

func (p *pubsubPacket) SetAttribute(key string, value message.IValue) {
	p.Attributes[key] = value
}

func (p *pubsubPacket) GetAttribute(key string) message.IValue {
	return p.Attributes[key]
}

func (p *pubsubPacket) GetAttributes() map[string]message.IValue {
	return p.Attributes
}

func (p *pubsubPacket) GetData() message.IData {
	return p.Data
}

func (p *pubsubPacket) SetData(data message.IData) {
	p.Data = &pubsubPacketData{
		data.String(),
	}
}

type GCPPubSub struct {
	ProjectId            string
	subscriptionHandler  func(message.IPacket)
	maxSubscribeMessages int
}

func (g *GCPPubSub) Publish(channel string, packet message.IPacket) error {
	var (
		ctx context.Context
		err error

		client        *pubsub.Client
		topic         *pubsub.Topic
		pubsubMessage *pubsub.Message

		isTopicExisted bool
		attributes     map[string]string
		packetAttr     map[string]message.IValue
	)
	ctx = context.Background()

	// Create client
	client, err = pubsub.NewClient(ctx, g.ProjectId)
	if err != nil {
		return err
	}
	defer client.Close()

	topic = client.Topic(channel)
	isTopicExisted, err = topic.Exists(ctx)
	if err != nil {
		return err
	}

	if !isTopicExisted {
		// Topic creation will be implement in the future
		return fmt.Errorf("topic %s does not exist", channel)
	}

	// Setup message
	// Convert the packet attributes to a map of string
	packetAttr = packet.GetAttributes()
	if packetAttr != nil {
		attributes = make(map[string]string)
		for key, value := range packetAttr {		
			// Convert the value to string
			attributes[key] = fmt.Sprintf("%v", value)
		
		}
	}

	// Convert the packet header as the special attribute with prefix defined in constant
	packetHeaders := packet.GetHeaders()
	if packetHeaders != nil {
		if attributes == nil {
			attributes = make(map[string]string)
		}
		for key, value := range packetHeaders {			
			// Convert the value to string
			attributes[fmt.Sprintf("%s%s", GCPPubSubMessageHeaderPrefix, key)] = fmt.Sprintf("%v", value)			
		}
	}

	pubsubMessage = &pubsub.Message{
		Data:       []byte(packet.GetData().String()),
		Attributes: attributes,
	}

	result := topic.Publish(ctx, pubsubMessage)

	// Block until the result is returned
	_, err = result.Get(ctx)
	return err
}

func (g *GCPPubSub) Receive(channel string) error {
	var (
		ctx              context.Context
		cancel           context.CancelFunc
		sub              *pubsub.Subscription
		receivedMessages int = 0
	)
	ctx = context.Background()
	client, err := pubsub.NewClient(ctx, g.ProjectId)
	if err != nil {
		return err
	}
	defer client.Close()

	sub = client.Subscription(channel)
	ctx, cancel = context.WithCancel(ctx)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// Convert the pubsub message to a packet
		packetData := &pubsubPacketData{
			string(msg.Data),
		}

		packet := &pubsubPacket{
			Headers:    make(map[string]message.IValue),
			Attributes: make(map[string]message.IValue),
			Data:       packetData,
		}
		for key, value := range msg.Attributes {
			// Check if the attribute contains special prefix
			if len(key) > len(GCPPubSubMessageHeaderPrefix) && key[:len(GCPPubSubMessageHeaderPrefix)] == GCPPubSubMessageHeaderPrefix {
				// Add the header to the packet
				packet.SetHeader(key[len(GCPPubSubMessageHeaderPrefix):], value)
				continue
			}

			// Add the attribute to the packet
			packet.SetAttribute(key, value)
		}

		// Call the subscription handler
		g.subscriptionHandler(packet)
		msg.Ack()

		// Cancel the subscription after receiving a number of messages
		receivedMessages++
		if receivedMessages >= g.maxSubscribeMessages {
			cancel()
		}
	})
	return err
}

func (g *GCPPubSub) SetParameter(key string, value interface{}) {
	switch key {
	case ParameterSubscriptionHandler:
		g.subscriptionHandler = value.(func(message.IPacket))
	case GCPPubSubParameterMaxSubscribeMessages:
		g.maxSubscribeMessages = value.(int)
	}
}

func NewGCPPubSub(args map[string]interface{}) IMessageSystem {
	return &GCPPubSub{
		ProjectId:            args["project_id"].(string),
		subscriptionHandler:  nil,
		maxSubscribeMessages: 1,
	}
}
