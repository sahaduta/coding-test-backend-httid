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
			// 400
			case errors.Is(err, apperror.ErrInvalidCategoryId):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.Is(err, apperror.ErrInvalidCustomUrl):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.Is(err, apperror.ErrInvalidNewsArticleId):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.Is(err, apperror.ErrInvalidInput):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)

			// 401
			case errors.Is(err, apperror.ErrInvalidCred):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			case errors.Is(err, apperror.ErrNotAuthorized):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			case errors.Is(err, apperror.ErrMissingMetadata):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)

			// 404
			case errors.Is(err, apperror.ErrCategoryIdNotFound):
				ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
			case errors.Is(err, apperror.ErrNewsArticleIdNotFound):
				ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
			case errors.Is(err, apperror.ErrCustomUrlNotFound):
				ctx.AbortWithStatusJSON(http.StatusNotFound, resp)

			// 500
			default:
				resp.Message = apperror.ErrInternal.Error()
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			}
		}
	}
}
