package repositories

import (
	"context"
	"go-cinema-api/models/domain"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *domain.Movie) error
	GetMovieByID(ctx context.Context, id string) (*domain.Movie, error)
}