package api

import "softnet/pkg/repository"

type RegisterWebhookByClientIDInput struct {
	ClientID string
	WebhookURL string
}
type WebhookService interface {
	RegisterWebhookByClientID(in *RegisterWebhookByClientIDInput) error
}

type DefaultWebhookService struct {
	client repository.ClientApplicationInformationRepository
}

func (d *DefaultWebhookService) RegisterWebhookByClientID(in *RegisterWebhookByClientIDInput) (error) {
	err := d.client.RegisterCallbackURL(&repository.RegisterCallbackURLInput{
		ID:in.ClientID,
		Url:in.WebhookURL,
	})
	return err
}
