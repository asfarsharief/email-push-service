package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to the "jobs" subject
	_, err = nc.Subscribe("jobs", func(msg *nats.Msg) {
		fmt.Println("Received job:", string(msg.Data))
	})

	if err != nil {
		log.Fatal("Error subscribing:", err)
	}

	fmt.Println("Listening for jobs... Press Ctrl+C to exit.")
	select {} // Keep running
}
