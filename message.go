package pubsub

import "time"

type Message struct {
	timestamp time.Time
	topic     string
	Data      []byte
}

func (m Message) GetTimestamp() time.Time {
	return m.timestamp
}

func (m Message) GetTopic() string {
	return m.topic
}
