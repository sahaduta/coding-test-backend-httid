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

type NewsArticleRepository interface {
	FindAllNewsArticles(ctx context.Context, payload *dto.NewsArticlesRequest) ([]*entity.NewsArticle, error)
	Count(ctx context.Context, payload *dto.NewsArticlesRequest) (int, error)
	FindNewsArticleDetail(ctx context.Context, productClassification entity.NewsArticle) (*entity.NewsArticle, error)
	CreateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) (uint, error)
	UpdateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error
	DeleteNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error
}
type newsArticleRepository struct {
	db *gorm.DB
}

func NewNewsArticleRepository(db *gorm.DB) NewsArticleRepository {
	return &newsArticleRepository{
		db: db,
	}
}

func (r *newsArticleRepository) FindAllNewsArticles(ctx context.Context, payload *dto.NewsArticlesRequest) ([]*entity.NewsArticle, error) {
	newsArticle := entity.NewsArticle{}
	categories := make([]*entity.NewsArticle, 0)
	q := r.db.WithContext(ctx).Model(&newsArticle).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Order(payload.SortBy + " " + payload.Sort).
		Offset((payload.Page - 1) * payload.Limit)
	err := q.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *newsArticleRepository) Count(ctx context.Context, payload *dto.NewsArticlesRequest) (int, error) {
	var total int64 = 0
	newsArticle := entity.NewsArticle{}
	err := r.db.WithContext(ctx).Model(newsArticle).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Count(&total).
		Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *newsArticleRepository) FindNewsArticleDetail(ctx context.Context, newsArticle entity.NewsArticle) (*entity.NewsArticle, error) {
	q := r.db.WithContext(ctx).Model(&newsArticle).
		Where("news_articles.id = ?", newsArticle.Id)
	err := q.First(&newsArticle).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, apperror.ErrNewsArticleIdNotFound
		default:
			return nil, err
		}
	}
	return &newsArticle, nil
}

func (r *newsArticleRepository) CreateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) (uint, error) {
	result := r.db.WithContext(ctx).Model(&entity.NewsArticle{}).Create(&newsArticle)
	if result.Error != nil {
		return 0, result.Error
	}
	return newsArticle.Id, nil
}

func (r *newsArticleRepository) UpdateNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error {
	fmt.Println(newsArticle)
	err := r.db.WithContext(ctx).Updates(newsArticle).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *newsArticleRepository) DeleteNewsArticle(ctx context.Context, newsArticle *entity.NewsArticle) error {
	err := r.db.WithContext(ctx).Delete(newsArticle).Error
	if err != nil {
		return err
	}
	return nil
}
