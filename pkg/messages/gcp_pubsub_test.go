package messages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)
type PacketTest struct {
	Data string
}

func (p PacketTest) String() string {
	return p.Data
}

func (p PacketTest) AddAttribute(key string, value interface{}) {
	// Do nothing
}

func (p PacketTest) GetAttribute(key string) interface{} {
	return nil
}

func (p PacketTest) RemoveAttribute(key string) {
	// Do nothing
}

func TestGCPPubSub_Publish(t *testing.T) {	
	projectID := os.Getenv("PROJECT_ID")
	t.Log("Project ID: ", projectID)
	channel := "test-topic"
	message := PacketTest{Data: "test-message-from-golang-essentials"}

	g := &GCPPubSub{ProjectId: projectID}

	err := g.Publish(channel, message)
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
	g.SetParameter("subscription_handler", func(p Packet) {
		t.Log(p.String())
	})
	g.SetParameter("max_subscribe_messages", 1)

	err := g.Receive(channel)
	assert.NoError(t, err)		
}
