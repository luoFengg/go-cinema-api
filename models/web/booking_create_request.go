package web

type BookingCreateRequest struct {
	ShowtimeID string   `json:"showtime_id" validate:"required"`             // Film apa dan jam berapa
	SeatIDs    []string `json:"seat_ids" validate:"required,min=1,required"` // Kursi yang dipesan apa aja? (Array ID)
}