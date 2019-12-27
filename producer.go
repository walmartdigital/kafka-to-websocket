package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/segmentio/kafka-go"
)

type myLogger struct {
}

// Printf ...
func (l *myLogger) Printf(a string, b ...interface{}) {
	glog.Errorf(a, b)
}

func mainProducer(kafkaURL string) {
	fmt.Println("initializing producer...")

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Logger:   &myLogger{},
	})

	defer w.Close()

	for {
		go produce(w)
		time.Sleep(1 * time.Second)
	}
}

func produce(w *kafka.Writer) {
	key := "key-a"
	message := "message " + time.Now().String()
	ctx := context.Background()
	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)

	<-ctx.Done()
	fmt.Println("message published!!! " + ctx.Err().Error())
}
