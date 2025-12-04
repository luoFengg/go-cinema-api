package repositories

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindByEmailorUsername(ctx context.Context, identifier web.UserLoginRequest) (*domain.User, error)
}