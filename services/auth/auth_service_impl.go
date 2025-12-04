package services

import (
	"context"
	"go-cinema-api/exceptions"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
	repositories "go-cinema-api/repositories/user"
	"go-cinema-api/utils"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo repositories.UserRepository
}

func NewAuthService(authRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		authRepo: authRepo,
	}
}

func (service *AuthServiceImpl) RegisterUser(ctx context.Context, request web.UserCreateRequest) (web.UserRegisterResponse, error) {
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserRegisterResponse{}, err
	}

	newUser := &domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Role:     "customer",
	}

	err = service.authRepo.CreateUser(ctx, newUser)
	if err != nil {
		errMsg := strings.ToLower(err.Error())
		if strings.Contains(errMsg, "duplicate") || strings.Contains(errMsg, "unique") {
			if strings.Contains(errMsg, "email") {
				return web.UserRegisterResponse{}, exceptions.NewConflictError("Email already in use")
			}
			if strings.Contains(errMsg, "username") {
				return web.UserRegisterResponse{}, exceptions.NewConflictError("Username already in use")
			}
			// Handle primary key duplicate (users_pkey)
			if strings.Contains(errMsg, "users_pkey") {
				return web.UserRegisterResponse{}, exceptions.NewConflictError("User registration failed - please try again")
			}
		}
		return web.UserRegisterResponse{}, err
	}
	
	return web.UserRegisterResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
		Role:	  newUser.Role,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}


func (service *AuthServiceImpl) LoginUser(ctx context.Context, identifier web.UserLoginRequest) (web.UserLoginResponse, error) {
	// 1. Find user by email or username
	user, err := service.authRepo.FindByEmailorUsername(ctx, identifier)
	if err != nil {
		return web.UserLoginResponse{}, exceptions.NewUnauthorizedError("Username or email not found")
	}

	// 2. Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(identifier.Password))
	if err != nil {
		return web.UserLoginResponse{}, exceptions.NewUnauthorizedError("Invalid password")
	}

	// 3. Generate JWT token
	token, errToken := utils.GenerateToken(user.ID, user.Role, os.Getenv("JWT_SECRET"))

	if errToken != nil {
		return web.UserLoginResponse{}, errToken
	}

	return web.UserLoginResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username: user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}, nil
}