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
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
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

func (h *categoryHandler) CreateCategory(ctx *gin.Context) {
	categoryRequest := dto.CategoryRequest{}
	err := ctx.ShouldBindJSON(&categoryRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	category := entity.Category{
		Name: categoryRequest.Name,
	}
	createdId, err := h.categoryUsecase.CreateCategory(ctx, &category)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: gin.H{constant.UserId: createdId},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *categoryHandler) UpdateCategory(ctx *gin.Context) {
	param := ctx.Param("category-id")
	categoryId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidCategoryId)
		return
	}

	categoryRequest := dto.CategoryRequest{}
	err = ctx.ShouldBindJSON(&categoryRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}

	category := entity.Category{
		Id:   uint(categoryId),
		Name: categoryRequest.Name,
	}

	err = h.categoryUsecase.UpdateCategory(ctx, &category)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *categoryHandler) DeleteCategory(ctx *gin.Context) {
	param := ctx.Param("category-id")
	categoryId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidCategoryId)
		return
	}

	category := entity.Category{Id: uint(categoryId)}

	err = h.categoryUsecase.DeleteCategory(ctx, &category)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusCreated, resp)
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
