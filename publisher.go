package pubsub

import (
	"sync"

	"github.com/apex/log"
)

type Publisher struct {
	topics map[string]*topic
	mu     sync.Mutex
	log    log.Interface
}

var pub *Publisher = &Publisher{
	topics: map[string]*topic{},
	mu:     sync.Mutex{},
	log:    log.Log,
}

// return the singleton client
func GetClient() *Publisher {
	return pub
}

// set a new logger
func (p *Publisher) SetLogger(log log.Interface) {
	p.log = log
}

// returns or creates a new topic
func Topic(name string) NotifySubscriber {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	pub.log.Debugf("[pubsub] get topic %s", name)
	t, ok := pub.topics[name]
	if !ok {
		pub.log.Debugf("[pubsub] create new '%s' topic", name)
		t = &topic{
			name:        name,
			mu:          sync.Mutex{},
			subscribers: make(map[string]Receiver),
		}
	}
	pub.topics[name] = t
	return t
}
