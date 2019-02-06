package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"softnet/pkg/repository"
	"time"
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
				CurrencyCode string
				Amount       float64
				Ref1         string
				Ref2         string
				AgentID      string
				ExpireAt     time.Time
			}

			var body PostPaymentBody
			err := c.BindJSON(&body)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			output, err := in.PaymentService.CreateNewPaymentSession(&CreateNewPaymentSessionInput{
				ClientID: session.ID,
				AgentID:  body.AgentID,
			})
			if err != nil {
				c.AbortWithError(http.StatusServiceUnavailable, err)
				return
			}

			c.JSON(http.StatusCreated, output)

		})
	}
	return r
}
