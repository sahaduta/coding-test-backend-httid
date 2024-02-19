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

type CommentHandler interface {
	GetCommentsByNewsArticleId(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
}

type commentHandler struct {
	commentUsecase usecase.CommentUsecase
}

func NewCommentHandler(uc usecase.CommentUsecase) CommentHandler {
	return &commentHandler{commentUsecase: uc}
}

func (h *commentHandler) GetCommentsByNewsArticleId(ctx *gin.Context) {
	newsArticleId := ctx.Param("news-article-id")
	newsArticleIdInt, err := strconv.Atoi(newsArticleId)
	if err != nil {
		ctx.Error(apperror.ErrInvalidNewsArticleId)
		return
	}
	param := dto.CommentsRequest{}
	ctx.ShouldBindQuery(&param)
	sanitizeCommentsParam(&param)
	param.NewsArticleId = uint(newsArticleIdInt)
	commentsResponse, err := h.commentUsecase.FindCommentsByNewsArticleId(ctx.Request.Context(), &param)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: commentsResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *commentHandler) CreateComment(ctx *gin.Context) {
	commentRequest := dto.CommentRequest{}
	err := ctx.ShouldBindJSON(&commentRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	comment := entity.Comment{
		NewsArticleId: commentRequest.NewsArticleId,
		Content:       commentRequest.Content,
		Name:          commentRequest.Name,
	}
	createdId, err := h.commentUsecase.CreateComment(ctx, &comment)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: gin.H{"id": createdId},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func sanitizeCommentsParam(param *dto.CommentsRequest) {
	if param.SortBy != "content" && param.SortBy != "id" {
		param.SortBy = "id"
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
