package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/handler"
	"github.com/sahaduta/coding-test-backend-httid/middleware"
)

type RouterOpts struct {
	AuthHandler        handler.AuthHandler
	CategoryHandler    handler.CategoryHandler
	NewsArticleHandler handler.NewsArticleHandler
	CustomPageHandler  handler.CustomPageHandler
	CommentHandler     handler.CommentHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	r.Use(middleware.HandleErrors())
	r.Use(middleware.HandleLogging())
	r.Use(middleware.HandleTimeout())
	// r.Use(middleware.HandleXApiKey())

	public := r.Group("")
	public.POST("/login", opts.AuthHandler.HandleLogin)
	public.POST("/comments/", opts.CommentHandler.CreateComment)
	public.GET("/comments/:news-article-id", opts.CommentHandler.GetCommentsByNewsArticleId)

	private := r.Group("")
	private.Use(middleware.HandleAuth())
	private.GET("/categories", opts.CategoryHandler.GetAllCategories)
	private.GET("/categories/:category-id", opts.CategoryHandler.GetCategoryDetail)
	private.POST("/categories/", opts.CategoryHandler.CreateCategory)
	private.PUT("/categories/:category-id", opts.CategoryHandler.UpdateCategory)
	private.DELETE("/categories/:category-id", opts.CategoryHandler.DeleteCategory)

	private.GET("/news-article", opts.NewsArticleHandler.GetAllNewsArticles)
	private.GET("/news-article/:news-article-id", opts.NewsArticleHandler.GetNewsArticleDetail)
	private.POST("/news-article/", opts.NewsArticleHandler.CreateNewsArticle)
	private.PUT("/news-article/:news-article-id", opts.NewsArticleHandler.UpdateNewsArticle)
	private.DELETE("/news-article/:news-article-id", opts.NewsArticleHandler.DeleteNewsArticle)

	private.GET("/custom-page", opts.CustomPageHandler.GetAllCustomPages)
	private.GET("/custom-page/:custom-url", opts.CustomPageHandler.GetCustomPageDetail)
	private.POST("/custom-page/", opts.CustomPageHandler.CreateCustomPage)
	private.PUT("/custom-page/:custom-url", opts.CustomPageHandler.UpdateCustomPage)
	private.DELETE("/custom-page/:custom-url", opts.CustomPageHandler.DeleteCustomPage)

	return r
}
