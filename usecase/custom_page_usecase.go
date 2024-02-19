package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/repository"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"gorm.io/gorm"
)

type CustomPageUsecase interface {
	GetAllCustomPages(ctx context.Context, payload *dto.CustomPagesRequest) (*dto.CustomPagesResponse, error)
	FindCustomPageDetail(ctx context.Context, customPage *entity.CustomPage) (*entity.CustomPage, error)
	CreateCustomPage(ctx context.Context, customPage *entity.CustomPage) (uint, error)
	UpdateCustomPage(ctx context.Context, payload *dto.CustomPageRequest) error
	DeleteCustomPage(ctx context.Context, customPage *entity.CustomPage) error
}

func NewCustomPageUsecase(customPageRepository repository.CustomPageRepository) CustomPageUsecase {
	return &customPageUsecase{customPageRepository}
}

type customPageUsecase struct {
	customPageRepository repository.CustomPageRepository
}

func (uc *customPageUsecase) GetAllCustomPages(ctx context.Context, payload *dto.CustomPagesRequest) (*dto.CustomPagesResponse, error) {
	totalItem, err := uc.customPageRepository.Count(ctx, payload)
	if err != nil {
		return nil, err
	}

	customPages, err := uc.customPageRepository.FindAllCustomPages(ctx, payload)
	if err != nil {
		return nil, err
	}

	items := make([]*dto.CustomPageResponse, 0)
	for _, v := range customPages {
		items = append(items, dto.CustomPageToCustomPageResponse(v))
	}

	totalPage := totalItem / payload.Limit
	if totalItem%payload.Limit != 0 {
		totalPage++
	}

	customPagesResponse := dto.CustomPagesResponse{
		Items:     items,
		TotalItem: totalItem,
		TotalPage: totalPage,
	}

	return &customPagesResponse, nil
}

func (uc *customPageUsecase) FindCustomPageDetail(ctx context.Context, payload *entity.CustomPage) (*entity.CustomPage, error) {
	customPage, err := uc.customPageRepository.FindCustomPageDetail(ctx, *payload)

	if err != nil {
		fmt.Println(err, "bbbbb")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrCustomUrlNotFound
		}
		return nil, err
	}

	return customPage, nil
}

func (uc *customPageUsecase) CreateCustomPage(ctx context.Context, customPage *entity.CustomPage) (uint, error) {
	_, err := uc.customPageRepository.FindCustomPageDetail(ctx, entity.CustomPage{CustomUrl: customPage.CustomUrl})
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return uc.customPageRepository.CreateCustomPage(ctx, customPage)
}

func (uc *customPageUsecase) UpdateCustomPage(ctx context.Context, payload *dto.CustomPageRequest) error {
	customPage, err := uc.customPageRepository.FindCustomPageDetail(ctx, entity.CustomPage{CustomUrl: payload.OldUrl})
	if err != nil {
		fmt.Println(err, "bbbbb")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrCustomUrlNotFound
		}
		return err
	}
	customPage.Content = payload.Content
	customPage.CustomUrl = payload.CustomUrl
	customPage.UserId = payload.UserId

	return uc.customPageRepository.UpdateCustomPage(ctx, customPage)
}

func (uc *customPageUsecase) DeleteCustomPage(ctx context.Context, customPage *entity.CustomPage) error {
	customPage, err := uc.customPageRepository.FindCustomPageDetail(ctx, *customPage)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrCustomUrlNotFound
		}
		return err
	}

	return uc.customPageRepository.DeleteCustomPage(ctx, customPage)
}
