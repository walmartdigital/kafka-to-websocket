package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/segmentio/kafka-go"
)

// Params to initiate kafka consumer
type Params struct {
	GroupID string
	Topic   string
	Brokers string
}

// Run the kafka consumer
func Run(params *Params) {
	fmt.Println("initializing consumer...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(params.Brokers, ","),
		GroupID:  params.GroupID,
		Topic:    params.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}
