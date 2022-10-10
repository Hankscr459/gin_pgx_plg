package router

import (
	"net/http"
	"pgx_plg/plugins/configs"

	"github.com/gin-gonic/gin"
)

func RegisterConfigRoutes(rg *gin.RouterGroup) {
	configRoute := rg.Group("/config")
	configRoute.GET("/create-table", func(ctx *gin.Context) {
		err := configs.Migrate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
