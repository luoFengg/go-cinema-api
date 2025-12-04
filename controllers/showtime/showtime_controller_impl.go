package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/showtime"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowtimeControllerImpl struct {
	showtimeService services.ShowtimeService
}

func NewShowtimeController(showtimeService services.ShowtimeService) ShowtimeController {
	return &ShowtimeControllerImpl{
		showtimeService: showtimeService,
	}
}

func (controller *ShowtimeControllerImpl) CreateShowtime(ctx *gin.Context) {
	var showtime web.ShowtimeCreateRequest

	err := ctx.ShouldBindJSON(&showtime)
	if  err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = controller.showtimeService.CreateShowtime(ctx.Request.Context(), showtime)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Success: true,
		Message: "Showtime created successfully",
		Data:    showtime,
	})
}

func (controller *ShowtimeControllerImpl) GetShowtimeList(ctx *gin.Context) {
	showtimes, err := controller.showtimeService.GetShowtimeList(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}	

	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Showtimes retrieved successfully",
		Data:    showtimes,
	})
}