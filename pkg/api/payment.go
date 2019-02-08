package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"softnet/pkg/api/proto"
	"softnet/pkg/repository"
)

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

func (d *DefaultPaymentService) GetTransactionList(c context.Context, in *proto.GetTransactionListInput) (*proto.GetTransactionListOutput, error) {

	out, err := d.Payment.GetPayments(&repository.GetPaymentsInput{
		Offset: int(in.Offset),
		Limit:  int(in.Limit),
	})

	var result = []*proto.Transaction{}
	if err != nil {
		return nil, err
	}
	for _, payment := range out.Result {

		histories := []*proto.History{}
		for _, history := range payment.History {
			histories = append(histories, &proto.History{
				Status: history.Status,
				Memo:   history.Memo,
			})
		}
		result = append(result, &proto.Transaction{
			ID:              payment.ID.String(),
			Currency:        payment.CurrencyCode,
			Amount:          float32(payment.Amount),
			History:         histories,
			CashInReference: payment.CashInReference,
			CashInType:      payment.CashInType,
			PaymentProvider: &proto.PaymentAgent{
				ID:   payment.CashInAgent.ID,
				Type: payment.CashInAgent.Type,
			},
		})
	}

	return &proto.GetTransactionListOutput{
		Transactions: result,
	}, nil
}

func (d *DefaultPaymentService) ConfirmPaymentCashIn(context.Context, *proto.ConfirmPaymentCashInInput) (*proto.ConfirmPaymentCashInOutput, error) {
	return &proto.ConfirmPaymentCashInOutput{}, nil
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
	var history *repository.CommitHistoryInput
	if err != nil {
		history = &repository.CommitHistoryInput{
			PaymentID: paymentRecord.ID.String(),
			Status:    repository.PaymentStatusFailed,
			Memo:      err.Error(),
		}
	} else {

		ref, err := json.Marshal(cashInRef.Data)
		if err != nil {
			return nil, err
		}
		if _, err := d.Payment.UpdateCashInInformation(&repository.UpdateCashInInformationInput{
			CashInReference: string(ref),
			CashInType:      cashInRef.Type,
			ID:              paymentRecord.ID,
		}); err != nil {
			return nil, err
		}
		history = &repository.CommitHistoryInput{
			PaymentID: paymentRecord.ID.String(),
			Status:    repository.PaymentStatusReadyToCashIn,
			Memo:      fmt.Sprintf("Response from %s", paymentRecord.CashInAgentID),
		}
	}

	if err := d.Payment.CommitHistory(history); err != nil {
		return nil, err
	}

	return &CreateNewPaymentSessionOutput{
		ID:   paymentRecord.ID.String(),
		Type: cashInRef.Type,
		Data: cashInRef.Data,
	}, nil

}
