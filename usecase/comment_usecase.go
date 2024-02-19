package usecase

import (
	"context"

	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/repository"
)

type CommentUsecase interface {
	FindCommentsByNewsArticleId(ctx context.Context, payload *dto.CommentsRequest) (*dto.CommentsResponse, error)
	CreateComment(ctx context.Context, comment *entity.Comment) (uint, error)
}

func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{commentRepository}
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

func (uc *commentUsecase) FindCommentsByNewsArticleId(ctx context.Context, payload *dto.CommentsRequest) (*dto.CommentsResponse, error) {
	totalItem, err := uc.commentRepository.Count(ctx, payload)
	if err != nil {
		return nil, err
	}

	comments, err := uc.commentRepository.FindCommentsByNewsArticleId(ctx, payload)
	if err != nil {
		return nil, err
	}

	items := make([]*dto.CommentResponse, 0)
	for _, v := range comments {
		items = append(items, dto.CommentToCommentResponse(v))
	}

	totalPage := totalItem / payload.Limit
	if totalItem%payload.Limit != 0 {
		totalPage++
	}

	commentsResponse := dto.CommentsResponse{
		Items:     items,
		TotalItem: totalItem,
		TotalPage: totalPage,
	}

	return &commentsResponse, nil
}

func (uc *commentUsecase) CreateComment(ctx context.Context, comment *entity.Comment) (uint, error) {
	return uc.commentRepository.CreateComment(ctx, comment)
}
