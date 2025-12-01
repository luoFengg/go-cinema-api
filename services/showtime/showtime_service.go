package services

import (
	"context"
	"go-cinema-api/models/web"
)

type ShowtimeService interface {
	CreateShowtime(ctx context.Context, request web.ShowtimeCreateRequest) error
}