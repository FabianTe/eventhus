package eventhus

import "context"

// EventStore saves the events from an aggregate
type EventStore interface {
	// Deprecated: Future versions will change this functions signature to the signature of SaveWithContext.
	// You should use SaveWithContext for now until it is renamed to Load
	Save(events []Event, version int) error
	SaveWithContext(ctx context.Context, events []Event, version int) error

	// Deprecated: Future versions will change this functions signature to the signature of SafeSaveWithContext.
	// You should use SafeSaveWithContext for now until it is renamed to Load
	SafeSave(events []Event, version int) error
	SafeSaveWithContext(ctx context.Context, events []Event, version int) error

	Load(aggregateID string) ([]Event, error)
	// Deprecated: Future versions will change this functions signature to the signature of LoadWithContext.
	// You should use LoadWithContext for now until it is renamed to Load
	LoadWithContext(ctx context.Context, aggregateID string) ([]Event, error)
}
