package services

import (
	"context"
	"go-cinema-api/models/domain"
	repositories "go-cinema-api/repositories/studio"
)

type StudioServiceImpl struct {
	repo repositories.StudioRepository
}

func NewStudioService(repo repositories.StudioRepository) StudioService {
	return &StudioServiceImpl{
		repo: repo,
	}
}

func (service *StudioServiceImpl) CreateStudio(ctx context.Context, name string, capacity int) error {
	// 1. Siapkan data studio

	newStudio := domain.Studio{
		Name:     name,
		Capacity: capacity,
	}

	// 2. ALGORITMA GENERATE KURSI
	// Aturan: 1 Baris maksimal 10 kursi
	seatsPerRow := 10
	var seats []domain.Seat

	// Hitung total baris yang dibutuhkan 
	// Misal: 25 kursi / 10 = 2 baris sisa 5. Jadi, totalnya nanti ada 3 baris.
	totalRows := capacity / seatsPerRow
	if capacity % seatsPerRow != 0 {
		totalRows++ // Tambah 1 baris untuk sisa kursi
	}

	currentRowChar := 'A' // Mulai dari Huruf A
	remainingSeats := capacity 

	for row := 0; row < totalRows; row++ {
		// Tentukan berapa kursi di baris ini (maksimal 10, atau sisa kursi)
		seatsInThisRow := seatsPerRow
		if remainingSeats < seatsPerRow {
			seatsInThisRow = remainingSeats
		}

		// Loop untuk membuat nomor kursi (1, 2, 3...)
		for seatNum := 1; seatNum <= seatsInThisRow; seatNum++ {
			newSeat := domain.Seat {
				Row:   string(currentRowChar),
				Number: seatNum,
			}
			seats = append(seats, newSeat)
		}

		// Kurangi sisa kursi yang harus dibuat
		remainingSeats -= seatsInThisRow
		// Pindah ke huruf selanjutnya (A -> B)
		currentRowChar++ 
	}

	// 3. Masukkan hasil generate ke struct Studio
	newStudio.Seats = seats

	// 4. Panggil Repository untuk simpan ke DB
	err := service.repo.CreateStudioWithSeats(ctx, &newStudio)
	if err != nil {
		return err
	}
	return nil
}