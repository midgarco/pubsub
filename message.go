package pubsub

import "time"

type Message struct {
	timestamp time.Time
	Data      []byte
}

func (m Message) GetTimestamp() time.Time {
	return m.timestamp
}
