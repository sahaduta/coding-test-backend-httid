package repository

import (
	"context"
	"errors"

	"github.com/sahaduta/coding-test-backend-httid/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOneByUsername(ctx context.Context, username string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{}
	err := r.db.WithContext(ctx).
		Where("username = ?", username).First(&user).Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, nil
		default:
			return nil, err
		}
	}
	return &user, nil
}
