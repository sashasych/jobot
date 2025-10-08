-- Create users table
-- This is the base table for all users in the Telegram bot

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tg_user_name VARCHAR(255),
    tg_chat_id VARCHAR(255) NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT true,
    is_premium BOOLEAN NOT NULL DEFAULT false,
    role VARCHAR(50) NOT NULL CHECK (role IN ('employee', 'employer')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes for faster lookups
CREATE INDEX idx_users_tg_chat_id ON users(tg_chat_id);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_created_at ON users(created_at DESC);

-- Add comments for documentation
COMMENT ON TABLE users IS 'Base users table for Telegram bot - contains both employees and employers';
COMMENT ON COLUMN users.id IS 'Primary key - unique user ID';
COMMENT ON COLUMN users.tg_user_name IS 'Telegram username (optional, may not be set)';
COMMENT ON COLUMN users.tg_chat_id IS 'Unique Telegram chat ID for user identification and messaging';
COMMENT ON COLUMN users.is_active IS 'Whether the user account is active';
COMMENT ON COLUMN users.is_premium IS 'Premium subscription status';
COMMENT ON COLUMN users.role IS 'User role: employee (job seeker) or employer (company/recruiter)';
COMMENT ON COLUMN users.created_at IS 'Timestamp when user record was created';
COMMENT ON COLUMN users.updated_at IS 'Timestamp when user record was last updated';

