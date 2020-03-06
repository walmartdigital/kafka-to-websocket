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
func Run(params *Params, c chan []byte) {
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
		c <- m.Value
	}

	r.Close()
}
