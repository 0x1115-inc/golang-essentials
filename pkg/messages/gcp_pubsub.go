package messages

const GCPProvider = "GCPPubSubMessageSystem"
type GCPPubSub struct {
}

func init() {
	Register(GCPProvider, NewGCPPubSub)
}

func (g *GCPPubSub) Publish(channel string, message string) error {
	return nil
}

func (g *GCPPubSub) Subscribe(channel string) error {
	return nil
}

func (g *GCPPubSub) Unsubscribe(channel string) error {
	return nil
}

func NewGCPPubSub(args ...map[string]interface{}) MessageSystem {
	return &GCPPubSub{}
}