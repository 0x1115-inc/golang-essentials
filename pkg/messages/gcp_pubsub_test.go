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
		"subscription_handler": func(p Packet) {
			t.Log(p.String())
		},
		"subscription_max_messages": 1,

	}
	g := NewGCPPubSub(args)

	err := g.Receive(channel)
	assert.NoError(t, err)
}