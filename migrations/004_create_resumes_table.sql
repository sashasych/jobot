-- Create resumes table
-- Resumes are employee resume files stored as Telegram file IDs

CREATE TABLE IF NOT EXISTS resumes (
    resume_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL UNIQUE REFERENCES employees(employee_id) ON DELETE CASCADE,
    tg_file_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_resumes_employee_id ON resumes(employee_id);
CREATE INDEX idx_resumes_created_at ON resumes(created_at DESC);

-- Add comments
COMMENT ON TABLE resumes IS 'Employee resumes stored as Telegram file references';
COMMENT ON COLUMN resumes.resume_id IS 'Primary key - unique resume ID';
COMMENT ON COLUMN resumes.employee_id IS 'Foreign key to employees table';
COMMENT ON COLUMN resumes.tg_file_id IS 'Telegram file ID for resume document';
COMMENT ON COLUMN resumes.created_at IS 'Timestamp when resume was uploaded';
COMMENT ON COLUMN resumes.updated_at IS 'Timestamp when resume was last updated';

