package main

import (
	v1 "github/zine0/short_url/internal/api/v1"
	"github/zine0/short_url/internal/config"
	"github/zine0/short_url/internal/db"
	"github/zine0/short_url/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitConfig()
	db.InitDB()

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	v1.RegisterService(api)

	router.GET("/:key", service.Get)

	router.Run(":8080")
}
