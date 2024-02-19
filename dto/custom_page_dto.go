package dto

import (
	"github.com/sahaduta/coding-test-backend-httid/entity"
)

type CustomPageRequest struct {
	CustomUrl string `json:"custom_url" binding:"required"`
	Content   string `json:"content" binding:"required"`
	OldUrl    string
	UserId    uint
}

type CustomPagesRequest struct {
	Search string `form:"s"`
	Sort   string `form:"sort"`
	SortBy string `form:"sort-by"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type CustomPageResponse struct {
	Id        uint   `json:"id"`
	CustomUrl string `json:"custom_url"`
	Content   string `json:"content"`
}

type CustomPagesResponse struct {
	Items     []*CustomPageResponse `json:"items"`
	TotalPage int                   `json:"total_page"`
	TotalItem int                   `json:"total_item"`
}

func CustomPageToCustomPageResponse(customPage *entity.CustomPage) *CustomPageResponse {
	return &CustomPageResponse{
		Id:        customPage.Id,
		CustomUrl: customPage.CustomUrl,
		Content:   customPage.Content,
	}
}
