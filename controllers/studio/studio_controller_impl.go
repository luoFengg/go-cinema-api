package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/studio"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudioControllerImpl struct {
	studioService services.StudioService
}

func NewStudioController(studioService services.StudioService) StudioController {
	return &StudioControllerImpl{
		studioService: studioService,
	}
}

func (controller *StudioControllerImpl) CreateStudio(ctx *gin.Context) {
	var requset web.StudioCreateRequest

	err := ctx.ShouldBindJSON(&requset)
	if err != nil {
		ctx.JSON (http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = controller.studioService.CreateStudio(ctx.Request.Context(), requset.Name, requset.Capacity)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON (http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Studio created successfully",
		Data:    requset,
	})

}