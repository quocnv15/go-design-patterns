package main

import "fmt"

type Notifier interface {
	Send(message string)
}

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

func main() {
	// I don't want my users init a new notifier like this.
	// They should call to something to produce a notifier with its specific type
	// CreateNotifier(type) Notifier
	s := Service{notifier: SMSNotifier{}}
	s.SendNotification("Hello")

	e := Service{notifier: EmailNotifier{}}
	e.SendNotification("Hello")
}
