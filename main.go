package main

import (
	"context"
	"os"
	"pgx_plg/plugins/configs"
	router "pgx_plg/routers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn := configs.ConnectDb()
	defer conn.Close(context.Background())
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	basepath := server.Group("/v1")
	router.RegisterUserRoutes(basepath)
	router.RegisterConfigRoutes(basepath)
	server.Run(os.Getenv("PORT"))
}
