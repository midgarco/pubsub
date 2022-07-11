# Simple PubSub module

Example:

```
package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/midgarco/pubsub"
)

func main() {
	log.SetHandler(cli.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	client := pubsub.GetClient()
	client.SetLogger(log.Log)

	wg := sync.WaitGroup{}

	foo := &Foo{wg: &wg}
	bar := &Bar{wg: &wg}

	// create topic
	t := pubsub.Topic("new-event")
	t2 := pubsub.Topic("create-event")

	// subscribers
	foosubid := t.Subscribe(foo)
	_ = t.Subscribe(bar)
	_ = t2.Subscribe(bar)

	// send out event message
	wg.Add(3)
	t.Notify(context.Background(), &pubsub.Message{
		Data: []byte("hello world"),
	})
	t2.Notify(context.Background(), &pubsub.Message{
		Data: []byte("hello new world"),
	})

	wg.Wait()

	// remove foo sub
	t.Unsubscribe(foosubid)

	// send out event message
	wg.Add(2)
	t.Notify(context.Background(), &pubsub.Message{
		Data: []byte("hello world, again"),
	})
	t2.Notify(context.Background(), &pubsub.Message{
		Data: []byte("hello new world, again"),
	})

	wg.Wait()

	log.Info("done.")
}

type Foo struct {
	wg *sync.WaitGroup
}

func (f *Foo) Receive(ctx context.Context, msg *pubsub.Message) {
	defer f.wg.Done()

	log.Infof("message received:[%s] %s", msg.GetTimestamp().Format(time.RFC3339), string(msg.Data))
}

type Bar struct {
	wg *sync.WaitGroup
}

func (b *Bar) Receive(ctx context.Context, msg *pubsub.Message) {
	defer b.wg.Done()

	log.Infof("message received:[%s] %s", msg.GetTimestamp().Format(time.RFC3339), string(msg.Data))
}
```
