package controllers

import "github.com/gin-gonic/gin"

type AuthController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}