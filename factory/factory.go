// Abstract Factory Pattern
package main

import "fmt"

type INotifyFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	println("Sending notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notiticationType string) (INotifyFactory, error) {
	if notiticationType == "SMS" {
		return new(SMSNotification), nil
	}
	if notiticationType == "Email" {
		return new(EmailNotification), nil
	}
	return nil, fmt.Errorf("no notification type")
}

func getMethod(factory INotifyFactory) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	smsFactory.SendNotification()
	emailFactory.SendNotification()
	getMethod(smsFactory)
	getMethod(emailFactory)
}
