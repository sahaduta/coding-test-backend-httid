package usecase

import (
	"context"
	"errors"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/repository"
	"gorm.io/gorm"
)

type CategoryUsecase interface {
	GetAllCategories(ctx context.Context, payload *dto.CategoriesRequest) (*dto.CategoriesResponse, error)
	FindCategoryDetail(ctx context.Context, category *entity.Category) (*entity.Category, error)
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepository}
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func (uc *categoryUsecase) GetAllCategories(ctx context.Context, payload *dto.CategoriesRequest) (*dto.CategoriesResponse, error) {
	totalItem, err := uc.categoryRepository.Count(ctx, payload)
	if err != nil {
		return nil, err
	}

	categories, err := uc.categoryRepository.FindAllCategories(ctx, payload)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	items := make([]*dto.CategoryResponse, 0)
	for _, v := range categories {
		items = append(items, dto.CategoryToCategoryResponse(v))
	}

	totalPage := totalItem / payload.Limit
	if totalItem%payload.Limit != 0 {
		totalPage++
	}

	categoriesResponse := dto.CategoriesResponse{
		Items:     items,
		TotalItem: totalItem,
		TotalPage: totalPage,
	}
	if err != nil {
		return nil, err
	}

	return &categoriesResponse, nil
}

func (uc *categoryUsecase) FindCategoryDetail(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	return uc.categoryRepository.FindCategoryDetail(ctx, category)
}
