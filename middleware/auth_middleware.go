package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"github.com/sahaduta/coding-test-backend-httid/shared/constant"
)

func HandleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := constant.AppSecretKey
		tokenString, err := extractToken(ctx)
		if err != nil {
			handleAuthError(ctx, err)
			return
		}

		token, err := parseToken(tokenString, secretKey)
		if err != nil {
			handleAuthError(ctx, apperror.ErrNotAuthorized)
			return
		}

		claims, ok := getTokenClaims(token)
		if !ok {
			handleAuthError(ctx, apperror.ErrNotAuthorized)
			return
		}

		userID, err := extractClaimsData(claims)
		if err != nil {
			handleAuthError(ctx, err)
			return
		}

		ctx.Set(constant.UserId, userID)
		ctx.Next()
	}
}

func extractToken(ctx *gin.Context) (string, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		return "", apperror.ErrMissingMetadata
	}
	return tokenString[len("Bearer "):], nil
}

func parseToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}

func getTokenClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func extractClaimsData(claims jwt.MapClaims) (float64, error) {
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, apperror.ErrNotAuthorized
	}
	if !ok {
		return 0, apperror.ErrNotAuthorized
	}
	return userID, nil
}

func handleAuthError(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
