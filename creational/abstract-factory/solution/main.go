package main

import "fmt"

type Notifier interface {
	Send(message string)
}
type NotifierType int

const (
	SMS NotifierType = iota
	Email
)

type SMSNotifier struct{}
type EmailNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("SMS: %s (Sender: SMS) \n", message)
}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Email: %s (Sender: Email)\n", message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t NotifierType) Notifier {
	switch t {
	case SMS:
		return &SMSNotifier{}
	case Email:
		return &EmailNotifier{}
	default:
		return nil
	}
}

func main() {
	s := Service{notifier: CreateNotifier(SMS)}
	s.SendNotification("Hello")

	e := Service{notifier: CreateNotifier(Email)}
	e.SendNotification("Hello")
}
