package web

import "time"

type CheckOverlappingShowtimeCreateRequest struct {
	StudioID  string    `json:"studio_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}