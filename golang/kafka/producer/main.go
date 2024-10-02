package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "coffee_orders"
	msgCount := 0
	// 1. Create a new consumer and start it
	consumer, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer started!")

	// 2. Handle OS sinals - Used to stop the process
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Create a Go-routin to run consumer / producer
	doneChan := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				panic(err)
			case msg := <-partitionConsumer.Messages():
				msgCount++
				fmt.Printf("Recevied Order count: %d | Topic: %s | message: %s\n", msgCount, topic, msg.Value)
				order := string(msg.Value)
				fmt.Printf("Brewing coffee for order: %s\n", order)
			case <-sigChan:
				fmt.Println("Interrupt is detected")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	fmt.Printf("Processed: %d\n", msgCount)

	// 4. Close the consumer on exit
	err = partitionConsumer.Close()
	if err != nil {
		panic(err)
	}
}

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
