package main

import (
	"fmt"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "publisher"
	natsURL   = "nats://localhost:4222"
	subject   = "order.created"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	message := "Hello, NATS Streaming!"

	if err := sc.Publish(subject, []byte(message)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message published: %s\n", message)

	// Ждем некоторое время, чтобы программа natssetup могла обработать сообщение
	time.Sleep(time.Second * 2)
}
