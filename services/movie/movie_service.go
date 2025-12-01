package services

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
)

type MovieService interface {
	CreateMovie(ctx context.Context, request web.MovieCreateRequest) (*domain.Movie, error)
}