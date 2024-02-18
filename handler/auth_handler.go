package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"github.com/sahaduta/coding-test-backend-httid/usecase"
)

type AuthHandler interface {
	HandleLogin(ctx *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

func (ah *authHandler) HandleLogin(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		err = apperror.ErrInvalidInput
		ctx.Error(err)
		return
	}

	accessToken, err := ah.authUsecase.Login(ctx, &entity.User{Username: loginRequest.Username, Password: loginRequest.Password})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: gin.H{"token": accessToken}})
}
