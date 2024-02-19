package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/handler"
	"github.com/sahaduta/coding-test-backend-httid/middleware"
)

type RouterOpts struct {
	AuthHandler     handler.AuthHandler
	CategoryHandler handler.CategoryHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	r.Use(middleware.HandleErrors())

	public := r.Group("")
	public.POST("/login", opts.AuthHandler.HandleLogin)
	public.GET("/categories", opts.CategoryHandler.GetAllCategories)
	public.GET("/categories/:category-id", opts.CategoryHandler.GetCategoryDetail)
	public.POST("/categories/", opts.CategoryHandler.CreateCategory)
	public.PUT("/categories/:category-id", opts.CategoryHandler.UpdateCategory)
	public.DELETE("/categories/:category-id", opts.CategoryHandler.DeleteCategory)

	return r
}
