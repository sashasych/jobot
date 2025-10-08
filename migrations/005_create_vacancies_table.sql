-- Create vacancies table
-- Vacancies are job postings created by employers

CREATE TABLE IF NOT EXISTS vacancies (
    vacansie_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employer_id UUID NOT NULL REFERENCES employers(employer_id) ON DELETE CASCADE,
    tags TEXT[] NOT NULL DEFAULT '{}',
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(255) NOT NULL,
    salary VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_vacancies_employer_id ON vacancies(employer_id);
CREATE INDEX idx_vacancies_tags ON vacancies USING GIN(tags);
CREATE INDEX idx_vacancies_title ON vacancies(title);
CREATE INDEX idx_vacancies_location ON vacancies(location);
CREATE INDEX idx_vacancies_created_at ON vacancies(created_at DESC);

-- Add comments
COMMENT ON TABLE vacancies IS 'Job postings created by employers';
COMMENT ON COLUMN vacancies.vacansie_id IS 'Primary key - unique vacancy ID';
COMMENT ON COLUMN vacancies.employer_id IS 'Foreign key to employers table';
COMMENT ON COLUMN vacancies.tags IS 'Array of job tags, skills, or categories';
COMMENT ON COLUMN vacancies.title IS 'Job title';
COMMENT ON COLUMN vacancies.description IS 'Job description and requirements';
COMMENT ON COLUMN vacancies.location IS 'Job location';
COMMENT ON COLUMN vacancies.salary IS 'Salary information';
COMMENT ON COLUMN vacancies.created_at IS 'Timestamp when vacancy was posted';
COMMENT ON COLUMN vacancies.updated_at IS 'Timestamp when vacancy was last updated';

