package services

import (
	"context"
	"fmt"
	"go-cinema-api/exceptions"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
	bookingRepo "go-cinema-api/repositories/booking"
	movieRepo "go-cinema-api/repositories/movie"
	showtimeRepo "go-cinema-api/repositories/showtime"
	studioRepo "go-cinema-api/repositories/studio"
	"time"
)

type ShowtimeServiceImpl struct {
	showtimeRepo showtimeRepo.ShowtimeRepository
	movieRepo movieRepo.MovieRepository
	studioRepo studioRepo.StudioRepository
	bookingRepo bookingRepo.BookingRepository
}

func NewShowtimeService(showtimeRepo showtimeRepo.ShowtimeRepository, movieRepo movieRepo.MovieRepository, studioRepo studioRepo.StudioRepository, bookingRepo bookingRepo.BookingRepository) ShowtimeService {
	return &ShowtimeServiceImpl{
		showtimeRepo: showtimeRepo,
		movieRepo: movieRepo,
		studioRepo: studioRepo,
		bookingRepo: bookingRepo,
	}
}

func (s *ShowtimeServiceImpl) CreateShowtime(ctx context.Context, request web.ShowtimeCreateRequest) error {
	// 1. Ambil data Film ( butuh durasinya )
	movie, err := s.movieRepo.GetMovieByID(ctx, request.MovieID)
	if err != nil {
		return err
	}

	// 2. Hitung Waktu Selesia (EndTime)
	// Rumus: StartTime + Duration (menit)
	duration := time.Duration(movie.DurationMin) * time.Minute
	endTime := request.StartTime.Add(duration)



	// 3. Cek Bentrok (Validasi)
	showtime := web.CheckOverlappingShowtimeCreateRequest{
		StudioID:  request.StudioID,
		StartTime: request.StartTime,
		EndTime:   endTime,
	}

	isOverlap, errOverlap := s.showtimeRepo.CheckOverlappingShowtime(ctx, showtime)
	if errOverlap != nil {
		return fmt.Errorf("gagal mengecek jadwal: %v", errOverlap)
	}
	if  isOverlap {
		return exceptions.NewConflictError("jadwal bentrok dengan jadwal yang sudah ada")
	}

	// 4. Jika aman, bungkus dan simpan
	newShowtime := &domain.Showtime{
		StudioID:  request.StudioID,
		MovieID:   request.MovieID,
		StartTime: request.StartTime,
		EndTime:   endTime,
		Price:     request.Price,
	}
	err = s.showtimeRepo.CreateShowtime(ctx, newShowtime)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShowtimeServiceImpl) GetShowtimeList(ctx context.Context) ([]domain.Showtime, error) {
	showtimes, err := s.showtimeRepo.GetAllShowtimes(ctx)
	if err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (s *ShowtimeServiceImpl) GetShowtimeByID(ctx context.Context, showtimeID string) (domain.Showtime, error) {
	showtime, err := s.showtimeRepo.GetShowtimeByID(ctx, showtimeID)
	if err != nil {
		return domain.Showtime{}, err
	}
	return showtime, nil
}

func (s *ShowtimeServiceImpl) GetSeatMapForShowtime(ctx context.Context, showtimeID string) ([]web.SeatWithStatus, error) {
	// 1. Ambil shotime (dengan studio_id)
	showtime, err := s.showtimeRepo.GetShowtimeByID(ctx, showtimeID)
	if err != nil {
		return nil, err
	}

	// 2. Ambil studio + seats
	studio, err := s.studioRepo.GetStudioByID(ctx, showtime.StudioID)
	if err != nil {
		return nil, err
	}

	// 3. Ambil daftar seat_id yang sudah dibooking di showtime ini
	bookedSeatIDs, err := s.bookingRepo.GetBookedSeatIDsForShowtime(ctx, showtimeID)
	if err != nil {
		return nil, err
	}
	
	bookedSet := map[string]struct{}{}
	for _, seatID := range bookedSeatIDs {
		bookedSet[seatID] = struct{}{}
	}

	// 4. Gabungkan data studio + bookedSeatIDs jadi SeatWithStatus
	var result []web.SeatWithStatus
	for _, seat := range studio.Seats {
		status := "available"
		if !seat.IsAvailable {
			status = "maintenance"
		} else {
			if _, ok := bookedSet[seat.ID]; ok {
				status = "booked"
			}
		}
		result = append(result, web.SeatWithStatus{
			ID:          seat.ID,
			Row:         seat.Row,
			Number:      seat.Number,
			IsAvailable: seat.IsAvailable,
			Status:      status,
		})
	}
	return result, nil
}