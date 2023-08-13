package main

import (
	"fmt"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

const subject = "Order.>"

func main() {

	a := func(msg *nats.Msg) {
		fmt.Printf("in the handler")
		fmt.Printf("Received message from Publisher:%s", string(msg.Data))
		msg.Respond([]byte("Hey mali gayo taro message"))
	}
	// Create server connection
	nc, err := nats.Connect(nats.DefaultURL)
	fmt.Printf("%s", err)
	//defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	nc.QueueSubscribe(subject, "tiptop", a)
	nc.Flush()

	//natsConnection.Subscribe(subject, a)

	// Keep the connection alive
	runtime.Goexit()
}
