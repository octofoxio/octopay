package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"softnet/pkg/api/proto"
	"time"
)

var (
	PaymentStatusInitial                = proto.PaymentStatus_INITIAL.String()
	PaymentStatusReadyToCashIn          = proto.PaymentStatus_READY_TO_CASH_IN.String()         // Payment gateway ตอบกลับมาแล้ว
	PaymentStatusCashInConfirm          = proto.PaymentStatus_CASH_IN_CONFIRM.String()          // มีเงินเข้ามาแล้ว
	PaymentStatusCallbackAttemptSuccess = proto.PaymentStatus_CALLBACK_ATTEMPT_SUCCESS.String() // noti กลับไปที่ App client แล้ว
	PaymentStatusCallbackAttemptFailed  = proto.PaymentStatus_CALLBACK_ATTEMPT_FAILED.String()  // noti กลับไปที่ app client แล้ว แต่ failed
	PaymentStatusFailed                 = proto.PaymentStatus_FAILED.String()
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
type CommitHistoryOutput struct {
	History []PaymentStatusHistory
}
type UpdateCashInInformationInput struct {
	ID              *uuid.UUID
	CashInReference string
	CashInType      string
}
type UpdateCashInInformationOutput struct {
	Result PaymentModel
}
type GetPaymentsInput struct {
	Limit  int
	Offset int
}
type GetPaymentsOutput struct {
	Result []PaymentModel
}
type GetPaymentInput struct {
	ID string
}
type GetPaymentOutput struct {
	Result *PaymentModel
}
type PaymentRepository interface {
	Create(*CreateInput) (*CreateOutput, error)
	UpdateCashInInformation(*UpdateCashInInformationInput) (*UpdateCashInInformationOutput, error)
	CommitHistory(input *CommitHistoryInput) (*CommitHistoryOutput, error)
	GetPayments(input *GetPaymentsInput) (*GetPaymentsOutput, error)
	GetPayment(input *GetPaymentInput) (*GetPaymentOutput, error)
}

type DefaultPaymentRepository struct {
	DB *gorm.DB
}

func (d *DefaultPaymentRepository) GetPayment(input *GetPaymentInput) (*GetPaymentOutput, error) {
	var result PaymentModel
	ID, err := uuid.FromString(input.ID)
	if err != nil {
		return &GetPaymentOutput{}, err
	}
	err = d.DB.Debug().
		Preload("History").
		Preload("CashInAgent").
		Order("created_at desc").
		Find(&result, &PaymentModel{
			ID: ID,
		}).Error
	if gorm.IsRecordNotFoundError(err) {
		return &GetPaymentOutput{}, nil
	} else if err != nil {
		return &GetPaymentOutput{}, err
	} else {
		return &GetPaymentOutput{
			Result: &result,
		}, nil
	}
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

	var result PaymentModel
	if err := d.DB.Debug().
		Model(&PaymentModel{}).
		Update(&PaymentModel{
			ID:              *input.ID,
			CashInType:      input.CashInType,
			CashInReference: input.CashInReference,
		}).
		Find(&result, &PaymentModel{
			ID: *input.ID,
		}).Error; err != nil {
		return nil, err
	}

	return &UpdateCashInInformationOutput{
		Result: result,
	}, nil
}

func (d *DefaultPaymentRepository) CommitHistory(input *CommitHistoryInput) (*CommitHistoryOutput, error) {
	var history []PaymentStatusHistory
	ID := uuid.FromStringOrNil(input.PaymentID)
	if ID == uuid.Nil {
		return nil, fmt.Errorf("Invalid Payment ID")
	}
	st := d.DB.
		Create(&PaymentStatusHistory{
			PaymentID: ID,
			Memo:      input.Memo,
			Status:    input.Status,
		}).
		Find(&history, &PaymentStatusHistory{
			PaymentID: ID,
		})
	return &CommitHistoryOutput{}, st.Error
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
