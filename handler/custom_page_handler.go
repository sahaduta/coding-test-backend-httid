package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/dto"
	"github.com/sahaduta/coding-test-backend-httid/entity"
	"github.com/sahaduta/coding-test-backend-httid/shared/apperror"
	"github.com/sahaduta/coding-test-backend-httid/shared/constant"
	"github.com/sahaduta/coding-test-backend-httid/usecase"
)

type CustomPageHandler interface {
	GetAllCustomPages(ctx *gin.Context)
	GetCustomPageDetail(ctx *gin.Context)
	CreateCustomPage(ctx *gin.Context)
	UpdateCustomPage(ctx *gin.Context)
	DeleteCustomPage(ctx *gin.Context)
}

type customPageHandler struct {
	customPageUsecase usecase.CustomPageUsecase
}

func NewCustomPageHandler(uc usecase.CustomPageUsecase) CustomPageHandler {
	return &customPageHandler{customPageUsecase: uc}
}

func (h *customPageHandler) GetAllCustomPages(ctx *gin.Context) {
	param := dto.CustomPagesRequest{}
	ctx.ShouldBindQuery(&param)

	sanitizeCustomPagesParam(&param)

	customPagesResponse, err := h.customPageUsecase.GetAllCustomPages(ctx.Request.Context(), &param)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: customPagesResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *customPageHandler) GetCustomPageDetail(ctx *gin.Context) {
	customUrlParam := ctx.Param("custom-url")
	customPage := &entity.CustomPage{
		CustomUrl: customUrlParam,
	}
	customPage, err := h.customPageUsecase.FindCustomPageDetail(ctx.Request.Context(), customPage)
	if err != nil {
		ctx.Error(err)
		return
	}
	customPageResponse := dto.CustomPageToCustomPageResponse(customPage)
	resp := dto.Response{
		Data: customPageResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *customPageHandler) CreateCustomPage(ctx *gin.Context) {
	customPageRequest := dto.CustomPageRequest{}
	err := ctx.ShouldBindJSON(&customPageRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	userId := ctx.Value(constant.UserId).(float64)
	customPage := entity.CustomPage{
		CustomUrl: customPageRequest.CustomUrl,
		Content:   customPageRequest.Content,
		UserId:    uint(userId),
	}
	createdId, err := h.customPageUsecase.CreateCustomPage(ctx, &customPage)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: gin.H{"id": createdId},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *customPageHandler) UpdateCustomPage(ctx *gin.Context) {
	customUrlParam := ctx.Param("custom-url")
	customPageRequest := dto.CustomPageRequest{}
	err := ctx.ShouldBindJSON(&customPageRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	userId := ctx.Value(constant.UserId).(float64)
	customPageRequest.UserId = uint(userId)
	customPageRequest.OldUrl = customUrlParam

	err = h.customPageUsecase.UpdateCustomPage(ctx, &customPageRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *customPageHandler) DeleteCustomPage(ctx *gin.Context) {
	customUrlParam := ctx.Param("custom-url")
	customPage := entity.CustomPage{CustomUrl: customUrlParam}

	err := h.customPageUsecase.DeleteCustomPage(ctx, &customPage)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusOK, resp)
}

func sanitizeCustomPagesParam(param *dto.CustomPagesRequest) {
	if param.SortBy != "content" && param.SortBy != "id" {
		param.SortBy = "content"
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
