package pubsub

import "context"

type Receiver interface {
	Receive(ctx context.Context, msg *Message)
}

type Notifier interface {
	Notify(ctx context.Context, msg *Message)
}

type Subscriber interface {
	Subscribe(sub Receiver) string
}

type Unsubscriber interface {
	Unsubscribe(id string)
}

type NotifySubscriber interface {
	Notifier
	Subscriber
	Unsubscriber
}
