package eventQueue

import "time"

type Event struct {
	// Time of event
	date time.Time

	// timeout limit to fire event
	timeout time.Duration

	// event name
	label string

	// event active
	active bool

	data map[any]any
}

func (e Event) GetDate() (date time.Time) {
	return e.date
}

func (e Event) GetTimeout() (timeout time.Duration) {
	return e.timeout
}

func (e Event) GetLabel() (label string) {
	return e.label
}

func (e Event) GetActive() (active bool) {
	return e.active
}

func (e Event) GetData() (data map[any]any) {
	return e.data
}

// Set
//
// English:
//
// # Set a new event
//
// Português:
//
// Define um novo evento
func (e *Event) Set(label string, active bool, data map[any]any, timeout time.Duration) {
	e.date = time.Now()
	e.timeout = timeout
	e.label = label
	e.active = active
	e.data = data
}
