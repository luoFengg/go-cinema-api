package services

import (
	"context"
	"go-cinema-api/models/web"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, input web.PaymentNotificationInput) error
}