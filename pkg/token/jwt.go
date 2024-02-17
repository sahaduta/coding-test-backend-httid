package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/shared/constant"
)

type JWTManager interface {
	GenerateToken(payload *dto.TokenPayload) (string, error)
}

type JWTManagerImpl struct{}

func NewJWTHelper() JWTManager {
	return &JWTManagerImpl{}
}

func (h *JWTManagerImpl) GenerateToken(payload *dto.TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":     constant.AppIssuer,
			"iat":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"user_id": payload.UserId,
			"role":    payload.Role,
		})

	return token.SignedString([]byte(constant.AppSecretKey))
}
