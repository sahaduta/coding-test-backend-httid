package repository

import (
	"context"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindCommentsByNewsArticleId(ctx context.Context, payload *dto.CommentsRequest) ([]*entity.Comment, error)
	Count(ctx context.Context, payload *dto.CommentsRequest) (int, error)
	CreateComment(ctx context.Context, comment *entity.Comment) (uint, error)
}
type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Count(ctx context.Context, payload *dto.CommentsRequest) (int, error) {
	var total int64 = 0
	comment := entity.Comment{}
	err := r.db.WithContext(ctx).Model(comment).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Count(&total).
		Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *commentRepository) FindCommentsByNewsArticleId(ctx context.Context, payload *dto.CommentsRequest) ([]*entity.Comment, error) {
	comment := entity.Comment{}
	comments := make([]*entity.Comment, 0)
	q := r.db.WithContext(ctx).Model(&comment).
		Where("news_article_id = ?", payload.NewsArticleId).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Order(payload.SortBy + " " + payload.Sort).
		Offset((payload.Page - 1) * payload.Limit)
	err := q.Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) CreateComment(ctx context.Context, comment *entity.Comment) (uint, error) {
	result := r.db.WithContext(ctx).Model(&entity.Comment{}).Create(&comment)
	if result.Error != nil {
		return 0, result.Error
	}
	return comment.Id, nil
}
