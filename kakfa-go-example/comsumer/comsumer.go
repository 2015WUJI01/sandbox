package comsumer

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"log"
)

type Reader struct {
	Reader *kafka.Reader
}

func NewReader() *Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   "groupz",
		Topic:     "test",
		Partition: 0,
		// StartOffset: kafka.LastOffset,
	})
	log.Printf("connect ok")
	return &Reader{Reader: reader}
}

func (r Reader) FetchMessage(ctx context.Context, messages chan<- kafka.Message) error {
	for {
		message, err := r.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case messages <- message:
			log.Printf("message read and sent to a channel: %v\n", string(message.Value))
		}
	}
}

func (r Reader) CommitMessage(ctx context.Context, messageCommitChan <-chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
		case msg := <-messageCommitChan:
			err := r.Reader.CommitMessages(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "Reader.CommitMessage")
			}
			log.Printf("committed an message: %v\n", string(msg.Value))
		}
	}
}
