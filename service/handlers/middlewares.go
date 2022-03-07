package handlers

import (
	"github.com/gin-gonic/gin"
)

func CustomizedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// customized middleware implement
	}
}
