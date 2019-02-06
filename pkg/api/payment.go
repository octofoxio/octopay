package api

import "softnet/pkg/repository"

type CreateNewPaymentSessionInput repository.CreateInput
type CreateNewPaymentSessionOutput struct {
}
type PaymentService interface {
	CreateNewPaymentSession(input *CreateNewPaymentSessionInput) (*CreateNewPaymentSessionOutput, error)
}

type DefaultPaymentService struct {
	Payment repository.PaymentRepository
}

func (d *DefaultPaymentService) CreateNewPaymentSession(input *CreateNewPaymentSessionInput) (*CreateNewPaymentSessionOutput, error) {

	_, err := d.Payment.Create(&repository.CreateInput{
		Ref2:         input.Ref2,
		Ref1:         input.Ref1,
		CurrencyCode: input.CurrencyCode,
		Amount:       input.Amount,
		AgentID:      "Sandbox",
		ClientID:     input.ClientID,
	})

	if err != nil {
		return nil, err
	}

	return &CreateNewPaymentSessionOutput{}, nil

}
