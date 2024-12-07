package messages

import (
	"testing"	
	"github.com/stretchr/testify/assert"
)

func TestGCPPubSub_Publish(t *testing.T) {
	projectID := "test-project"
	channel := "test-topic"
	message := Packet{Data: "test message"}

	g := &GCPPubSub{ProjectId: projectID}

	err := g.Publish(channel, message)
	assert.NoError(t, err)
}

func TestGCPPubSub_Receive(t *testing.T) {
	projectID := "test-project"
	channel := "test-subscription"

	g := &GCPPubSub{ProjectId: projectID}

	packet, err := g.Receive(channel)
	assert.NoError(t, err)
	assert.Nil(t, packet)
}

func TestGCPPubSub_Subscribe(t *testing.T) {
	projectID := "test-project"
	channel := "test-subscription"

	g := &GCPPubSub{ProjectId: projectID}

	err := g.Subscribe(channel)
	assert.NoError(t, err)
}

func TestGCPPubSub_Unsubscribe(t *testing.T) {
	projectID := "test-project"
	channel := "test-subscription"

	g := &GCPPubSub{ProjectId: projectID}

	err := g.Unsubscribe(channel)
	assert.NoError(t, err)
}

func TestNewGCPPubSub(t *testing.T) {
	args := map[string]interface{}{
		"project_id": "test-project",
	}

	g := NewGCPPubSub(args)
	assert.Equal(t, "test-project", g.ProjectId)
}