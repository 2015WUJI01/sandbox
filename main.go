package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

const (
	kfkAddr = "localhost:9092,localhost:9093,localhost:9094"
)

func main() {

	var conn *kafka.Conn
	var err error
	for _, addr := range strings.Split(kfkAddr, ",") {
		if conn, err = kafka.Dial("tcp", addr); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	brokers, _ := conn.Brokers()
	for _, b := range brokers {
		log.Infof("Broker available: ID=%v Addr=%v:%v Rack=%v", b.ID, b.Host, b.Port, b.Rack)
	}

	P1 := kafka.Writer{
		Addr:         kafka.TCP(strings.Split(kfkAddr, ",")...),
		Topic:        "test4",
		Balancer:     &kafka.RoundRobin{},
		MaxAttempts:  10,
		BatchSize:    100, // 100 条消息
		BatchBytes:   1 * 1024 * 1024,
		BatchTimeout: 10 * time.Millisecond,
		ReadTimeout:  10 * time.Second, // 默认
		WriteTimeout: 10 * time.Second, // 默认
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Completion: func(msgs []kafka.Message, err error) {
			switch err.(type) {
			case nil:
				for _, msg := range msgs {
					layout := "kafka.Message(topic/partition/offset/value): | %10v | %2v | %6v | %10v"
					log.Infof(layout, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
				}
			case kafka.WriteErrors:
				log.Errorf("[WriteError] %s", err)
			default:
				log.Errorf("[UnknowError] %s", err)
			}
		},
		Compression: kafka.Lz4,
		// Logger:                 kafka.LoggerFunc(log.Infof),
		ErrorLogger:            kafka.LoggerFunc(log.Errorf),
		Transport:              nil,
		AllowAutoTopicCreation: true,
	}
	ctx := context.Background()

	t := time.NewTicker(500 * time.Millisecond)
	var i = 0
	for range t.C {
		i++
		msgs := []kafka.Message{
			{Value: bytes.NewBufferString(fmt.Sprintf("%d", i)).Bytes()},
			// {Value: bytes.NewBufferString(fmt.Sprintf("%d", time.Now().UnixNano())).Bytes()},
		}
		_ = P1.WriteMessages(ctx, msgs...)
	}
	// switch err := P1.WriteMessages(ctx, msgs...).(type) {
	// case nil:
	// 	for _, msg := range msgs {
	// 		log.Infof("kafka.Message: Offset=%v Value=%v", msg.Offset, string(msg.Value))
	// 	}
	// case kafka.WriteErrors:
	// 	log.Errorf("[WriteError] %s", err)
	// default:
	// 	log.Errorf("[UnknowError] %s", err)
	// }
	_ = conn.Close()
}
