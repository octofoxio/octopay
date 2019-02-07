package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http"
	"softnet/pkg/repository"
)

type NewAPIInput struct {
	Middleware     Middleware
	ClientAppInfo  repository.ClientApplicationInformationRepository
	PaymentService PaymentService
}

func NewAPI(in *NewAPIInput) *gin.Engine {
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

	paymentResource := r.Group("/payment")
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
			output, err := in.PaymentService.CreateNewPaymentSession(&CreateNewPaymentSessionInput{
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
	return r
}
