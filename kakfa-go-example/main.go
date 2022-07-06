package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {

	writer := &kafka.Writer{
		Addr:  kafka.TCP("127.0.0.1:9092"),
		Topic: "topic_1",
	}
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		err := writer.WriteMessages(ctx, kafka.Message{
			Value: []byte(fmt.Sprintf("this is msg %d", i)),
		})

		if err != nil {
			log.Println("Error", err)
			return
		} else {
			log.Println(i, "write success")
		}
	}
}
