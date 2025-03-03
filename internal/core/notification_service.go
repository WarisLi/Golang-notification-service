package core

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	events "github.com/WarisLi/Golang-shared-events"
)

type NotificationService interface {
	HandleEvent(topic string, eventBytes []byte)
}

type nofiticationImpl struct {
	notiSender NotificationSender
}

func NewNotificationService(notiSender NotificationSender) NotificationService {
	return &nofiticationImpl{notiSender: notiSender}
}

func (n nofiticationImpl) HandleEvent(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.LowProductQuantityNotificationEvent{}).Name():
		event := events.LowProductQuantityNotificationEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		message := fmt.Sprintf("Product \"%v\" are reaching low stock quantity (%v items remaining).", event.Name, event.Quantity)
		err = n.notiSender.SendTextNotification(message)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.NewOrderNotificationEvent{}).Name():
		event := events.NewOrderNotificationEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		message := fmt.Sprintf("New order (order id :%v)", event.OrderId)
		err = n.notiSender.SendTextNotification(message)
		if err != nil {
			log.Println(err)
			return
		}

	default:
		log.Println("no event handler")
	}

}
