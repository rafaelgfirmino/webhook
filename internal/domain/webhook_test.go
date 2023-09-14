package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewWebhook(t *testing.T) {
	id := uuid.New()
	tests := []struct {
		name          string
		input         Webhook
		expected      Webhook
		expectedError error
	}{
		{
			name:          "Invalid Url",
			input:         Webhook{Url: "1", Authorization: "", Id: id},
			expected:      Webhook{TenantId: id, Url: "1", Authorization: "", IsEnabled: true},
			expectedError: InvalidWebhookUrl,
		},
		{
			name:          "Invalid AuthorizationIsRequired",
			input:         Webhook{Url: "http://teste.com.br", Authorization: "", Id: id},
			expected:      Webhook{TenantId: id, Url: "http://teste.com.br", Authorization: "", IsEnabled: true},
			expectedError: InvalidAuthorization,
		},
		{
			name:          "Invalid webhook registry",
			input:         Webhook{Url: "", Authorization: "", Id: id},
			expected:      Webhook{TenantId: id, Url: "", Authorization: "", IsEnabled: true},
			expectedError: InvalidWebhookUrl,
		},
		{
			name:          "Valid Webhook",
			input:         Webhook{Url: "http://teste.ccom.br", Authorization: "123"},
			expected:      Webhook{Url: "http://teste.ccom.br", Authorization: "123", TenantId: id, IsEnabled: true, Id: id},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			webhook, err := NewWebhook(test.input.Url, test.input.Authorization, id)
			test.expected.Id = webhook.Id
			assert.Equal(t, test.expected, webhook)
			if test.expectedError != nil {
				assert.Equal(t, err, test.expectedError)
			}
		})
	}
}
