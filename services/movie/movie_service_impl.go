package services

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
	repositories "go-cinema-api/repositories/movie"
)

type MovieServiceImpl struct {
	movieRepo repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) MovieService {
	return &MovieServiceImpl{
		movieRepo: movieRepo,
	}
}

func (service *MovieServiceImpl) CreateMovie(ctx context.Context, request web.MovieCreateRequest) (*domain.Movie, error) {
	newMovie := &domain.Movie{
		Title:       request.Title,
		Description: request.Description,
		DurationMin: request.DurationMin,
		ReleaseDate: request.ReleaseDate,
	}
	err := service.movieRepo.CreateMovie(ctx, newMovie)
	if err != nil {
		return nil, err
	}
	return newMovie, nil
}