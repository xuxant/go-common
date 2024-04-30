package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	errs "github.com/xuxant/go-common/errors"
)

func PanicHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Status(500)
				c.Writer.WriteHeaderNow()
				panic(err)
			}
		}()
		c.Next()
	}
}

func AddHealthCheckRoute(group *gin.RouterGroup) gin.IRoutes {
	return group.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
}

func AddCorsMiddleware(origin, credentials, headers, methods string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", credentials)
		c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
		c.Writer.Header().Set("Access-Control-Allow-Methods", methods)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SendError(c *gin.Context, err error) {
	var serviceError *errs.Error
	if !errors.As(err, &serviceError) {
		serviceError = errs.Internal.WithDetail(err)
	}
	log.Ctx(c.Request.Context()).Error().Msgf("error executing request: %v (%v)", serviceError.Message, serviceError.Detail)
	c.AbortWithStatusJSON(serviceError.HttpStatus, serviceError)
}
