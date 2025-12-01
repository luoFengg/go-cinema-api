package repositories

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
)

type ShowtimeRepository interface {
	CreateShowtime(ctx context.Context, showtime *domain.Showtime) error
	CheckOverlappingShowtime(ctx context.Context, request web.CheckOverlappingShowtimeCreateRequest) (bool, error)
}