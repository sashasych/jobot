-- Create employees table
-- Employees are job seekers who use the Telegram bot

CREATE TABLE IF NOT EXISTS employees (
    employee_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    tags TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_employees_user_id ON employees(user_id);
CREATE INDEX idx_employees_tags ON employees USING GIN(tags);
CREATE INDEX idx_employees_created_at ON employees(created_at DESC);

-- Add comments
COMMENT ON TABLE employees IS 'Job seekers who use the Telegram bot';
COMMENT ON COLUMN employees.employee_id IS 'Primary key - unique employee ID';
COMMENT ON COLUMN employees.user_id IS 'Foreign key to users table';
COMMENT ON COLUMN employees.tags IS 'Array of skills, interests, or job preferences';
COMMENT ON COLUMN employees.created_at IS 'Timestamp when employee record was created';
COMMENT ON COLUMN employees.updated_at IS 'Timestamp when employee record was last updated';

