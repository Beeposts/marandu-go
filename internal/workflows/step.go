package workflows

import (
	"github.com/beeposts/marandu-go/internal/channels"
)

type StepId uint32

type Step struct {
	Id         StepId
	Name       string
	ChannelId  channels.ChannelId
	Order      uint8
	WorkflowId WorkflowId
}
