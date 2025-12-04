package repositories

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repo *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	return repo.DB.WithContext(ctx).Create(user).Error
}

func (repo *UserRepositoryImpl) FindByEmailorUsername(ctx context.Context, identifier web.UserLoginRequest) (*domain.User, error) {
	var user domain.User
	
	err := repo.DB.WithContext(ctx).First(&user, "email = ? OR username = ?", identifier.Identifier, identifier.Identifier).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}