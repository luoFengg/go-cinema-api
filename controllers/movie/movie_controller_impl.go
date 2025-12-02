package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/movie"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieControllerImpl struct {
	movieService services.MovieService
}

func NewMovieController(movieService services.MovieService) MovieController {
	return &MovieControllerImpl{
		movieService: movieService,
	}
}

func (controller *MovieControllerImpl) CreateMovie(ctx *gin.Context) {
	var request web.MovieCreateRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	movie, err := controller.movieService.CreateMovie(ctx.Request.Context(), request)
	if err != nil {
		ctx.Error(err)
		return
	}
	
	ctx.JSON(200, web.WebResponse{
		Success: true,
		Message: "Movie created successfully",
		Data:    movie,
	})
}

func (controller *MovieControllerImpl) GetAllMovies(ctx *gin.Context) {
	movies, err := controller.movieService.GetMovies(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}
	
	ctx.JSON(200, web.WebResponse{
		Success: true,
		Message: "Movies retrieved successfully",
		Data:    movies,
	})
}