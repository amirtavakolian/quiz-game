package sms

type SMSProvider interface {
	Send(smsMessage *Message) error
}
