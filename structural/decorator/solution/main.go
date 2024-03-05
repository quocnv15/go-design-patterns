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
// Like add to stack

func (n *NotifierDecorator) Send(message string) {
	n.notifier.Send(message)
	if n.core != nil {
		n.core.Send(message)
	}
}


func (nd NotifierDecorator) StackDecorate(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

// like add to queue
// func (nd NotifierDecorator) QueueDecorate(n Notifier) NotifierDecorator {
// 	if nd.core == nil {
// 		return NotifierDecorator{
// 			core:     &nd,
// 			notifier: n,
// 		}
// 	}
// 	return NotifierDecorator{
// 		core:     nd.core,
// 		notifier: n,
// 	}
// }

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
		StackDecorate(TelegramNotifier{}).
		StackDecorate(SMSNotifier{}).
		StackDecorate(TelegramNotifier{})
		//  		Sending message: Hello world (Sender: Telegram)
		//  		Sending message: Hello world (Sender: SMS)
		//  		Sending message: Hello world (Sender: Telegram)
		//  		Sending message: Hello world (Sender: Email)
	s := Service{
		notifier: &notifier,
	}
	s.SendNotification("Hello world")
}
