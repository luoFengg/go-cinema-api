package services

import "context"

type StudioService interface {
	CreateStudio(ctx context.Context,name string, capacity int) error
}