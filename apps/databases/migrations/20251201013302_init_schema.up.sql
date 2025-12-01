-- 1. Table Studios
CREATE TABLE studios (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  capacity INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- 2. Table Movies
CREATE TABLE movies (
  id VARCHAR(100) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  duration_min INTEGER NOT NULL,
  release_date TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- 3. Table Seats
CREATE TABLE seats (
  id VARCHAR(100) PRIMARY KEY,
  studio_id VARCHAR(100) NOT NULL,
  row VARCHAR(5) NOT NULL,
  number INTEGER NOT NULL,
  is_available BOOLEAN NOT NULL DEFAULT true,

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  
  CONSTRAINT fk_seat_studio FOREIGN KEY (studio_id) REFERENCES studios(id) ON DELETE CASCADE,
  CONSTRAINT uk_seat_position UNIQUE (studio_id, row, number)
);

-- 4. Table Showtimes
CREATE TABLE showtimes (
  id VARCHAR(100) PRIMARY KEY,
  studio_id VARCHAR(100) NOT NULL,
  movie_id VARCHAR(100) NOT NULL,
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  
  CONSTRAINT fk_showtime_studio FOREIGN KEY (studio_id) REFERENCES studios(id) ON DELETE CASCADE,
  CONSTRAINT fk_showtime_movie FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE
);

-- Index untuk mempercepat pencarian jadwal
CREATE INDEX idx_showtime_studio ON showtimes(studio_id);
CREATE INDEX idx_showtime_movie ON showtimes(movie_id);


-- migrate -path apps/databases/migrations -database "postgresql://root:root@localhost:5433/cinema_db?sslmode=disable" up