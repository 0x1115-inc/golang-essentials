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
const GCPProvider = "GCPPubSubMessageSystem"

type GCPPubSub struct {
	ProjectId string
	SubscriptionId string
}

func init() {
	Register(GCPProvider, NewGCPPubSub)
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

func (g *GCPPubSub) Receive(channel string) (Packet, error) {
	return nil, nil
}

func (g *GCPPubSub) Subscribe(channel string) error {
	return nil
}

func (g *GCPPubSub) Unsubscribe(channel string) error {
	return nil
}

func NewGCPPubSub(args map[string]interface{}) MessageSystem {
	return &GCPPubSub{
		ProjectId: args["project_id"].(string),
	}
}
