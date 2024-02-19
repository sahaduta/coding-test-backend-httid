package dto

import (
	"github.com/sahaduta/coding-test-backend-httid/entity"
)

type NewsArticleRequest struct {
	CategoryId uint   `json:"category_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type NewsArticlesRequest struct {
	Search string `form:"s"`
	Sort   string `form:"sort"`
	SortBy string `form:"sort-by"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type NewsArticleResponse struct {
	Id         uint   `json:"id"`
	CategoryId uint   `json:"category_id"`
	Content    string `json:"content"`
}

type NewsArticlesResponse struct {
	Items     []*NewsArticleResponse `json:"items"`
	TotalPage int                    `json:"total_page"`
	TotalItem int                    `json:"total_item"`
}

func NewsArticleToNewsArticleResponse(newsArticle *entity.NewsArticle) *NewsArticleResponse {
	return &NewsArticleResponse{
		Id:         newsArticle.Id,
		CategoryId: newsArticle.CategoryId,
		Content:    newsArticle.Content,
	}
}
