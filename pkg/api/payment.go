package api

import "softnet/pkg/repository"

type CreateNewPaymentSessionInput repository.CreateInput
type CreateNewPaymentSessionOutput struct {
	ID   string      `json:"id"`
	Type string      `json:"type"`
	Data interface{} `json:"responsePayload"`
}
type PaymentService interface {
	CreateNewPaymentSession(input *CreateNewPaymentSessionInput) (*CreateNewPaymentSessionOutput, error)
}

type DefaultPaymentService struct {
	Payment repository.PaymentRepository
	Agent   repository.CashInAgentRepository
}

func (d *DefaultPaymentService) CreateNewPaymentSession(input *CreateNewPaymentSessionInput) (*CreateNewPaymentSessionOutput, error) {

	paymentRecord, err := d.Payment.Create(&repository.CreateInput{
		Ref1:         input.Ref1,
		CurrencyCode: input.CurrencyCode,
		Amount:       input.Amount,
		AgentID:      input.AgentID,
		ClientID:     input.ClientID,
	})

	if err != nil {
		return nil, err
	}

	cashInRef, err := d.Agent.CreatePaymentReference(&repository.CreatePaymentReferenceInput{
		PaymentID:    paymentRecord.ID.String(),
		Ref1:         paymentRecord.Ref1,
		CurrencyCode: paymentRecord.CurrencyCode,
		Amount:       paymentRecord.Amount,
		AgentID:      paymentRecord.CashInAgentID,
	})

	if err != nil {

		if err := d.Payment.CommitHistory(&repository.CommitHistoryInput{
			PaymentID: paymentRecord.ID.String(),
			Status:    repository.PaymentStatusFailed,
			Memo:      err.Error(),
		}); err != nil {
			return nil, err
		}

		return nil, err
	}

	return &CreateNewPaymentSessionOutput{
		ID:   paymentRecord.ID.String(),
		Type: cashInRef.Type,
		Data: cashInRef.Data,
	}, nil

}
