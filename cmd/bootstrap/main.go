package main

import (
	"log"
	"softnet/cmd/sandbox/lib"
	"softnet/pkg/api"
	"softnet/pkg/repository"
)

func main() {
	db, _ := repository.NewGORMSqlLiteConnection("./db.sqlite")
	repository.DefaultGenesis(db)
	clientAppInfo := repository.DefaultClientApplicationInformationRepository{
		DB: db,
	}
	payment := repository.DefaultPaymentRepository{
		DB: db,
	}
	agent := repository.DefaultCashInAgentRepository{
		DB: db,
	}
	paymentService := api.DefaultPaymentService{
		Payment: &payment,
		Agent:   &agent,
	}

	m := api.DefaultMiddleware{
		ClientAppInfo: &clientAppInfo,
	}

	server := api.NewAPI(&api.NewAPIInput{
		Middleware:     &m,
		ClientAppInfo:  &clientAppInfo,
		PaymentService: &paymentService,
	})

	go func() {
		lib.StartSandboxServer(db)
	}()

	if e := server.Run(":3000"); e != nil {
		log.Fatal(e)
	}

}
