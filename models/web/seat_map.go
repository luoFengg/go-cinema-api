package web

type SeatMapResponse struct {
	ShowtimeID string           `json:"showtime_id"`
	StudioID   string           `json:"studio_id"`
	Seats      []SeatWithStatus `json:"seats"`
}

type SeatWithStatus struct {
	ID          string `json:"id"`
	Row         string `json:"row"`
	Number      int    `json:"number"`
	IsAvailable bool   `json:"is_available"`
	Status      string `json:"status"` // "available" | "booked" | "maintenance"
}