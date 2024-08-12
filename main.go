package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//config
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "my-group",
	}
	// Buat consumer
	c, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatal("failed to create consumer: %s", err)
	}
	defer c.Close()

	// Subscrib ke topic
	err = c.Subscribe("testing", nil)
	if err != nil {
		log.Fatal("failed to subscribe: %s", err)
	}

	for {
		msg, err := c.ReadMessage(100)
		if err == nil {
			fmt.Println("message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// fmt.Println(err)
		}
	}

}
