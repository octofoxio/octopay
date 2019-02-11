package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	proto2 "softnet/pkg/api/proto"
	"softnet/pkg/repository"
)

type NewAPIInput struct {
	Middleware    Middleware
	ClientAppInfo repository.ClientApplicationInformationRepository
	PaymentGRPC   proto2.PaymentServer
	PaymentAPI    PaymentService
	Webhook       WebhookService
}

func NewAPI(in *NewAPIInput) (*gin.Engine, func()) {
	r := gin.Default()

	authenticationResource := r.Group("/auth")
	{
		authenticationResource.POST("/", func(c *gin.Context) {
			type PostAuthBody struct {
				Public string `json:"publicKey"`
				Secret string `json:"secretKey"`
			}
			var body PostAuthBody
			if c.BindJSON(&body) != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if _, err := in.ClientAppInfo.ValidateCredential(&repository.ValidateCredentialInput{
				ID:     body.Public,
				Secret: body.Secret,
			}); err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Status(http.StatusCreated)
			c.Writer.Header().Set("AccessToken", fmt.Sprintf("%s::%s", body.Public, body.Secret))
		})
	}

	webhookResource := r.Group("/webhooks")
	{
		webhookResource.Use(in.Middleware.WithValidateClientCredentials)
		webhookResource.Use(in.Middleware.WithSessionLogger)

		webhookResource.POST("/", func(c *gin.Context) {
			logger, _ := GetLogger(c)
			session := MustGetClient(c)

			type Body struct {
				URL         string `json:"url"`
				HTTPHeaders []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"httpHeaders"`
			}
			var body Body
			err := c.BindJSON(&body)

			err = in.Webhook.RegisterWebhookByClientID(&RegisterWebhookByClientIDInput{
				WebhookURL: body.URL,
				ClientID:   session.ID,
			})

			if err != nil {
				logger.Error(err)
			}

		})
	}

	paymentResource := r.Group("/sources")
	{
		paymentResource.Use(in.Middleware.WithValidateClientCredentials)
		paymentResource.POST("/", func(c *gin.Context) {

			session := MustGetClient(c)

			type PostPaymentBody struct {
				PaymentChannel string `json:"paymentChannel"`
				Amount         int    `json:"amount"`
				Currency       string `json:"currency"`
				RequestPayload struct {
					Reference     string        `json:"reference"`
					ReferenceDate string        `json:"referenceDate"`
					ProductList   []interface{} `json:"productList"`
					FirstName     string        `json:"firstName"`
					LastName      string        `json:"lastName"`
					Email         string        `json:"email"`
					CreditData    string        `json:"creditData"`
					Remark        string        `json:"remark"`
				} `json:"requestPayload"`
			}

			var body PostPaymentBody
			err := c.BindJSON(&body)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			referencePayload, err := json.Marshal(body.RequestPayload)
			output, err := in.PaymentAPI.CreateNewPaymentSession(&CreateNewPaymentSessionInput{
				ClientID:     session.ID,
				AgentID:      body.PaymentChannel,
				Amount:       float64(body.Amount) / 100,
				CurrencyCode: body.Currency,
				Ref1:         string(referencePayload),
			})
			if err != nil {
				err := c.AbortWithError(http.StatusServiceUnavailable, err)
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
				return
			}

			type PostPaymentResult struct {
				Status string      `json:"status"`
				Data   interface{} `json:"data"`
			}
			c.JSON(http.StatusCreated, &PostPaymentResult{
				Status: "success",
				Data:   output,
			})

		})
	}
	return r, func() {
		s := grpc.NewServer()
		proto2.RegisterPaymentServer(s, in.PaymentGRPC)
		reflection.Register(s)
		lis, err := net.Listen("tcp",
			fmt.Sprintf(":%s", "3009"))
		if err != nil {
			panic(err)
		}
		fmt.Printf("PaymentModel GRPC Running on :%s\n", "3009")
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}

	}
}
