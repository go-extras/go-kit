package pubsub_test

import (
	"bytes"
	"log"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/go-extras/go-kit/pubsub"
)

func TestPublisher(t *testing.T) {
	c := qt.New(t)

	// instantiate logrus debug logger
	// Create a buffer to capture the logs
	var buf bytes.Buffer

	// Create a new logger that writes to the buffer
	debugLog := log.New(&buf, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)

	p := pubsub.NewPublisher[string](5, pubsub.WithLogger[string](debugLog))

	// Subscribe to the publisher
	sub1 := p.Subscribe()

	// Publish some messages
	p.Publish("message 1")
	p.Publish("message 2")

	// Receive messages from the subscription
	c.Assert(<-sub1, qt.Equals, "message 1")
	c.Assert(<-sub1, qt.Equals, "message 2")

	// Unsubscribe from the publisher
	p.Unsubscribe(sub1)

	// Publish a message after unsubscribing
	p.Publish("message 3")

	// Ensure that the unsubscribed subscription did not receive the message
	select {
	case msg, ok := <-sub1:
		c.Fatal("channel must not be closed and must have received a message", qt.Commentf("msg: %v, ok: %v", msg, ok))
	default:
		// ok
	}

	// Subscribe to the publisher again
	sub2 := p.Subscribe()

	// Publish some messages
	p.Publish("message 4")
	p.Publish("message 5")

	// Receive messages from the new subscription
	c.Assert(<-sub2, qt.Equals, "message 4")
	c.Assert(<-sub2, qt.Equals, "message 5")

	// Unsubscribe from the publisher again
	p.Unsubscribe(sub2)

	// Publish a message after unsubscribing again
	p.Publish("message 6")

	// Ensure that the unsubscribed subscription did not receive the message
	select {
	case msg, ok := <-sub2:
		c.Fatal("channel must not be closed and must have received a message", qt.Commentf("msg: %v, ok: %v", msg, ok))
	default:
		// ok
	}

	expected := "dropping message because subscriber is too slow (message buffer is full)\n"
	c.Assert(buf.String(), qt.Not(qt.Contains), expected)
}

func TestPublisher_NoReceive(t *testing.T) {
	c := qt.New(t)

	// instantiate logrus debug logger
	// Create a buffer to capture the logs
	var buf bytes.Buffer

	// Create a new logger that writes to the buffer
	debugLog := log.New(&buf, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)

	p := pubsub.NewPublisher[string](2, pubsub.WithLogger[string](debugLog))

	// Subscribe to the publisher
	// Never receive messages from the subscription
	sub1 := p.Subscribe()

	// Publish some messages
	p.Publish("message 1")
	p.Publish("message 2")

	// Publish a message above the buffer size
	p.Publish("message 3")

	expected := "dropping message because subscriber is too slow (message buffer is full)\n"
	c.Assert(buf.String(), qt.Contains, expected)

	// Unsubscribe from the publisher
	p.Unsubscribe(sub1)
}
