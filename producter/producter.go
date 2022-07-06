package producter

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Writer struct {
	Writer *kafka.Writer
}

func (w Writer) WriteMessage(ctx context.Context, msgs chan kafka.Message, msgCommitChan chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case m := <-msgs:
			err := w.Writer.WriteMessages(ctx, kafka.Message{
				// Topic:     "test",
				// Partition: 0,
				Value: m.Value,
			})
			if err != nil {
				return err
			}
			select {
			case <-ctx.Done():
			case msgCommitChan <- m:
			}
		}
	}
}

func NewWriter() *Writer {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("127.0.0.1:9092"),
		Topic: "test",
		Async: true,
	}
	return &Writer{Writer: writer}
}
