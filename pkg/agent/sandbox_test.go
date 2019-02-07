package agent

import (
	"context"
	"softnet/pkg/agent/proto"
	"softnet/pkg/repository"
	"testing"
	"time"
)

func TestSandboxCounterAgent_CreateNewPaymentReference(t *testing.T) {

	db, _ := repository.NewGORMSqlLiteConnection("./db.sqlite")
	db.AutoMigrate(&CashInTransaction{})

	sandbox := SandboxCounterAgent{
		DB: db,
	}

	out, _ := sandbox.CreateNewPaymentReference(context.Background(), &proto.CreateNewPaymentReferenceInput{
		Amount:       100.50,
		PaymentID:    "TEST",
		CurrencyCode: "THB",
		ExpireAt:     time.Now().Unix(),
		Ref1:         "REFERENCE",
	})
	if len(out.Error) > 0 {
		t.Error(out.Error)
	}
	t.Log(out.Result.Code)
}
