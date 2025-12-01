package controllers

import "github.com/gin-gonic/gin"

type StudioController interface {
	CreateStudio(ctx *gin.Context)
}