package main

import (
	"log"
	"net/http"
	"server/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}

	server.Use(cors.New(corsConfig))

	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	router := server.Group("api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok"})
	})

	log.Fatal(server.Run(":" + config.Port))
}
