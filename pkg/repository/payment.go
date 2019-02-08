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
	PaymentStatusFailed        = "FAILED"
)

type PaymentStatusHistory struct {
	gorm.Model
	PaymentID uuid.UUID
	Status    string
	Memo      string
}

type PaymentModel struct {
	ID        uuid.UUID `gorm:"type:uuid; primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ExpireAt      time.Time
	History       []PaymentStatusHistory `gorm:"ForeignKey:PaymentID;AssociationForeignKey:ID"`
	CashInAgentID string                 `gorm:"type:uuid;not null"`
	CashInAgent   CashInAgent

	CurrencyCode string  `gorm:"not null"`
	Amount       float64 `gorm:"not null"`
	Ref1         sql.NullString
	Ref2         sql.NullString

	CashInReference string
	CashInType      string

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
	Status    string
	Memo      string
}
type UpdateCashInInformationInput struct {
	ID              *uuid.UUID
	CashInReference string
	CashInType      string
}
type UpdateCashInInformationOutput struct {
}
type GetPaymentsInput struct {
	Limit  int
	Offset int
}
type GetPaymentsOutput struct {
	Result []PaymentModel
}
type PaymentRepository interface {
	Create(*CreateInput) (*CreateOutput, error)
	UpdateCashInInformation(*UpdateCashInInformationInput) (*UpdateCashInInformationOutput, error)
	CommitHistory(input *CommitHistoryInput) error
	GetPayments(input *GetPaymentsInput) (*GetPaymentsOutput, error)
}

type DefaultPaymentRepository struct {
	DB *gorm.DB
}

func (d *DefaultPaymentRepository) GetPayments(input *GetPaymentsInput) (*GetPaymentsOutput, error) {
	var result []PaymentModel
	err := d.DB.Debug().
		Preload("History").
		Preload("CashInAgent").
		Limit(input.Limit).
		Offset(input.Offset).
		Order("created_at desc").
		Find(&result, &PaymentModel{}).
		Error
	return &GetPaymentsOutput{
		Result: result,
	}, err
}

func (d *DefaultPaymentRepository) UpdateCashInInformation(input *UpdateCashInInformationInput) (*UpdateCashInInformationOutput, error) {

	if err := d.DB.Debug().
		Model(&PaymentModel{}).
		Update(&PaymentModel{
			ID:              *input.ID,
			CashInType:      input.CashInType,
			CashInReference: input.CashInReference,
		},
		).Error; err != nil {
		return nil, err
	}

	return &UpdateCashInInformationOutput{}, nil
}

func (d *DefaultPaymentRepository) CommitHistory(input *CommitHistoryInput) error {
	st := d.DB.Create(&PaymentStatusHistory{
		PaymentID: uuid.FromStringOrNil(input.PaymentID),
		Memo:      input.Memo,
		Status:    input.Status,
	})
	return st.Error
}

func (d *DefaultPaymentRepository) Create(in *CreateInput) (*CreateOutput, error) {

	if in.ExpireAt.IsZero() {
		in.ExpireAt = time.Now().Local().Add(1 * time.Hour)
	}

	ID := uuid.NewV4()
	var result PaymentModel
	st := d.DB.Debug().
		Create(
			&PaymentModel{
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
		Find(&result, &PaymentModel{
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
