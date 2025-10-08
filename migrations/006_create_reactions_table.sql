-- Create reactions table
-- Reactions represent employee interactions with vacancies (like/dislike)

CREATE TABLE IF NOT EXISTS reactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL REFERENCES employees(employee_id) ON DELETE CASCADE,
    vacancy_id UUID NOT NULL REFERENCES vacancies(vacansie_id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(employee_id, vacancy_id)
);

-- Create indexes
CREATE INDEX idx_reactions_employee_id ON reactions(employee_id);
CREATE INDEX idx_reactions_vacancy_id ON reactions(vacancy_id);
CREATE INDEX idx_reactions_created_at ON reactions(created_at DESC);

-- Add comments
COMMENT ON TABLE reactions IS 'Employee reactions (likes/dislikes) to job vacancies';
COMMENT ON COLUMN reactions.id IS 'Primary key - unique reaction ID';
COMMENT ON COLUMN reactions.employee_id IS 'Foreign key to employees table - employee who reacted';
COMMENT ON COLUMN reactions.vacancy_id IS 'Foreign key to vacancies table - vacancy that was reacted to';
COMMENT ON COLUMN reactions.created_at IS 'Timestamp when reaction was created';

