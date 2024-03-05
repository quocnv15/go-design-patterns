package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SMSNotifier struct{}

func (sms SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type TelegramNotifier struct{}

func (notifier TelegramNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Telegram)\n", message)
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

// like add to queue
func (n *NotifierDecorator) Send(message string) {
	if n.core != nil {
		n.core.Send(message)
	}
	n.notifier.Send(message)
}

func (nd NotifierDecorator) Decorate(n Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: n,
	}
}

func NewNotifierDecorator(n Notifier) NotifierDecorator {
	return NotifierDecorator{
		notifier: n,
	}
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(EmailNotifier{}).
		Decorate(SMSNotifier{}).
		Decorate(TelegramNotifier{}).
		Decorate(TelegramNotifier{})
	s := Service{
		notifier: &notifier,
	}
	s.SendNotification("Hello world")
}
