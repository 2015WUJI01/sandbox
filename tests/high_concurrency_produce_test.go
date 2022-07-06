package tests

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestHighConcurrencyWriting(t *testing.T) {
	// to produce messages
	topic := "test"

	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: topic,
		// 如果无 topic，则自动生成
		AllowAutoTopicCreation: true,
	}

	// for i := 0; i < 10; i++ {
	_ = writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("msg"),
	})
	// t.Logf("已生产一条消息:%d\n", i)
	// }
}
