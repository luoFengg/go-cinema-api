package services

import (
	"context"
	"fmt"
	"go-cinema-api/exceptions"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
	movieRepo "go-cinema-api/repositories/movie"
	showtimeRepo "go-cinema-api/repositories/showtime"
	"time"
)

type ShowtimeServiceImpl struct {
	showtimeRepo showtimeRepo.ShowtimeRepository
	movieRepo movieRepo.MovieRepository
}

func NewShowtimeService(showtimeRepo showtimeRepo.ShowtimeRepository, movieRepo movieRepo.MovieRepository) ShowtimeService {
	return &ShowtimeServiceImpl{
		showtimeRepo: showtimeRepo,
		movieRepo: movieRepo,
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