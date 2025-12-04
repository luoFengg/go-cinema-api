CREATE TABLE "users" (
  id VARCHAR(100) PRIMARY KEY,
  
  username VARCHAR(30) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL, -- Email tidak boleh kembar
  password VARCHAR(255) NOT NULL,     -- Akan berisi Hash panjang
  role VARCHAR(50) NOT NULL DEFAULT 'customer', -- Pilihan: 'admin' atau 'customer'
  
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())

  CONSTRAINT check_role CHECK (role IN ('admin', 'customer'))

);

-- Index untuk mempercepat login (pencarian by email)
CREATE INDEX "idx_users_email" ON "users" ("email");