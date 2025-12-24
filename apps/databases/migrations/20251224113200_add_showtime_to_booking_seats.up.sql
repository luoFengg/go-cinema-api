-- Tambahkan kolom showtime_id ke booking_seats untuk unique constraint
ALTER TABLE booking_seats 
ADD COLUMN showtime_id VARCHAR(100);

-- Update existing data (jika ada) dengan mengambil showtime_id dari bookings
UPDATE booking_seats bs
SET showtime_id = b.showtime_id
FROM bookings b
WHERE bs.booking_id = b.id;

-- Set NOT NULL setelah data di-update
ALTER TABLE booking_seats 
ALTER COLUMN showtime_id SET NOT NULL;

-- Tambahkan foreign key ke showtimes
ALTER TABLE booking_seats 
ADD CONSTRAINT fk_bookingseat_showtime 
FOREIGN KEY (showtime_id) REFERENCES showtimes(id);

-- CONSTRAINT UTAMA: Mencegah double booking di level database
-- Satu kursi hanya bisa di-booking sekali per showtime
ALTER TABLE booking_seats 
ADD CONSTRAINT uq_showtime_seat UNIQUE (showtime_id, seat_id);
