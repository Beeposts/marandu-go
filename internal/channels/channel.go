package channels

type ChannelId uint16

type Channel struct {
	Id            ChannelId
	Name          string
	Configuration string
	ChannelType   ChannelType
	ChannelVendor uint8
}
