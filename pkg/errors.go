package pkg

type WebhookError struct {
	StatusCode int
	Status     string
	Reason     string
}

// Error returns the error message.
func (e WebhookError) Error() string {
	if e.Reason != "" {
		return e.Reason
	}
	return "Webhooks Error"
}
