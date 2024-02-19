package dto

import (
	"github.com/sahaduta/coding-test-backend-httid/entity"
)

type CommentRequest struct {
	NewsArticleId uint   `json:"news_article_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Content       string `json:"content" binding:"required"`
}

type CommentsRequest struct {
	Search        string `form:"s"`
	Sort          string `form:"sort"`
	SortBy        string `form:"sort-by"`
	Limit         int    `form:"limit"`
	Page          int    `form:"page"`
	NewsArticleId uint
}

type CommentResponse struct {
	Id            uint   `json:"id"`
	NewsArticleId uint   `json:"news_article_id"`
	Name          string `json:"name"`
	Content       string `json:"content"`
}

type CommentsResponse struct {
	Items     []*CommentResponse `json:"items"`
	TotalPage int                `json:"total_page"`
	TotalItem int                `json:"total_item"`
}

func CommentToCommentResponse(comment *entity.Comment) *CommentResponse {
	return &CommentResponse{
		Id:            comment.Id,
		NewsArticleId: comment.NewsArticleId,
		Name:          comment.Name,
		Content:       comment.Content,
	}
}
