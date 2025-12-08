ALTER TABLE bookings
DROP COLUMN IF EXISTS payment_token,
DROP COLUMN IF EXISTS payment_url;