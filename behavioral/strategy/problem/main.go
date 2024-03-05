package main

import "fmt"

type NotficationService struct {
	notifierType string
}

func (s NotficationService) SendNotification(message string) {
	if s.notifierType == "email" {
		// Send Email
		fmt.Println("Email: ", message)
	} else {
		// Send SMS
		fmt.Println("SMS: ", message)

	}

}

func main() {

	sms := NotficationService{notifierType: "sms"}
	sms.SendNotification("Hello")

	email := NotficationService{notifierType: "email"}
	email.SendNotification("Hello")
}
