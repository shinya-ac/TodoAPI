package settings

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	config "github.com/shinya-ac/TodoAPI/configs"
	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *errDomain.Error:
				if errors.Is(err, errDomain.NotFoundErr) {
					ReturnNotFound(c, e)
				}
				ReturnStatusBadRequest(c, e)
			default:
				ReturnStatusInternalServerError(c, e)
			}
		}
	}
}

func ApiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeys := []string{config.Config.APIKey1, config.Config.APIKey2, config.Config.APIKey3}
		apiKey := c.GetHeader("Todo-API-Key")

		valid := false
		for _, key := range apiKeys {
			if apiKey == key {
				valid = true
				break
			}
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "APIkeyが有効ではありません。"})
			return
		}

		c.Next()
	}
}
