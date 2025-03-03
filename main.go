package main

import (
	"context"
	"fmt"
	"os"

	events "github.com/WarisLi/Golang-shared-events"

	"github.com/WarisLi/Golang-notification-service/internal/adapters"
	"github.com/WarisLi/Golang-notification-service/internal/core"
	"github.com/joho/godotenv"
	"gopkg.in/Shopify/sarama.v1"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{os.Getenv("KAFKA_SERVERS")},
		os.Getenv("KAFKA_GROUP"),
		config,
	)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	notiSender := adapters.NewAPIClient()
	notiService := core.NewNotificationService(notiSender)
	notiHandler := adapters.NewConsumerHandler(notiService)

	fmt.Println("Notification consumer start.")
	for {
		consumerGroup.Consume(context.Background(), events.Topics, notiHandler)
	}
}
