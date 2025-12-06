-- 1. Tabel Header Transaksi
CREATE TABLE bookings (
    id VARCHAR(100) PRIMARY KEY,
    
    user_id VARCHAR(100) NOT NULL,
    showtime_id VARCHAR(100) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- Pilihan: 'pending', 'paid', 'cancelled'
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT check_status CHECK (status IN ('pending', 'paid', 'cancelled')),
    CONSTRAINT fk_booking_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_booking_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes(id)
);
-- 2. Tabel Detail Kursi (Pivot Table)
CREATE TABLE booking_seats (
    id VARCHAR(100) PRIMARY KEY,
    
    booking_id VARCHAR(100) NOT NULL,
    seat_id VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,

    -- Relasi
  CONSTRAINT fk_booking FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
  CONSTRAINT fk_seat FOREIGN KEY (seat_id) REFERENCES seats(id),
  
  -- Validasi Unik (PENTING)
  CONSTRAINT uq_booking_seat UNIQUE (booking_id, seat_id)
);

