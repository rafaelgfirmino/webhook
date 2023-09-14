package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	tests := []struct {
		name          string
		input         Event
		expected      Event
		expectedError error
	}{
		{
			name:          "Invalid Name",
			input:         Event{Name: "", Description: "Teste"},
			expected:      Event{Name: "", Description: "Teste"},
			expectedError: ErrorEventNameIsRequired,
		},
		{
			name:          "Invalid Description",
			input:         Event{Name: "Teste", Description: ""},
			expected:      Event{Name: "Teste", Description: ""},
			expectedError: ErrorEventDescriptionIsRequired,
		},
		{
			name:          "Valid Event",
			input:         Event{Name: "Teste", Description: "Teste"},
			expected:      Event{Name: "Teste", Description: "Teste"},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event, err := NewEvent(test.input.Name, test.input.Description)
			assert.Equal(t, test.expected, event)
			if test.expectedError != nil {
				assert.Equal(t, err, test.expectedError)
			}
		})
	}
}
