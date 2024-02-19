package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/shared/constant"
)

func HandleTimeout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx2, cancel := context.WithTimeout(ctx.Request.Context(),
			time.Duration(constant.GetTimeout())*time.Second)
		defer cancel()
		ctx.Request = ctx.Request.WithContext(ctx2)
		ctx.Next()
	}
}
