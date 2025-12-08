package services

import (
	"context"
	"go-cinema-api/exceptions"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
	repositories "go-cinema-api/repositories/booking"
	showtimeRepository "go-cinema-api/repositories/showtime"
	"go-cinema-api/utils"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type BookingServiceImpl struct {
	bookingRepo repositories.BookingRepository
	showtimeRepo showtimeRepository.ShowtimeRepository
	DB *gorm.DB
}

func NewBookingService(bookingRepo repositories.BookingRepository, showtimeRepo showtimeRepository.ShowtimeRepository, DB *gorm.DB) BookingService {
	return &BookingServiceImpl{
		bookingRepo: bookingRepo,
		showtimeRepo: showtimeRepo,
		DB: DB,
	}
}

func (service *BookingServiceImpl) CreateBooking(ctx context.Context, userID string, request web.BookingCreateRequest) (*domain.Booking, error) {
	var savedBooking *domain.Booking

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var errBooking error
		// 1. Validasi Input Awal
		if len(request.SeatIDs) == 0 {
			return exceptions.NewBadRequestError("seat_ids cannot be empty")
		}

		// ---------------------------------------------------------
        // TAHAP 1: LOCKING (Pessimistic Lock) - SOLUSI RACE CONDITION
        // Lock dulu kursi fisiknya di tabel master seats.
        // Jika ada orang lain yang mau booking kursi ini di detik yang sama,
        // database akan memaksa mereka menunggu di baris ini sampai kita selesai.
        // ---------------------------------------------------------
		errLock := service.bookingRepo.LockSeats(tx, request.SeatIDs)
		if errLock != nil {
			return errLock
		}

		// ---------------------------------------------------------
        // TAHAP 2: VALIDASI LOGIC (Check Availability)
        // Setelah kita kunci, baru kita cek: "Ada ngga sih bookingan aktif di kursi ini?"
        // ---------------------------------------------------------
		isBooked, err := service.bookingRepo.CheckSeatsAvailability(ctx, tx, request.ShowtimeID, request.SeatIDs)
		if err != nil {
			return err
		}
		
		if isBooked {
			return exceptions.NewConflictError("one or more seats are already booked")
		}

		// 3. Ambil Harga & Logic Booking
		showtime, errGetShowtime := service.showtimeRepo.GetShowtimeByID(ctx, request.ShowtimeID)
		if errGetShowtime != nil {
			return errGetShowtime
		}

		ticketPrice := showtime.Price
		totalPrice := ticketPrice * float64(len(request.SeatIDs))

		newBooking := domain.Booking{
			UserID:     userID,
			ShowtimeID: request.ShowtimeID,
			TotalPrice: totalPrice,
			Status:     "pending",
			BookingSeats: []domain.BookingSeat{},
		}

		// for loop untuk buat detail tiket
		for _, seatID := range request.SeatIDs {
			bookingSeat := domain.BookingSeat{
				SeatID: seatID,
				Price:  ticketPrice,
			}
			newBooking.BookingSeats = append(newBooking.BookingSeats, bookingSeat)
		}
		
		savedBooking, errBooking = service.bookingRepo.CreateBooking(ctx, tx, &newBooking)
		if errBooking != nil {
			return errBooking
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	err = service.DB.
	Preload("User").
	Preload("Showtime").
	Preload("Showtime.Studio").
	Preload("Showtime.Movie").
	Preload("BookingSeats").
	Preload("BookingSeats.Seat").
	First(savedBooking, "id = ?", savedBooking.ID).Error

	if err != nil {
		return nil, err
	}

	// ==========================================
    // REQUEST KE MIDTRANS (SNAP)
    // ==========================================

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  savedBooking.ID,
			GrossAmt: int64(savedBooking.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: savedBooking.User.FirstName,
			LName: savedBooking.User.LastName,
			Email: savedBooking.User.Email,
		},
		Expiry: &snap.ExpiryDetails{
			Unit:  "minutes",
			Duration: 5,
		},

	}

	// Panggil API Midtrans Snap untuk bikin transaksi using initialized client
	snapResp, errSnap := utils.SnapClient.CreateTransaction(snapReq)
	if errSnap != nil {
		return nil, errSnap
	}

	// Simpan payment_url & payment_token ke tabel bookings
	err = service.bookingRepo.UpdatePaymentInfo(ctx, nil, savedBooking.ID, snapResp.RedirectURL, snapResp.Token)
	if err != nil {
		return nil, err
	}

	// Update objek savedBooking biar ada payment_url & payment_tokennya
	savedBooking.PaymentURL = snapResp.RedirectURL
	savedBooking.PaymentToken = snapResp.Token

	return savedBooking, nil
}

func (service *BookingServiceImpl) GetBookingsByUserID(ctx context.Context, userID string) ([]domain.Booking, error) {
	bookings, err := service.bookingRepo.GetBookingsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}