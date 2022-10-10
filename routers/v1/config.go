package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterConfigRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/config")
	configRoute.GET("/installation", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
