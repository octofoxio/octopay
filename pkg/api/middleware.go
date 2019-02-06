package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
	"softnet/pkg/repository"
	"strings"
)

type Middleware interface {
	WithValidateClientCredentials(c *gin.Context)
	WithSessionLogger(c *gin.Context)
}

type DefaultMiddleware struct {
	ClientAppInfo repository.ClientApplicationInformationRepository
}

const (
	MiddlewareKeyClientSession = "client-session"
	MiddlewareKeyClientLogger  = "client-logger"
)

type ClientSession struct {
	ID   string
	Name string
}

func (m *DefaultMiddleware) WithSessionLogger(c *gin.Context) {
	if _, exists := c.Get(MiddlewareKeyClientSession); !exists {
		m.WithValidateClientCredentials(c)
	}

	session := c.MustGet(MiddlewareKeyClientSession).(ClientSession)
	c.Set(MiddlewareKeyClientLogger, log.WithFields(log.Fields{
		"ClientID": session.ID,
	}))
}

func (m *DefaultMiddleware) WithValidateClientCredentials(c *gin.Context) {
	token := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	tokens := strings.Split(token[1], "::")
	if len(tokens) == 2 {
		output, err := m.ClientAppInfo.ValidateCredential(&repository.ValidateCredentialInput{
			ID:     tokens[0],
			Secret: tokens[1],
		})
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid Token",
			})
			return
		} else if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		c.Set(MiddlewareKeyClientSession, ClientSession(*output))
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func MustGetClient(c *gin.Context) *ClientSession {
	r, e := GetClient(c)
	if !e {
		panic(e)
	}
	return r
}
func GetClient(c *gin.Context) (*ClientSession, bool) {
	r, exists := c.Get(MiddlewareKeyClientSession)
	if logger, ok := r.(ClientSession); !ok {
		return nil, false
	} else {
		return &logger, exists
	}
}
func GetLogger(c *gin.Context) (*log.Entry, bool) {
	r, exists := c.Get(MiddlewareKeyClientLogger)

	if logger, ok := r.(*log.Entry); !ok {
		return nil, false
	} else {
		return logger, exists
	}

}
