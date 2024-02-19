package middleware

import (
	"errors"
	"net/http"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"

	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			resp := dto.Response{
				Message: err.Error(),
			}

			switch {
			// 401
			case errors.Is(err, apperror.ErrInvalidInput):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)

			// 404
			case errors.Is(err, apperror.ErrCategoryIdNotFound):
				ctx.AbortWithStatusJSON(http.StatusNotFound, resp)

			// 500
			default:
				resp.Message = apperror.ErrInternal.Error()
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			}
		}
	}
}
