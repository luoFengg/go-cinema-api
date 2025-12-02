package repositories

import (
	"context"
	"go-cinema-api/models/domain"
)

type StudioRepository interface {
	CreateStudioWithSeats(ctx context.Context, studio *domain.Studio) error
	GetStudioByID(ctx context.Context, studioID string) (*domain.Studio, error)
}