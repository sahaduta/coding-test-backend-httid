package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
)

func HandleXApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		XApiKey := ctx.Request.Header.Get("X-API-KEY")
		if XApiKey != "loremipsum" {
			handleAuthError(ctx, apperror.ErrNotAuthorized)
			return
		}
		ctx.Next()
	}
}
