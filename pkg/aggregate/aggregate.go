package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type Now func() time.Time

type Version string

func (v Version) String() string {
	return string(v)
}

func NewVersion() Version {
	return Version(uuid.New().String())
}

type eventRecorder struct {
	events []interface{}
}

func (r *eventRecorder) Record(event interface{}) {
	r.events = append(r.events, event)
}

func (r *eventRecorder) Events() []interface{} {
	return r.events
}

func (r *eventRecorder) Clear() {
	r.events = []interface{}{}
}

type Root struct {
	eventRecorder eventRecorder
}

func (r *Root) AddEvent(event interface{}) {
	r.eventRecorder.Record(event)
}

func (r *Root) Events() []interface{} {
	return r.eventRecorder.Events()
}

func (r *Root) ClearEvents() {
	r.eventRecorder.Clear()
}
