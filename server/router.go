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
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	r.Use(middleware.HandleErrors())

	public := r.Group("")
	public.POST("/login", opts.AuthHandler.HandleLogin)

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

	return r
}
