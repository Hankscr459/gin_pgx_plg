package router

import (
	"fmt"
	"net/http"
	"pgx_plg/plugins/configs"

	"github.com/gin-gonic/gin"
)

var db, _ = configs.ConnectDb()

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/read", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})

	userRoute.GET("/create", func(ctx *gin.Context) {
		sqlStatement := `
			INSERT INTO users (name, date_of_birth, description)
			VALUES ('Jonathan', '2000/12/30', 'description')`
		_, err := db.Exec(ctx, sqlStatement)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})
}
