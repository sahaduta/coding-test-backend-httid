package server

import (
	"github.com/sahaduta/coding-test-backend-httid/handler"
	"github.com/sahaduta/coding-test-backend-httid/pkg/hasher"
	"github.com/sahaduta/coding-test-backend-httid/pkg/token"
	"github.com/sahaduta/coding-test-backend-httid/repository"
	"github.com/sahaduta/coding-test-backend-httid/usecase"
	"gorm.io/gorm"
)

func GetRouterOpts(db *gorm.DB) RouterOpts {
	bcryptHasher := hasher.NewBcryptHasher()
	jwt := token.NewJWTHelper()

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, jwt, bcryptHasher)
	authHandler := handler.NewAuthHandler(authUsecase)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	opts := RouterOpts{
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
	}

	return opts
}
