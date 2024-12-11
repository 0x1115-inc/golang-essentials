package messages

type MessageSystem interface {
	Publish(channel string, message Packet) error
	Receive(channel string) error	
	SetParameter(key string, value interface{})
}
