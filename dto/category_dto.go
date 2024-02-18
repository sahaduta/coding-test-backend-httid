package dto

import (
	"github.com/sahaduta/coding-test-backend-httid/entity"
)

type CategoryRequest struct {
	Name string `json:"name" binding:"required,min=5"`
}

type CategoriesRequest struct {
	Search string `form:"s"`
	Sort   string `form:"sort"`
	SortBy string `form:"sort-by"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type CategoryResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoriesResponse struct {
	Items     []*CategoryResponse `json:"items"`
	TotalPage int                 `json:"total_page"`
	TotalItem int                 `json:"total_item"`
}

func CategoryToCategoryResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
