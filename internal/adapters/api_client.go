package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/WarisLi/Golang-notification-service/internal/core"
)

type APIClient struct {
	client *http.Client
}

func NewAPIClient() core.NotificationSender {
	return &APIClient{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *APIClient) SendTextNotification(message string) error {
	url := "https://api.line.me/v2/bot/message/push"
	data := map[string]any{
		"to": os.Getenv("LINE_NOTIFICATION_USER_ID"),
		"messages": []map[string]string{{
			"type": "text",
			"text": message,
		}},
	}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")))

	_, err = a.client.Do(req)
	if err != nil {
		return err
	}

	log.Println("Notification sent successfully.")

	return nil
}
