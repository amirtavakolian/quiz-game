package sms

import (
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/configloader"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/service/appservice"
)

const (
	kavenegarProvider = "kavenegar"
	smsirProvider     = "smsir"
)

type SMSNotifier struct {
	provider SMSProvider
}

func NewNotifier() SMSNotifier {
	cfgLoader := configloader.NewConfigLoader()
	smsProvider := appservice.NewAppService(cfgLoader)
	currentSmsProvider, apiKey := smsProvider.GetSmsProvider()
	var n SMSNotifier

	switch currentSmsProvider {
	case kavenegarProvider:
		n = SMSNotifier{provider: KavenegarAdapter{ApiKey: apiKey}}
	case smsirProvider:
		n = SMSNotifier{provider: SmsirAdapter{ApiKey: apiKey}}
	default:
		panic("unknown SMS provider: " + currentSmsProvider)
	}
	
	return n
}

func (n *SMSNotifier) SendSMS(smsMessage *Message) error {
	return n.provider.Send(smsMessage)
}
