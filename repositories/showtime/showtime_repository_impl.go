package repositories

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"

	"gorm.io/gorm"
)

type ShowtimeRepositoryImpl struct {
	DB *gorm.DB
}

func NewShowtimeRepository(db * gorm.DB) ShowtimeRepository {
	return &ShowtimeRepositoryImpl{
		DB: db,
	}
}

// 1. Fungsi Simpan Jadwal
func (repo *ShowtimeRepositoryImpl) CreateShowtime(ctx context.Context, showtime *domain.Showtime) error {
	err := repo.DB.WithContext(ctx).Create(showtime).Error
	if err != nil {
		return err
	}
	return nil
}

// 2. LOGIC INTI: Cek Bentrok
// Rumus: (StartA < EndB) AND (EndA > StartB)
func (repo *ShowtimeRepositoryImpl) CheckOverlappingShowtime(ctx context.Context, request web.CheckOverlappingShowtimeCreateRequest) (bool, error) {
	var count int64

	err := repo.DB.WithContext(ctx).Model(&domain.Showtime{}).Where("studio_id = ?", request.StudioID).
	Where("start_time < ?", request.EndTime). // Jadwal yang ada MULAINYA sebelum jadwal baru SELESAI
	Where("end_time > ?", request.StartTime). // DAN Jadwal yang ada SELESAINYA setelah jadwal baru MULAI
	Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *ShowtimeRepositoryImpl) GetAllShowtimes(ctx context.Context) ([]domain.Showtime, error) {
	var showtimes []domain.Showtime
	err := repo.DB.WithContext(ctx).Preload("Movie").Preload("Studio").Find(&showtimes).Error
	if err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (repo *ShowtimeRepositoryImpl) GetShowtimeByID(ctx context.Context, showtimeID string) (domain.Showtime, error) {
	var showtime domain.Showtime

	err := repo.DB.WithContext(ctx).Preload("Movie").Preload("Studio").First(&showtime, "id = ?", showtimeID).Error
	if err != nil {
		return domain.Showtime{}, err
	}
	return showtime, nil
}