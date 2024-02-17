package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/coding-test-backend-httid/handler"
)

type RouterOpts struct {
	AuthHandler handler.AuthHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	public := r.Group("")
	public.POST("/login", opts.AuthHandler.HandleLogin)

	return r
}
