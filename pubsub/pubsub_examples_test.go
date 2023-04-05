package pubsub_test

import (
	"fmt"
	"sync"

	"github.com/go-extras/go-kit/pubsub"
)

func ExamplePublisher() {
	publisher := pubsub.NewPublisher[string](5)
	subscriber1 := publisher.Subscribe()
	subscriber2 := publisher.Subscribe()

	var wg sync.WaitGroup
	wg.Add(2)

	var sub1, sub2 []string

	// start two goroutines to read from the subscribers
	go func() {
		for msg := range subscriber1 {
			sub1 = append(sub1, msg)
		}
		wg.Done()
	}()

	go func() {
		for msg := range subscriber2 {
			sub2 = append(sub2, msg)
		}
		wg.Done()
	}()

	// publish some messages
	publisher.Publish("hello")
	publisher.Publish("world")

	// unsubscribe subscriber1 and publish another message
	publisher.Unsubscribe(subscriber1)
	publisher.Publish("goodbye")

	// unsubscribe subscriber2
	publisher.Unsubscribe(subscriber2)

	close(subscriber1)
	close(subscriber2)

	wg.Wait()

	for _, v1 := range sub1 {
		fmt.Printf("Subscriber 1: %v\n", v1)
	}

	for _, v1 := range sub2 {
		fmt.Printf("Subscriber 2: %v\n", v1)
	}

	// Output: Subscriber 1: hello
	//Subscriber 1: world
	//Subscriber 2: hello
	//Subscriber 2: world
	//Subscriber 2: goodbye
}
