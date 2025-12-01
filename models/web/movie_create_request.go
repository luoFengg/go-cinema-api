package web

import "time"

type MovieCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DurationMin int    `json:"duration_min" binding:"required"`
	ReleaseDate time.Time `json:"release_date"`
}