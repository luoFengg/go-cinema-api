ALTER TABLE bookings
ADD COLUMN payment_url TEXT,
ADD COLUMN payment_token VARCHAR(255);