// Copyright 2024 0x1115 Inc

// Guide to setup the Google Cloud Pub/Sub message system
// Pre-requisites:
// 1. Create a Google Cloud Platform (GCP) account
// 2. Create a project in GCP
// 3. Enable the Pub/Sub API for the project

package messages

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// GCPProvider is the provider name for the Google Cloud Pub/Sub message system when register with the registry.
const (
	GCPPubSubProvider = "GCPPubSub"
	GCPPubSubParameterMaxSubscribeMessages = "max_subscribe_messages"
)

func init() {
	Register(GCPPubSubProvider, NewGCPPubSub)
}


type pubsubPacket struct {
	message *pubsub.Message
}

func (p *pubsubPacket) String() string {
	return string(p.message.Data)
}

func (p *pubsubPacket) AddAttribute(key string, value interface{}) {
	// Do nothing
}

func (p *pubsubPacket) GetAttribute(key string) interface{} {
	return nil
}

func (p *pubsubPacket) RemoveAttribute(key string) {
	// Do nothing
}

func (p *pubsubPacket) SetData(data interface{}) {
	// Do nothing
}

type GCPPubSub struct {
	ProjectId            string
	subscriptionHandler  func(Packet)
	maxSubscribeMessages int
}


func (g *GCPPubSub) Publish(channel string, message Packet) error {
	ctx := context.Background()

	// Create client
	client, err := pubsub.NewClient(ctx, g.ProjectId)
	if err != nil {
		return err
	}
	defer client.Close()

	topic := client.Topic(channel)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message.String()),
	})

	// Block until the result is returned
	_, err = result.Get(ctx)
	return err
}

func (g *GCPPubSub) Receive(channel string) error {
	var (
		ctx context.Context
		cancel context.CancelFunc
		sub *pubsub.Subscription
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
		// Call the subscription handler
		g.subscriptionHandler(&pubsubPacket{message: msg})
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
		g.subscriptionHandler = value.(func(Packet))	
	case GCPPubSubParameterMaxSubscribeMessages:
		g.maxSubscribeMessages = value.(int)
	}
}


func NewGCPPubSub(args map[string]interface{}) MessageSystem {
	return &GCPPubSub{
		ProjectId:            args["project_id"].(string),
		subscriptionHandler:  nil,
		maxSubscribeMessages: 1,
	}
}
