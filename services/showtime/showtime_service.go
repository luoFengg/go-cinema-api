package services

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
)

type ShowtimeService interface {
	CreateShowtime(ctx context.Context, request web.ShowtimeCreateRequest) error
	GetShowtimeList(ctx context.Context) ([]domain.Showtime, error)
	GetShowtimeByID(ctx context.Context, showtimeID string) (domain.Showtime, error)
	GetSeatMapForShowtime(ctx context.Context, showtimeID string) ([]web.SeatWithStatus, error)
}