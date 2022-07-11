package pubsub

import (
	"context"
	"sync"
	"time"

	"github.com/rs/xid"
)

type topic struct {
	name        string
	subscribers map[string]Receiver
	mu          sync.Mutex
}

func (t *topic) Notify(ctx context.Context, msg *Message) {
	t.mu.Lock()
	defer t.mu.Unlock()

	msg.timestamp = time.Now()

	pub.log.Debugf("[pubsub] notify all subscribers for %s", t.name)
	// pass message to all subscrived Receivers
	for id, sub := range t.subscribers {
		pub.log.Debugf("[pubsub] emit to subscription %s", id)
		go sub.Receive(ctx, msg)
	}
}

func (t *topic) Subscribe(sub Receiver) string {
	t.mu.Lock()
	defer t.mu.Unlock()

	id := xid.New().String()

	pub.log.Debugf("[pubsub] new subscription: %s", id)
	t.subscribers[id] = sub

	return id
}

func (t *topic) Unsubscribe(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	pub.log.Debugf("[pubsub] remove subscription: %s", id)
	delete(t.subscribers, id)
}
