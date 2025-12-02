package repositories

import (
	"context"
	"go-cinema-api/models/domain"

	"gorm.io/gorm"
)

type StudioRepositoryImpl struct {
	DB *gorm.DB
}

func NewStudioRepository(db *gorm.DB) StudioRepository {
	return &StudioRepositoryImpl{
		DB: db,
	}
}

func (repo *StudioRepositoryImpl) CreateStudioWithSeats(ctx context.Context, studio *domain.Studio) error {
	// GORM Transaction: Penting!
	// Simpan Studio dan Kursi. Jika salah satu gagal, semua harus batal.
	return repo.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Simpan Data Studio (Nama, Kapasitas)
		// Karena di struct Studio ada field 'Seats', GORM otomatis menyimpan seats-nya juga jika datanya ada.
		if err := tx.WithContext(ctx).Create(&studio).Error; err != nil {
			return err // Rollback otomatis jika error
		}
		return nil
	})

}

func (repo *StudioRepositoryImpl) GetStudioByID(ctx context.Context, studioID string) (*domain.Studio, error) {
	var studio domain.Studio

	err := repo.DB.WithContext(ctx).Preload("Seats").First(&studio, "id = ?", studioID).Error
	if err != nil {
		return nil, err
	}

	return &studio, nil
}