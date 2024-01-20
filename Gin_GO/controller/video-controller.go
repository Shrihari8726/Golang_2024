package controller

import (
	"github.com/gin-gonic/gin"
	"mygo.com/hari/golang-gin-poc/entity"
	"mygo.com/hari/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controllerVideo struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controllerVideo{service: service}
}

func (c *controllerVideo) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controllerVideo) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
