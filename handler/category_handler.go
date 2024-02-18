package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"github.com/sahaduta/coding-test-backend-httid/shared/constant"
	"github.com/sahaduta/coding-test-backend-httid/usecase"
)

type CategoryHandler interface {
	GetAllCategories(ctx *gin.Context)
	GetCategoryDetail(ctx *gin.Context)
}

type categoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(uc usecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{categoryUsecase: uc}
}

func (h *categoryHandler) GetAllCategories(ctx *gin.Context) {
	param := dto.CategoriesRequest{}
	ctx.ShouldBindQuery(&param)

	sanitizeCategoriesParam(&param)

	categories, err := h.categoryUsecase.GetAllCategories(ctx.Request.Context(), &param)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: categories,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *categoryHandler) GetCategoryDetail(ctx *gin.Context) {
	categoryIdParam := ctx.Param("category-id")
	categoryIdParamInt, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		ctx.Error(apperror.ErrInvalidCategoryId)
		return
	}
	category := &entity.Category{
		Id: uint(categoryIdParamInt),
	}
	category, err = h.categoryUsecase.FindCategoryDetail(ctx.Request.Context(), category)
	if err != nil {
		ctx.Error(err)
		return
	}
	categoryResponse := dto.CategoryToCategoryResponse(category)
	resp := dto.Response{
		Data: categoryResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func sanitizeCategoriesParam(param *dto.CategoriesRequest) {
	if param.SortBy != "name" && param.SortBy != "id" {
		param.SortBy = "name"
	}
	if param.Sort != "asc" && param.Sort != "desc" {
		param.Sort = "asc"
	}
	if param.Limit == 0 {
		param.Limit = constant.DefaultLimit
	}
	if param.Page == 0 {
		param.Page = constant.DefaultPage
	}
}
