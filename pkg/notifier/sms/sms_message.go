package sms

import "fmt"

type Message struct {
	to      string
	content string
}

func NewSmsMessage() *Message {
	return &Message{}
}

func (n *Message) SetReceiverNumber(number string) *Message {
	n.to = number
	return n
}

func (n *Message) SetContent(content string, data ...any) *Message {
	n.content = fmt.Sprintf(content, data...)
	return n
}

func (n *Message) BuildCustomContent(t Template, data ...any) *Message {
	n.content = fmt.Sprintf(string(t), data...)
	return n
}
