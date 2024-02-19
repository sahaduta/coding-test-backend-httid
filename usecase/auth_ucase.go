package usecase

import (
	"context"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/pkg/hasher"
	"github.com/sahaduta/coding-test-backend-httid/pkg/token"
	"github.com/sahaduta/coding-test-backend-httid/repository"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
)

type AuthUsecase interface {
	Login(ctx context.Context, user *entity.User) (string, error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
	jwt      token.JWTManager
	hasher   hasher.Hasher
}

func NewAuthUsecase(authRepo repository.AuthRepository, jwt token.JWTManager, hasher hasher.Hasher) AuthUsecase {
	return &authUsecase{authRepo: authRepo, jwt: jwt, hasher: hasher}
}

func (uc *authUsecase) Login(ctx context.Context, user *entity.User) (string, error) {
	user, err := uc.authRepo.FindOneByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", apperror.ErrInvalidCred
	}

	err = uc.hasher.ComparePasswords([]byte(user.Password), []byte(user.Password))
	if err != nil {
		return "", apperror.ErrInvalidCred
	}
	tokenPayload := dto.TokenPayload{
		UserId: user.Id,
	}
	tokenString, err := uc.jwt.GenerateToken(&tokenPayload)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
