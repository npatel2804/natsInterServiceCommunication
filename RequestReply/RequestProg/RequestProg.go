package main

import (
	"fmt"
	"log"
	"time"

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
	msg, err := natsConnection.Request(subject, []byte("hello from Publisher"), 10*time.Second)

	fmt.Printf("Request Function error is : %s\n", err)
	log.Println("Published message on subject " + subject)
	log.Printf("The Received Response is :%s", string(msg.Data))
}
