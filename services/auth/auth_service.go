package services

import (
	"context"
	"go-cinema-api/models/web"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request web.UserCreateRequest) (web.UserRegisterResponse, error)
	LoginUser(ctx context.Context, identifier web.UserLoginRequest) ( web.UserLoginResponse, error)
}