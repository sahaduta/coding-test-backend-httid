package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAllCategories(ctx context.Context, payload *dto.CategoriesRequest) ([]*entity.Category, error)
	Count(ctx context.Context, payload *dto.CategoriesRequest) (int, error)
	FindCategoryDetail(ctx context.Context, productClassification entity.Category) (*entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) (uint, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, category *entity.Category) error
}
type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) FindAllCategories(ctx context.Context, payload *dto.CategoriesRequest) ([]*entity.Category, error) {
	category := entity.Category{}
	categories := make([]*entity.Category, 0)
	q := r.db.WithContext(ctx).Model(&category).
		Where("name ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Order(payload.SortBy + " " + payload.Sort).
		Offset((payload.Page - 1) * payload.Limit)
	err := q.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Count(ctx context.Context, payload *dto.CategoriesRequest) (int, error) {
	var total int64 = 0
	category := entity.Category{}
	err := r.db.WithContext(ctx).Model(category).
		Where("name ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Count(&total).
		Offset((payload.Page - 1) * payload.Limit).
		Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *categoryRepository) FindCategoryDetail(ctx context.Context, category entity.Category) (*entity.Category, error) {
	q := r.db.WithContext(ctx).Model(&category).
		Where("categories.id = ?", category.Id)
	err := q.First(&category).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, apperror.ErrCategoryIdNotFound
		default:
			return nil, err
		}
	}
	return &category, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (uint, error) {
	result := r.db.WithContext(ctx).Model(&entity.Category{}).Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.Id, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	fmt.Println(category)
	err := r.db.WithContext(ctx).Updates(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, category *entity.Category) error {
	err := r.db.WithContext(ctx).Delete(category).Error
	if err != nil {
		return err
	}
	return nil
}
