package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/handler"
)

type RouterOpts struct {
	AuthHandler     handler.AuthHandler
	CategoryHandler handler.CategoryHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	public := r.Group("")
	public.POST("/login", opts.AuthHandler.HandleLogin)
	public.GET("/categories", opts.CategoryHandler.GetAllCategories)
	public.GET("/categories/:category-id", opts.CategoryHandler.GetCategoryDetail)

	return r
}
