package channels

type ChannelType uint8

const (
	Email ChannelType = iota
	InApp
	Push
	Sms
	Chat
	Webhook
)
