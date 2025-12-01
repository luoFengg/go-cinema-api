package controllers

import "github.com/gin-gonic/gin"

type ShowtimeController interface {
	CreateShowtime(ctx *gin.Context)
}