package main

import (
	"log"
	"runtime"
	"fmt"

	nats"github.com/nats-io/nats.go"
)

const subject = "Order.>"

func main() {

	a := func(msg *nats.Msg){
		fmt.Printf("in the handler")
		fmt.Printf("Received message from Publisher:%s",string(msg.Data))
	}
	// Create server connection
	natsConnection, err := nats.Connect(nats.DefaultURL)
	fmt.Printf("%s",err)
	//defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(subject,a) 

	// Keep the connection alive
	runtime.Goexit()
}
