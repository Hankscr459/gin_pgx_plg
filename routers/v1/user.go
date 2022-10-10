package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/read", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
