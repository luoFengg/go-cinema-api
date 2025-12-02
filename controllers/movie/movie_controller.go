package controllers

import "github.com/gin-gonic/gin"

type MovieController interface {
	CreateMovie(ctx *gin.Context)
	GetAllMovies(ctx *gin.Context)
}