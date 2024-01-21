package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gitdump "github.com/tpkeeper/gin-dump"
	"mygo.com/hari/golang-gin-poc/controller"
	"mygo.com/hari/golang-gin-poc/middlewares"
	"mygo.com/hari/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {

	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gitdump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome!!",
		})
	})
	server.Run(":8080")
}
