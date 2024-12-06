package messages

type MessageSystem interface {
	Publish(channel string, message string) error
	Subscribe(channel string) error
	Unsubscribe(channel string) error
}

