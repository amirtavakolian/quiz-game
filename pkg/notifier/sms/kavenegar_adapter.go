package sms

import (
	"github.com/kavenegar/kavenegar-go"
)

type KavenegarAdapter struct {
	ApiKey string
}

func (k KavenegarAdapter) Send(smsMessage *Message) error {

	api := kavenegar.New(k.ApiKey)

	receptor := []string{smsMessage.to}

	if _, err := api.Message.Send("", receptor, smsMessage.content, nil); err != nil {
		return err
	}

	return nil
}
