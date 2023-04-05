// Package pubsub provides a simple publish-subscribe messaging system
// with support for multiple subscribers and optional buffering of messages.
// This package defines a Publisher type that can be used to publish messages
// of any type to all current subscribers, and a Subscriber type that can be used
// to receive messages of the same type from a Publisher.
//
// To use this package, first create a new Publisher using NewPublisher,
// passing in a buffer length and any optional configuration options.
//
// Example usage:
// // Create a new publisher with a buffer of 10 messages and a custom logger
// p := pubsub.NewPublisher[string](10, pubsub.WithLogger[string](myLogger))
//
// Then, to subscribe to messages from the publisher, call the Subscribe method
// on the Publisher instance:
//
// Example usage:
// // Subscribe to messages from the publisher
// sub := p.Subscribe()
//
// You can then receive messages by reading from the Subscriber channel:
//
// Example usage:
// // Receive messages from the subscriber channel
// msg := <-sub
//
// To publish a message to all current subscribers, simply call the Publish method
// on the Publisher instance with the message as an argument:
//
// Example usage:
// // Publish a message to all subscribers
// p.Publish("Hello, world!")
//
// You can also unsubscribe from a Publisher by calling the Unsubscribe method
// with the Subscriber channel that you want to unsubscribe:
//
// Example usage:
// // Unsubscribe from the publisher
// p.Unsubscribe(sub)
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package pubsub

import (
	"log"
	"os"
	"sync"

	"github.com/go-extras/go-kit/logger"
)

// The PublisherOption type is a functional option that can be used to configure
// a new Publisher instance.
type PublisherOption[T any] func(*Publisher[T])

// WithLogger is a PublisherOption that sets a custom logger on the Publisher.
func WithLogger[T any](l logger.PrimitiveLogger) PublisherOption[T] {
	return func(p *Publisher[T]) {
		p.logger = l
	}
}

// The Subscriber type is a channel of a specific message type that can be used
// to receive messages from a Publisher.
type Subscriber[T any] chan T

// The Publisher type represents a single publisher that can broadcast messages
// to all current subscribers.
type Publisher[T any] struct {
	mu           sync.RWMutex
	subscribers  map[Subscriber[T]]struct{}
	bufferLength int
	logger       logger.PrimitiveLogger
}

// NewPublisher creates a new Publisher with a buffer length and any optional
// configuration options.
func NewPublisher[T any](bufferLength int, opts ...PublisherOption[T]) *Publisher[T] {
	p := &Publisher[T]{
		subscribers:  make(map[Subscriber[T]]struct{}),
		bufferLength: bufferLength,
		logger:       log.New(os.Stderr, "pubsub.Publisher:", log.LstdFlags),
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Subscribe returns a new Subscriber channel that can be used to receive messages
// from the Publisher.
func (p *Publisher[T]) Subscribe() Subscriber[T] {
	ch := make(chan T, p.bufferLength)
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers[ch] = struct{}{}
	return ch
}

// Unsubscribe removes a Subscriber channel from the Publisher's list of subscribers.
func (p *Publisher[T]) Unsubscribe(ch Subscriber[T]) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.subscribers, ch)
}

// Publish broadcasts a message to all current subscribers.
// If a subscriber's channel buffer is full, the message will be dropped
// and a warning will be logged.
func (p *Publisher[T]) Publish(msg T) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for ch := range p.subscribers {
		select {
		case ch <- msg:
		default:
			p.logger.Print("dropping message because subscriber is too slow (message buffer is full)\n")
		}
	}
}
