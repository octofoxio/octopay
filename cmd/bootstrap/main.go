package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	url2 "net/url"
	"softnet/cmd/sandbox/lib"
	"softnet/pkg/api"
	"softnet/pkg/repository"
	"time"
)

func main() {
	db, _ := repository.NewGORMSqlLiteConnection("./db.sqlite")
	repository.DefaultGenesis(db)
	clientAppInfo := repository.DefaultClientApplicationInformationRepository{
		DB: db.Debug(),
	}
	payment := repository.DefaultPaymentRepository{
		DB: db.Debug(),
	}
	agent := repository.DefaultCashInAgentRepository{
		DB: db.Debug(),
	}
	paymentService := api.DefaultPaymentService{
		Payment: &payment,
		Agent:   &agent,
	}

	m := api.DefaultMiddleware{
		ClientAppInfo: &clientAppInfo,
	}

	server, grpc := api.NewAPI(&api.NewAPIInput{
		Middleware:    &m,
		ClientAppInfo: &clientAppInfo,
		PaymentAPI:    &paymentService,
		PaymentGRPC:   &paymentService,
	})

	go func() {
		target := "http://localhost:3001"
		url, _ := url2.Parse(target)
		handler := httputil.NewSingleHostReverseProxy(url)
		handler.FlushInterval = 100 * time.Millisecond
		server.NoRoute(func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		})
	}()
	go func() {
		lib.StartSandboxServer(db)
	}()

	go grpc()

	if e := server.Run(":3000"); e != nil {
		log.Fatal(e)
	}

}
