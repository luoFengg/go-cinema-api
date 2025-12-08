package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/payment"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentControllerImpl struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		paymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) HandleWebhook(ctx *gin.Context) {
	var input web.PaymentNotificationInput

	// 1. Bind JSON
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 2. Panggil Service
	err = controller.paymentService.ProcessPayment(ctx.Request.Context(), input)
	if err != nil {
		log.Printf("Payment webhook processing error for order=%s: %v", input.OrderID, err)
		ctx.JSON(http.StatusInternalServerError, web.WebResponse{
			Success: false,
			Message: "failed to process payment notification",
			Data:    nil,
		})
		return
	}

	// 3. Response sukses (WAJIB 200 OK ke Midtrans)
	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Payment notification processed successfully",
		Data:    nil,
	})
}