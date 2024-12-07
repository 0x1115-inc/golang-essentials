package messages

type MessageSystem interface {
	Publish(channel string, message Packet) error
	Receive(channel string) (Packet, error)
	Subscribe(channel string) error
	Unsubscribe(channel string) error
}

