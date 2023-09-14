package domain

import (
	"context"
	"net/url"

	"github.com/firmino/webhook/pkg"
	"github.com/google/uuid"
)

type Webhook struct {
	Id            uuid.UUID
	Authorization string
	Url           string
	Description   string
	IsEnabled     bool
	TenantId      uuid.UUID
}

var (
	InvalidWebhookUrl = pkg.WebhookError{
		Reason:     "Url is not valid",
		StatusCode: 412,
		Status:     "webhook-url",
	}

	InvalidAuthorization = pkg.WebhookError{
		Reason:     "Authorization is not valid",
		StatusCode: 412,
		Status:     "webhook-authorization",
	}
)

func NewWebhook(url, authorization string, tenantId uuid.UUID) (Webhook, error) {
	w := Webhook{
		Id:            uuid.New(),
		Url:           url,
		Authorization: authorization,
		IsEnabled:     true,
		TenantId:      tenantId,
	}
	err := w.validate()
	return w, err
}

func (w *Webhook) validate() error {
	_, err := url.ParseRequestURI(w.Url)
	if err != nil {
		return InvalidWebhookUrl
	}
	if w.Authorization == "" {
		return InvalidAuthorization
	}
	return nil
}

type WebhookRepository interface {
	Insert(ctx context.Context, webhook Webhook) (*Webhook, error)
	FindById(ctx context.Context, id uuid.UUID) (*Webhook, error)
	FindAll(ctx context.Context, pagination *pkg.Pagination) (*pkg.Pagination, error)
}

type WebhookUseCase interface {
	Save(ctx context.Context, webhook Webhook) (*Webhook, error)
	GetById(ctx context.Context, id uuid.UUID) (*Webhook, error)
	GetAll(ctx context.Context, pagination *pkg.Pagination) (*pkg.Pagination, error)
}
