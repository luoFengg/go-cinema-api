package repositories

import (
	"context"
	"go-cinema-api/models/domain"

	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &MovieRepositoryImpl{
		DB: db,
	}
}

func (repo *MovieRepositoryImpl) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	err := repo.DB.WithContext(ctx).Create(movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *MovieRepositoryImpl) GetMovieByID(ctx context.Context, id string) (*domain.Movie, error) {
	var  movie domain.Movie
	err := repo.DB.WithContext(ctx).First(&movie, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}