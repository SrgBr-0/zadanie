package main

import (
	"fmt"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "client-123"
	natsURL   = "nats://localhost:4222"
	subject   = "order.created"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	subscription, err := sc.Subscribe(subject, func(msg *stan.Msg) {
		// Обработка полученных данных, например, запись в БД
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	// Добавь код для обработки полученных данных, например, записи в БД
	// ...

	// Поддерживаем соединение для долгосрочной подписки
	select {
	case <-time.After(time.Hour):
		fmt.Println("Subscription timeout")
	}

	fmt.Println("NATS Setup complete...")
}
