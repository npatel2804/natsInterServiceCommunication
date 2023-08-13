package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

const (
	aggregate = "Order"
	event     = "OrderCreated"
)

// publishOrderCreated publish an event via NATS server
func main() {
	// Connect to NATS server
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()

	subject := "Order.OrderCreated"
	// Publish message on subject
	err := natsConnection.Publish(subject, []byte("hello from Publisher"))
	fmt.Printf("%s\n", err)
	log.Println("Published message on subject " + subject)
}
