// Copyright 2024 0x1115 Inc

package messagesystem

import (
	"os"
	"testing"

	"github.com/0x1115-inc/golang-essentials/pkg/messagesystem/message"
	"github.com/stretchr/testify/assert"
)

func TestGCPPubSub_Publish(t *testing.T) {
	projectID := os.Getenv("PROJECT_ID")
	t.Log("Project ID: ", projectID)
	channel := "test-topic"

	packet := pubsubPacket{
		Headers: map[string]message.IValue{
			"source":      "header-source",
			"messageType": "header-message-type",
		},
		Attributes: map[string]message.IValue{
			"attr1": "Attribute1",
			"attr2": "Attribute2",
		},
		Data: &pubsubPacketData{
			"test-message-from-golang-essentials",
		},
	}
	g := &GCPPubSub{ProjectId: projectID}

	err := g.Publish(channel, &packet)
	assert.NoError(t, err)
}

func TestGCPPubSub_Receive(t *testing.T) {
	var (
		args map[string]interface{}
	)
	projectID := os.Getenv("PROJECT_ID")
	t.Log("Project ID: ", projectID)
	channel := os.Getenv("CHANNEL")

	args = map[string]interface{}{
		"project_id": projectID,
	}

	g := NewGCPPubSub(args)
	g.SetParameter("subscription_handler", func(p message.IPacket) {
		t.Log(p.String())
	})
	g.SetParameter("max_subscribe_messages", 1)

	err := g.Receive(channel)
	assert.NoError(t, err)
}
