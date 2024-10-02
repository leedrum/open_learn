package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeeType   string `json:"coffee_type"`
}

func main() {
	http.HandleFunc("/orders", placeOrder)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func pushOrderToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}
	// Create connections
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	// Create kafka message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("Order is stored in topic %s, partition %d and offset %d \n", topic, partition, offset)

	return nil
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request type is invalid", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse request body into order
	order := new(Order)
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Default().Println(err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	// 2. Convert body into bytes
	orderInBytes, err := json.Marshal(order)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// 3. Send the bytes to kafka
	err = pushOrderToQueue("coffee_orders", orderInBytes)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Can not place the orders", http.StatusBadRequest)
	}

	// 4. Respond back to the user
	response := map[string]interface{}{
		"success": true,
		"msg":     "Order for " + order.CustomerName + " placed successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatalln(err)
		http.Error(w, "Something went wrong with encode response", http.StatusInternalServerError)
		return
	}
}
