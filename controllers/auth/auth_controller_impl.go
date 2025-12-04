package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (c *AuthControllerImpl) RegisterUser(ctx *gin.Context) {
	var registerRequest web.UserCreateRequest

	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	registerResponse, errResponse := c.authService.RegisterUser(ctx.Request.Context(), registerRequest)
	
	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    registerResponse,
	})
}

func (c *AuthControllerImpl) LoginUser(ctx *gin.Context) {
	var loginRequest web.UserLoginRequest

	err := ctx.ShouldBindJSON(&loginRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	loginResponse, errResponse := c.authService.LoginUser(ctx.Request.Context(), loginRequest)

	if errResponse != nil {
		ctx.Error(errResponse)
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "User logged in successfully",
		Data:    loginResponse,
	})
}