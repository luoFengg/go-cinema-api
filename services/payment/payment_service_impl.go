package services

import (
	"context"
	"go-cinema-api/models/web"
	repositories "go-cinema-api/repositories/booking"
	"go-cinema-api/utils"
	"log"

	"gorm.io/gorm"
)

type PaymentServiceImpl struct {
	bookingRepo repositories.BookingRepository
	DB *gorm.DB
}

func NewPaymentService(bookingRepo repositories.BookingRepository, db *gorm.DB) PaymentService {
	return &PaymentServiceImpl{
		bookingRepo: bookingRepo,
		DB: db,
	}
}

func (service *PaymentServiceImpl) ProcessPayment(ctx context.Context, input web.PaymentNotificationInput) error {
	// 1. VERIFIKASI (CHECK TRANSACTION STATUS) using shared CoreClient from utils
	checkResp, err := utils.CoreClient.CheckTransaction(input.OrderID)
	if err != nil {
		log.Printf("PaymentService: CheckTransaction error for order=%s: %v", input.OrderID, err)
		return err
	}

	if checkResp != nil {
		status := "pending"

		// 3. LOGIKA MAPPING (TERJEMAHAN BAHASA)
		// Midtrans punya banyak istilah status, Database saya cuma butuh: pending, paid, cancelled.


		switch checkResp.TransactionStatus {
		case "capture":
			// Case 1: Capture (Khusus Kartu Kredit)
			if checkResp.FraudStatus == "accept" {
				status = "paid"
			} else if checkResp.FraudStatus == "challenge" {
				status = "pending"
			}
		case "settlement":
			// Case 2: Settlement (Transfer Bank, E-Wallet, dll)
			// Ini artinya uang sudah dipastikan masuk (Settled).
			status = "paid"
		case "deny", "cancel", "expire":
		// Case 3: Gagal / Dibatalkan / Kadaluwarsa
			status = "cancelled"
		}

		// 4. UPDATE STATUS BOOKING DI DATABASE
		if status != "pending" {
			log.Printf("PaymentService: updating booking %s -> %s", input.OrderID, status)
			err := service.DB.Transaction(func(tx *gorm.DB) error {
				return service.bookingRepo.UpdateBookingStatus(ctx, tx, input.OrderID, status)
			})
			if err != nil {
				log.Printf("PaymentService: failed to update booking status for %s: %v", input.OrderID, err)
				return err
			}
		}
	}

	return nil
}