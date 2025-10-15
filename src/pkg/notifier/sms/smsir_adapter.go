package sms

import "fmt"

type SmsirAdapter struct {
	ApiKey string
}

func (smsir SmsirAdapter) Send(smsMessage *Message) error {
	return fmt.Errorf("SMS.ir provider not yet implemented")
}
