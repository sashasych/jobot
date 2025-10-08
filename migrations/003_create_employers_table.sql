-- Create employers table
-- Employers are companies/recruiters who post job vacancies

CREATE TABLE IF NOT EXISTS employers (
    employer_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    company_name VARCHAR(255) NOT NULL,
    company_description TEXT NOT NULL,
    company_website VARCHAR(255),
    company_location VARCHAR(255) NOT NULL,
    company_size VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_employers_user_id ON employers(user_id);
CREATE INDEX idx_employers_company_name ON employers(company_name);
CREATE INDEX idx_employers_company_location ON employers(company_location);
CREATE INDEX idx_employers_created_at ON employers(created_at DESC);

-- Add comments
COMMENT ON TABLE employers IS 'Companies/recruiters who post job vacancies';
COMMENT ON COLUMN employers.employer_id IS 'Primary key - unique employer ID';
COMMENT ON COLUMN employers.user_id IS 'Foreign key to users table';
COMMENT ON COLUMN employers.company_name IS 'Name of the company';
COMMENT ON COLUMN employers.company_description IS 'Description of the company';
COMMENT ON COLUMN employers.company_website IS 'Company website URL';
COMMENT ON COLUMN employers.company_location IS 'Company location/address';
COMMENT ON COLUMN employers.company_size IS 'Company size category (e.g., 1-10, 11-50, 51-200, etc.)';
COMMENT ON COLUMN employers.created_at IS 'Timestamp when employer record was created';
COMMENT ON COLUMN employers.updated_at IS 'Timestamp when employer record was last updated';

