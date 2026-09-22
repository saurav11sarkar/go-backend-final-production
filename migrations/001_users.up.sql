CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS users (
 id UUID PRIMARY KEY,
 full_name VARCHAR(120) NOT NULL,
 email VARCHAR(255) UNIQUE NOT NULL,
 password TEXT NOT NULL,
 role VARCHAR(30) NOT NULL DEFAULT 'user',
 status VARCHAR(30) NOT NULL DEFAULT 'active',
 profile_picture TEXT,
 reset_code VARCHAR(20),
 reset_code_expires_at TIMESTAMPTZ,
 created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
 updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);
CREATE INDEX IF NOT EXISTS idx_users_role_status ON users(role,status);
