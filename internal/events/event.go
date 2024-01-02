package events

import "github.com/beeposts/marandu-go/internal/workflows"

type EventId uint64

type Event struct {
	Id         EventId
	WorkflowId workflows.WorkflowId
	Data       string
}
