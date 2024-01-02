package workflows

type WorkflowId uint32

type Workflow struct {
	Id           WorkflowId
	Name         string
	FriendlyName string
}
