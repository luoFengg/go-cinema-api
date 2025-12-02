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

func (service *MovieServiceImpl) GetMovies(ctx context.Context) ([]domain.Movie, error) {
	movies, err := service.movieRepo.GetAllMovies(ctx)
	if err != nil {
		return nil, err
	}
	return movies, nil
}