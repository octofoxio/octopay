package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

const (
	_ = iota
	PaymentMethodRedirect
	PaymentMethodCounter
)

type CashInAgent struct {
	ID            string `gorm:"primary_key"`
	Endpoint      string
	PaymentMethod int
}
type CashInAgentRepository struct {
	DB *gorm.DB
}
type CreatePaymentReferenceInput struct {
	PaymentID    string
	CurrencyCode string
	Amount       float64
	Ref1         string
	Ref2         string
	AgentID      string
	ClientID     string
	ExpireAt     time.Time
}
type CreatePaymentReferenceOutput struct {
	Method    string
	Reference interface{}
}

type PaymentReferenceCounter struct {
	Type string
	Code string
	Memo string
}
type PaymentReferenceRedirect struct {
	PaymentURI string
}

func (s *CashInAgentRepository) CreatePaymentReference(in *CreatePaymentReferenceInput) (*CreatePaymentReferenceOutput, error) {
	var agent CashInAgent

	st := s.DB.Find(&agent, &CashInAgent{
		ID: in.AgentID,
	})
	if st.Error != nil {
		return nil, st.Error
	}

	type PaymentAgentRequestPayload struct {
		PaymentID    string
		CurrencyCode string
		Amount       float64
		Ref1         string
		Ref2         string
		ExpireAt     time.Time
	}
	resp, err := resty.R().
		SetBody(PaymentAgentRequestPayload{
			PaymentID:    in.PaymentID,
			ExpireAt:     in.ExpireAt,
			Amount:       in.Amount,
			CurrencyCode: in.CurrencyCode,
			Ref1:         in.Ref1,
			Ref2:         in.Ref2,
		}).
		Post(agent.Endpoint)
	if err != nil {
		return nil, err
	} else if resp.StatusCode() > http.StatusBadRequest {
		return nil, fmt.Errorf("Payment from %s return %s", agent.ID, http.StatusText(resp.StatusCode()))
	}

	type PaymentAgentResponsePayload struct {
		PaymentURI string
		Type       string
		Code       string
		Memo       string
	}
	var result PaymentAgentResponsePayload
	err = json.Unmarshal(resp.Body(), &result)

	if err != nil {
		return nil, err
	}

	var reference interface{}
	switch agent.PaymentMethod {
	case PaymentMethodCounter:
		reference = PaymentReferenceCounter{
			Code: result.Code,
			Type: result.Type,
			Memo: result.Memo,
		}
		break
	case PaymentMethodRedirect:
		reference = PaymentReferenceRedirect{
			PaymentURI: result.PaymentURI,
		}
		break
	}

	return &CreatePaymentReferenceOutput{
		Reference: reference,
	}, st.Error
}
