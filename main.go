package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {

	// to produce messages
	topic := "test"

	writer := &kafka.Writer{
		Addr:  kafka.TCP("127.0.0.1:9092"),
		Topic: topic,
		// 如果无 topic，则自动生成
		AllowAutoTopicCreation: true,
	}

	for i := 0; i < 10; i++ {
		_ = writer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte("msg"),
		})
		fmt.Printf("已生产 1 条消息:%d\n", i)
		_ = writer.WriteMessages(context.Background(),
			kafka.Message{
				Value: []byte("msg"),
			},
			kafka.Message{
				Value: []byte("msg"),
			},
			kafka.Message{
				Value: []byte("msg"),
			},
		)
		fmt.Printf("已生产 3 条消息:%d\n", i)
	}

	// writer := &kafka.Writer{
	// 	Addr:  kafka.TCP("127.0.0.1:9092"),
	// 	Topic: "topic_1",
	// }
	// ctx := context.Background()
	// for i := 0; i < 100; i++ {
	// 	err := writer.WriteMessages(ctx, kafka.Message{
	// 		Value: []byte(fmt.Sprintf("this is msg %d", i)),
	// 	})
	//
	// 	if err != nil {
	// 		log.Println("Error", err)
	// 		return
	// 	} else {
	// 		log.Println(i, "write success")
	// 	}
	// }
}
