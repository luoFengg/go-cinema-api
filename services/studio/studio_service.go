package services

import (
	"context"
	"go-cinema-api/models/domain"
)

type StudioService interface {
	CreateStudio(ctx context.Context,name string, capacity int) error
	GetStudioByID(ctx context.Context, studioID string) (*domain.Studio, error)
}