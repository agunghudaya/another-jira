package event

import "time"

// Event represents a domain event
type Event interface {
	// EventName returns the name of the event
	EventName() string
	
	// OccurredOn returns when the event occurred
	OccurredOn() time.Time
}

// EventHandler defines the interface for event handlers
type EventHandler interface {
	// Handle processes an event
	Handle(event Event) error
}

// EventBus defines the interface for the event bus
type EventBus interface {
	// Publish publishes an event to all subscribers
	Publish(event Event) error
	
	// Subscribe subscribes to events of a specific type
	Subscribe(eventName string, handler EventHandler) error
	
	// Unsubscribe unsubscribes from events
	Unsubscribe(eventName string, handler EventHandler) error
}

// BaseEvent provides a base implementation of the Event interface
type BaseEvent struct {
	Name      string
	Timestamp time.Time
}

// NewBaseEvent creates a new base event
func NewBaseEvent(name string) BaseEvent {
	return BaseEvent{
		Name:      name,
		Timestamp: time.Now(),
	}
}

// EventName implements the Event interface
func (e BaseEvent) EventName() string {
	return e.Name
}

// OccurredOn implements the Event interface
func (e BaseEvent) OccurredOn() time.Time {
	return e.Timestamp
}

// JiraIssueSyncedEvent represents the event when a Jira issue is synced
type JiraIssueSyncedEvent struct {
	BaseEvent
	IssueKey string
}

// NewJiraIssueSyncedEvent creates a new JiraIssueSyncedEvent
func NewJiraIssueSyncedEvent(issueKey string) JiraIssueSyncedEvent {
	return JiraIssueSyncedEvent{
		BaseEvent: NewBaseEvent("jira.issue.synced"),
		IssueKey:  issueKey,
	}
} 