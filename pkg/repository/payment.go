package repository

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

const (
	PaymentStatusInitial       = "INITIAL"
	PaymentStatusReadyToCashIn = "READY_TO_CASH_IN"
	PaymentStatusCashInConfirm = "CASH_IN_CONFIRM"
)

type PaymentStatusHistory struct {
	gorm.Model
	PaymentID uuid.UUID
	Status    string
	Memo      string
}

type Payment struct {
	ID        uuid.UUID `gorm:"type:uuid; primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ExpireAt      time.Time
	History       []PaymentStatusHistory
	CashInAgentID string `gorm:"type:uuid;not null"`
	CashInAgent   CashInAgent

	CurrencyCode string  `gorm:"not null"`
	Amount       float64 `gorm:"not null"`
	Ref1         sql.NullString
	Ref2         sql.NullString

	CashInReference string

	ClientID string
	Client   ClientApplicationInformation
}

type CreateInput struct {
	CurrencyCode string
	Amount       float64
	Ref1         string
	Ref2         string
	AgentID      string
	ClientID     string
	ExpireAt     time.Time
}
type CreateOutput struct {
	ID            *uuid.UUID
	CashInAgentID string
	CurrencyCode  string  `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Ref1          string
	Ref2          string
}
type CommitHistoryInput struct {
	PaymentID string
	History   PaymentStatusHistory
}
type PaymentRepository interface {
	Create(*CreateInput) (*CreateOutput, error)
	CommitHistory(input *CommitHistoryInput) error
}

type DefaultPaymentRepository struct {
	DB *gorm.DB
}

func (d *DefaultPaymentRepository) CommitHistory(input *CommitHistoryInput) error {
	st := d.DB.Create(&PaymentStatusHistory{
		PaymentID: uuid.FromStringOrNil(input.PaymentID),
		Memo:      input.History.Memo,
	})
	return st.Error
}

func (d *DefaultPaymentRepository) Create(in *CreateInput) (*CreateOutput, error) {

	if in.ExpireAt.IsZero() {
		in.ExpireAt = time.Now().Local().Add(1 * time.Hour)
	}

	ID := uuid.NewV4()
	var result Payment
	st := d.DB.
		Create(
			&Payment{
				ClientID:     in.ClientID,
				ID:           ID,
				CurrencyCode: in.CurrencyCode,
				Amount:       in.Amount,
				Ref1: sql.NullString{
					String: in.Ref1,
					Valid:  len(in.Ref1) > 0,
				},
				Ref2: sql.NullString{
					String: in.Ref2,
					Valid:  len(in.Ref2) > 0,
				},
				History: []PaymentStatusHistory{
					{
						Status: PaymentStatusInitial,
						Memo:   "Initial by Create new payment",
					},
				},
				CashInAgentID: in.AgentID,
				ExpireAt:      in.ExpireAt,
			},
		).
		Find(&result, &Payment{
			ID: ID,
		})

	return &CreateOutput{
		ID:            &ID,
		CashInAgentID: result.CashInAgentID,
		Amount:        result.Amount,
		CurrencyCode:  result.CurrencyCode,
		Ref1:          result.Ref1.String,
		Ref2:          result.Ref2.String,
	}, st.Error
}
