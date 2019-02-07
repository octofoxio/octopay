package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"softnet/pkg/agent/proto"
	"time"
)

const (
	AgentTypeSandbox = "SANDBOX"
)

type CashInAgent struct {
	ID       string `gorm:"primary_key"`
	Type     string
	Endpoint string
}
type CashInAgentRepository interface {
	CreatePaymentReference(in *CreatePaymentReferenceInput) (*CreatePaymentReferenceOutput, error)
}
type DefaultCashInAgentRepository struct {
	DB *gorm.DB
}
type CreatePaymentReferenceInput struct {
	PaymentID    string
	CurrencyCode string
	Amount       float64
	Ref1         string
	Ref2         string
	AgentID      string
	ExpireAt     time.Time
}
type CreatePaymentReferenceOutput struct {
	Type   string
	Method int
	Data   interface{}
}

type PaymentReferenceCounter struct {
	Code   string
	Format string
}
type PaymentReferenceRedirect struct {
	PaymentURI string
}

func GetResultFromPaymentAgent(cc *grpc.ClientConn, agent *CashInAgent, in *CreatePaymentReferenceInput) (*proto.CreateNewPaymentReferenceOutput, error) {
	switch agent.Type {
	case AgentTypeSandbox:
		sandboxAgent := proto.NewSandboxAgentClient(cc)
		result, err := sandboxAgent.CreateNewPaymentReference(context.Background(), &proto.CreateNewPaymentReferenceInput{
			Amount:       float32(in.Amount),
			CurrencyCode: in.CurrencyCode,
			Ref1:         in.Ref1,
			PaymentID:    in.PaymentID,
			ExpireAt:     in.ExpireAt.Unix(),
		})
		return result, err
	default:
		return nil, fmt.Errorf("Payment agent %s not supported", agent.Type)
	}
}

func (s *DefaultCashInAgentRepository) CreatePaymentReference(in *CreatePaymentReferenceInput) (*CreatePaymentReferenceOutput, error) {

	var agent CashInAgent

	st := s.DB.Find(&agent, &CashInAgent{
		ID: in.AgentID,
	})
	if st.Error != nil {
		return nil, st.Error
	}

	agentConn, err := grpc.Dial(agent.Endpoint, grpc.WithInsecure())

	agentPaymentResult, err := GetResultFromPaymentAgent(agentConn, &agent, in)

	if err != nil {
		return nil, err
	}

	if agentPaymentResult.Result == nil {
		return nil, fmt.Errorf("Payment agent error")
	}

	switch agentPaymentResult.Result.Type {
	case "barcode":
		return &CreatePaymentReferenceOutput{
			Type: agentPaymentResult.Result.Type,
			Data: PaymentReferenceCounter{
				Code:   agentPaymentResult.Result.Code,
				Format: agentPaymentResult.Result.Format,
			},
		}, st.Error
	case "redirect":
		return &CreatePaymentReferenceOutput{
			Type: agentPaymentResult.Result.Type,
			Data: PaymentReferenceRedirect{
				PaymentURI: agentPaymentResult.Result.PaymentURI,
			},
		}, st.Error
	default:
		return nil, fmt.Errorf("Error: payment method %s not supported", agent.Type)
	}

}
