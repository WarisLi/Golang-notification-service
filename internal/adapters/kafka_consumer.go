package adapters

import (
	"github.com/WarisLi/Golang-notification-service/internal/core"
	"gopkg.in/Shopify/sarama.v1"
)

type consumerHandler struct {
	eventHandler core.NotificationService
}

func NewConsumerHandler(eventHandler core.NotificationService) sarama.ConsumerGroupHandler {
	return &consumerHandler{eventHandler: eventHandler}
}

func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		obj.eventHandler.HandleEvent(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}

	return nil
}
