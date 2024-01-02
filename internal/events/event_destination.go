package events

type EventDestinationId uint64

type EventDestination struct {
	Id           EventDestinationId
	SubscriberId string
	Name         string
	LastName     string
	Email        string
	Phone        string
	Group        string
	EventId      EventId
}
