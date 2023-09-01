package domain

import (
	"github.com/firmino/webhook/pkg"
	"github.com/google/uuid"
)

const ErrorEventNameIsRequired = "Field Name is required"
const ErrorEventDescriptionIsRequired = "Field Description is required"

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
		return pkg.NewCustomError(pkg.PreconditionFailed, ErrorEventNameIsRequired)
	}
	if e.Description == "" {
		return pkg.NewCustomError(pkg.PreconditionFailed, ErrorEventDescriptionIsRequired)
	}
	return nil
}
