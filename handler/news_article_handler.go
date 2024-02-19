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

type NewsArticleHandler interface {
	GetAllNewsArticles(ctx *gin.Context)
	GetNewsArticleDetail(ctx *gin.Context)
	CreateNewsArticle(ctx *gin.Context)
	UpdateNewsArticle(ctx *gin.Context)
	DeleteNewsArticle(ctx *gin.Context)
}

type newsArticleHandler struct {
	newsArticleUsecase usecase.NewsArticleUsecase
}

func NewNewsArticleHandler(uc usecase.NewsArticleUsecase) NewsArticleHandler {
	return &newsArticleHandler{newsArticleUsecase: uc}
}

func (h *newsArticleHandler) GetAllNewsArticles(ctx *gin.Context) {
	param := dto.NewsArticlesRequest{}
	ctx.ShouldBindQuery(&param)

	sanitizeNewsArticlesParam(&param)

	newsArticlesResponse, err := h.newsArticleUsecase.GetAllNewsArticles(ctx.Request.Context(), &param)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: newsArticlesResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *newsArticleHandler) GetNewsArticleDetail(ctx *gin.Context) {
	newsArticleIdParam := ctx.Param("news-article-id")
	newsArticleIdParamInt, err := strconv.Atoi(newsArticleIdParam)
	if err != nil {
		ctx.Error(apperror.ErrInvalidNewsArticleId)
		return
	}
	newsArticle := &entity.NewsArticle{
		Id: uint(newsArticleIdParamInt),
	}
	newsArticle, err = h.newsArticleUsecase.FindNewsArticleDetail(ctx.Request.Context(), newsArticle)
	if err != nil {
		ctx.Error(err)
		return
	}
	newsArticleResponse := dto.NewsArticleToNewsArticleResponse(newsArticle)
	resp := dto.Response{
		Data: newsArticleResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *newsArticleHandler) CreateNewsArticle(ctx *gin.Context) {
	newsArticleRequest := dto.NewsArticleRequest{}
	err := ctx.ShouldBindJSON(&newsArticleRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	userId := ctx.Value(constant.UserId).(float64)
	newsArticle := entity.NewsArticle{
		CategoryId: newsArticleRequest.CategoryId,
		Content:    newsArticleRequest.Content,
		UserId:     uint(userId),
	}
	createdId, err := h.newsArticleUsecase.CreateNewsArticle(ctx, &newsArticle)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: gin.H{"id": createdId},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *newsArticleHandler) UpdateNewsArticle(ctx *gin.Context) {
	param := ctx.Param("news-article-id")
	newsArticleId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidNewsArticleId)
		return
	}

	newsArticleRequest := dto.NewsArticleRequest{}
	err = ctx.ShouldBindJSON(&newsArticleRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	userId := ctx.Value(constant.UserId).(float64)
	newsArticle := entity.NewsArticle{
		Id:         uint(newsArticleId),
		CategoryId: newsArticleRequest.CategoryId,
		Content:    newsArticleRequest.Content,
		UserId:     uint(userId),
	}

	err = h.newsArticleUsecase.UpdateNewsArticle(ctx, &newsArticle)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *newsArticleHandler) DeleteNewsArticle(ctx *gin.Context) {
	param := ctx.Param("news-article-id")
	newsArticleId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidNewsArticleId)
		return
	}

	newsArticle := entity.NewsArticle{Id: uint(newsArticleId)}

	err = h.newsArticleUsecase.DeleteNewsArticle(ctx, &newsArticle)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusNoContent, resp)
}

func sanitizeNewsArticlesParam(param *dto.NewsArticlesRequest) {
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
