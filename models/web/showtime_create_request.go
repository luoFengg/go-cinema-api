package web

import "time"

type ShowtimeCreateRequest struct {
	StudioID  string    `json:"studio_id" binding:"required"`
	MovieID   string    `json:"movie_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
}