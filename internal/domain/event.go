package domain

import (
	"github.com/firmino/webhook/pkg"
	"github.com/google/uuid"
)

var (
	ErrorEventNameIsRequired = pkg.WebhookError{
		Reason:     "Field Name is required",
		StatusCode: 412,
		Status:     "Precondition Failed",
	}

	ErrorEventDescriptionIsRequired = pkg.WebhookError{
		Reason:     "Field Description is required",
		StatusCode: 412,
		Status:     "Field Name is required",
	}
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
}

func NewEvent(name string, description string) (Event, error) {
	e := Event{Name: name, Description: description}
	err := e.validate()
	return e, err
}

func (e *Event) validate() error {
	if e.Name == "" {
		return ErrorEventNameIsRequired
	}
	if e.Description == "" {
		return ErrorEventDescriptionIsRequired
	}
	return nil
}
