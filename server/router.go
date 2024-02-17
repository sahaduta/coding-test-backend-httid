package server

import (
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.ContextWithFallback = true

	return r
}
