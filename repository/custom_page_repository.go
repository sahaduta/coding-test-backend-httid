package repository

import (
	"context"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"gorm.io/gorm"
)

type CustomPageRepository interface {
	FindAllCustomPages(ctx context.Context, payload *dto.CustomPagesRequest) ([]*entity.CustomPage, error)
	Count(ctx context.Context, payload *dto.CustomPagesRequest) (int, error)
	FindCustomPageDetail(ctx context.Context, customPage entity.CustomPage) (*entity.CustomPage, error)
	CreateCustomPage(ctx context.Context, customPage *entity.CustomPage) (uint, error)
	UpdateCustomPage(ctx context.Context, customPage *entity.CustomPage) error
	DeleteCustomPage(ctx context.Context, customPage *entity.CustomPage) error
}
type customPageRepository struct {
	db *gorm.DB
}

func NewCustomPageRepository(db *gorm.DB) CustomPageRepository {
	return &customPageRepository{
		db: db,
	}
}

func (r *customPageRepository) FindAllCustomPages(ctx context.Context, payload *dto.CustomPagesRequest) ([]*entity.CustomPage, error) {
	customPage := entity.CustomPage{}
	customPages := make([]*entity.CustomPage, 0)
	q := r.db.WithContext(ctx).Model(&customPage).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Order(payload.SortBy + " " + payload.Sort).
		Offset((payload.Page - 1) * payload.Limit)
	err := q.Find(&customPages).Error
	if err != nil {
		return nil, err
	}
	return customPages, nil
}

func (r *customPageRepository) Count(ctx context.Context, payload *dto.CustomPagesRequest) (int, error) {
	var total int64 = 0
	customPage := entity.CustomPage{}
	err := r.db.WithContext(ctx).Model(customPage).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Count(&total).
		Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *customPageRepository) FindCustomPageDetail(ctx context.Context, customPage entity.CustomPage) (*entity.CustomPage, error) {
	q := r.db.WithContext(ctx).Model(&customPage).
		Where("custom_pages.custom_url = ?", customPage.CustomUrl)
	err := q.First(&customPage).Error
	if err != nil {
		return nil, err
	}
	return &customPage, nil
}

func (r *customPageRepository) CreateCustomPage(ctx context.Context, customPage *entity.CustomPage) (uint, error) {
	result := r.db.WithContext(ctx).Model(&entity.CustomPage{}).Create(&customPage)
	if result.Error != nil {
		return 0, result.Error
	}
	return customPage.Id, nil
}

func (r *customPageRepository) UpdateCustomPage(ctx context.Context, customPage *entity.CustomPage) error {
	err := r.db.WithContext(ctx).Updates(customPage).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *customPageRepository) DeleteCustomPage(ctx context.Context, customPage *entity.CustomPage) error {
	err := r.db.WithContext(ctx).Delete(customPage).Error
	if err != nil {
		return err
	}
	return nil
}
