package kafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/gzip"
	_ "github.com/segmentio/kafka-go/lz4"
	_ "github.com/segmentio/kafka-go/snappy"
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
		MaxBytes: 10e6, // 10MB,
		MaxWait:  5 * time.Second,
	})

	go printStats(r)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("error reading message: %s\n", err.Error())
			break
		}
		c <- m.Value
		fmt.Printf("message received: %s\n", string(m.Value))
	}

	r.Close()
	Run(params, c)
}

func printStats(r *kafka.Reader) {
	for {
		fmt.Println(r.Stats())
		time.Sleep(5 * time.Second)
	}
}
