-- Rollback: hapus constraint dan kolom
ALTER TABLE booking_seats DROP CONSTRAINT IF EXISTS uq_showtime_seat;
ALTER TABLE booking_seats DROP CONSTRAINT IF EXISTS fk_bookingseat_showtime;
ALTER TABLE booking_seats DROP COLUMN IF EXISTS showtime_id;
