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
	message := PacketTest{Data: "test-message"}

	g := &GCPPubSub{ProjectId: projectID}

	err := g.Publish(channel, message)
	assert.NoError(t, err)
}
