package usecase

import (
	"context"
	"errors"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/repository"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"gorm.io/gorm"
)

type NewsArticleUsecase interface {
	GetAllNewsArticles(ctx context.Context, payload *dto.NewsArticlesRequest) (*dto.NewsArticlesResponse, error)
	FindNewsArticleDetail(ctx context.Context, newsArticle *entity.NewsArticle) (*entity.NewsArticle, error)
	CreateNewsArticle(ctx context.Context, payload *entity.NewsArticle) (uint, error)
	UpdateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error
	DeleteNewsArticle(ctx context.Context, address *entity.NewsArticle) error
}

func NewNewsArticleUsecase(newsArticleRepository repository.NewsArticleRepository) NewsArticleUsecase {
	return &newsArticleUsecase{newsArticleRepository}
}

type newsArticleUsecase struct {
	newsArticleRepository repository.NewsArticleRepository
}

func (uc *newsArticleUsecase) GetAllNewsArticles(ctx context.Context, payload *dto.NewsArticlesRequest) (*dto.NewsArticlesResponse, error) {
	totalItem, err := uc.newsArticleRepository.Count(ctx, payload)
	if err != nil {
		return nil, err
	}

	newsArticles, err := uc.newsArticleRepository.FindAllNewsArticles(ctx, payload)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	items := make([]*dto.NewsArticleResponse, 0)
	for _, v := range newsArticles {
		items = append(items, dto.NewsArticleToNewsArticleResponse(v))
	}

	totalPage := totalItem / payload.Limit
	if totalItem%payload.Limit != 0 {
		totalPage++
	}

	newsArticlesResponse := dto.NewsArticlesResponse{
		Items:     items,
		TotalItem: totalItem,
		TotalPage: totalPage,
	}
	if err != nil {
		return nil, err
	}

	return &newsArticlesResponse, nil
}

func (uc *newsArticleUsecase) FindNewsArticleDetail(ctx context.Context, payload *entity.NewsArticle) (*entity.NewsArticle, error) {
	return uc.newsArticleRepository.FindNewsArticleDetail(ctx, *payload)
}

func (uc *newsArticleUsecase) CreateNewsArticle(ctx context.Context, payload *entity.NewsArticle) (uint, error) {
	return uc.newsArticleRepository.CreateNewsArticle(ctx, payload)
}

func (uc *newsArticleUsecase) UpdateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error {
	_, err := uc.newsArticleRepository.FindNewsArticleDetail(ctx, *newsArticle)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrNewsArticleIdNotFound
		}
		return err
	}

	return uc.newsArticleRepository.UpdateNewsArticle(ctx, newsArticle)
}

func (uc *newsArticleUsecase) DeleteNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error {
	_, err := uc.newsArticleRepository.FindNewsArticleDetail(ctx, *newsArticle)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrNewsArticleIdNotFound
		}
		return err
	}

	return uc.newsArticleRepository.DeleteNewsArticle(ctx, newsArticle)
}
