package core

type NotificationSender interface {
	SendTextNotification(string) error
}
