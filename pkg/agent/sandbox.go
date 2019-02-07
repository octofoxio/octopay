package agent

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"softnet/pkg/agent/proto"
	"time"
)

type CashInTransaction struct {
	Code      string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Type      string
	Format    string
	Currency  string
	Amount    float64
	PaymentID string
}

type SandboxCounterAgent struct {
	DB *gorm.DB
}

func (s *SandboxCounterAgent) CreateNewPaymentReference(c context.Context, in *proto.CreateNewPaymentReferenceInput) (*proto.CreateNewPaymentReferenceOutput, error) {
	now := time.Now()
	ID := fmt.Sprintf("%s::%f::%d", in.PaymentID, in.Amount, now.Unix())
	hash := md5.New()
	hash.Write([]byte(ID))
	code := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	if err := s.DB.Create(&CashInTransaction{
		Type:      "barcode",
		Amount:    float64(in.Amount),
		Format:    "code128b",
		PaymentID: in.PaymentID,
		Currency:  in.CurrencyCode,
		Code:      code,
		CreatedAt: now,
	}).Error; err != nil {
		return &proto.CreateNewPaymentReferenceOutput{
			Error: err.Error(),
		}, nil
	}

	return &proto.CreateNewPaymentReferenceOutput{
		Result: &proto.CreateNewPaymentReferenceOutput_Output{
			Code:   code,
			Format: "code128b",
			Type:   "barcode",
		},
	}, nil

}

func NewSandboxCounterAgent(db *gorm.DB) *SandboxCounterAgent {

	db.AutoMigrate(&CashInTransaction{})
	return &SandboxCounterAgent{
		DB: db,
	}

}
