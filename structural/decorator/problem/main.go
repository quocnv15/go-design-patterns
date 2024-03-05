package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Println("Sending message: %s (Sender: Email)", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Println("Sending message: %s (Sender: SMS)", message)
}

type EmailSMSNotifier struct {
	emailNotifier EmailNotifier
	smsNotifier   SMSNotifier
}

func (n EmailSMSNotifier) Send(message string) {
	n.emailNotifier.Send(message)
	n.smsNotifier.Send(message)
}

type SMSEmailNotifier struct {
	emailNotifier EmailNotifier
	smsNotifier   SMSNotifier
}

func (n SMSEmailNotifier) Send(message string) {
	n.smsNotifier.Send(message)
	n.emailNotifier.Send(message)

}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	fmt.Println("####---EmailSMSNotifier---####")

	s := Service{
		notifier: EmailSMSNotifier{},
	}

	s.SendNotification("Hello world")
	fmt.Println("####---SMSEmailNotifie---####")
	s1 := Service{
		notifier: SMSEmailNotifier{},
	}

	s1.SendNotification("Hello world")
}
