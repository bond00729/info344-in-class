package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func listen(messages <-chan amqp.Delivery) {
	log.Println("listening for new messages...")
	for message := range messages {
		log.Println(string(message.Body))
	}
}

func main() {
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		mqAddr = "localhost:5672"
	}
	mqURL := fmt.Sprintf("amqp://%s", mqAddr)
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatalf("error connecting to rabbit mq: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error establishing connection: %v", err)
	}
	q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error declaring the queue: %v", err)
	}

	messages, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("error consuming channel: %v", err)
	}

	go listen(messages)

	neverEnd := make(chan bool)
	<-neverEnd
}
