package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var k string = `{
    "toAddress": "shariefasfar015@gmail.com",
    "tenantId": "company1",
    "userId": "asfar.sharief",
    "subject": "dummy",
    "body": "THIS IS DUMMY"
}`

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// message := fmt.Sprintf("Job #%d", i)
	err = nc.Publish("jobs", []byte(k))
	if err != nil {
		log.Println("Error publishing:", err)
	} else {
		fmt.Println("Published:", k)
	}

	// for i := 1; i <= 10; i++ {
	// 	message := fmt.Sprintf("Job #%d", i)
	// 	err := nc.Publish("jobs", []byte(message))
	// 	if err != nil {
	// 		log.Println("Error publishing:", err)
	// 	} else {
	// 		fmt.Println("Published:", message)
	// 	}
	// 	time.Sleep(1 * time.Second) // Simulate delay
	// }

	fmt.Println("All jobs published!")
}
