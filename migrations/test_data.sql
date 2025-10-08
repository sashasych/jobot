-- Test data for jobot database
-- This file contains sample data for testing purposes

-- Insert test users
INSERT INTO users (id, tg_user_name, tg_chat_id, is_active, is_premium, role, created_at, updated_at) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'john_doe', '111111111', true, false, 'employee', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440002', 'jane_smith', '222222222', true, true, 'employee', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440003', 'tech_recruiter', '333333333', true, false, 'employer', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440004', 'startup_hr', '444444444', true, true, 'employer', NOW(), NOW())
ON CONFLICT (tg_chat_id) DO NOTHING;

-- Insert test employees
INSERT INTO employees (employee_id, user_id, tags, created_at, updated_at) VALUES
    ('660e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 
     ARRAY['golang', 'postgresql', 'docker', 'backend'], NOW(), NOW()),
    ('660e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440002', 
     ARRAY['react', 'typescript', 'frontend', 'ui/ux'], NOW(), NOW())
ON CONFLICT (user_id) DO NOTHING;

-- Insert test employers
INSERT INTO employers (employer_id, user_id, company_name, company_description, company_website, company_location, company_size, created_at, updated_at) VALUES
    ('770e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440003',
     'TechCorp Inc', 'Leading technology company specializing in AI and ML', 'https://techcorp.example.com',
     'Москва, Россия', '51-200', NOW(), NOW()),
    ('770e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440004',
     'StartupHub', 'Fast-growing startup in fintech space', 'https://startuphub.example.com',
     'Санкт-Петербург, Россия', '11-50', NOW(), NOW())
ON CONFLICT (user_id) DO NOTHING;

-- Insert test resumes
INSERT INTO resumes (resume_id, employee_id, tg_file_id, created_at, updated_at) VALUES
    ('880e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001',
     'BAADAgADZAAD1234567890', NOW(), NOW()),
    ('880e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440002',
     'BAADAgADaAAD0987654321', NOW(), NOW())
ON CONFLICT (employee_id) DO NOTHING;

-- Insert test vacancies
INSERT INTO vacancies (vacansie_id, employer_id, tags, title, description, location, salary, created_at, updated_at) VALUES
    ('990e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001',
     ARRAY['golang', 'kubernetes', 'microservices', 'senior'],
     'Senior Backend Developer (Go)',
     'We are looking for an experienced Backend Developer to join our team. Must have 5+ years of experience with Go, Kubernetes, and microservices architecture.',
     'Москва (можно удалённо)',
     '250,000 - 350,000 руб/месяц',
     NOW(), NOW()),
    ('990e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440001',
     ARRAY['python', 'machine learning', 'tensorflow', 'middle'],
     'Middle ML Engineer',
     'Join our AI team! We need a Machine Learning Engineer with experience in Python, TensorFlow, and deep learning.',
     'Москва',
     '180,000 - 250,000 руб/месяц',
     NOW(), NOW()),
    ('990e8400-e29b-41d4-a716-446655440003', '770e8400-e29b-41d4-a716-446655440002',
     ARRAY['react', 'typescript', 'nextjs', 'frontend'],
     'Frontend Developer (React)',
     'Looking for a talented Frontend Developer to build amazing user interfaces. Experience with React, TypeScript, and Next.js required.',
     'Санкт-Петербург (гибрид)',
     '150,000 - 200,000 руб/месяц',
     NOW(), NOW()),
    ('990e8400-e29b-41d4-a716-446655440004', '770e8400-e29b-41d4-a716-446655440002',
     ARRAY['fullstack', 'nodejs', 'react', 'junior'],
     'Junior Full Stack Developer',
     'Great opportunity for a junior developer to grow! We offer mentorship and exciting projects.',
     'Удалённо',
     '80,000 - 120,000 руб/месяц',
     NOW(), NOW());

-- Insert test reactions
INSERT INTO reactions (id, employee_id, vacancy_id, created_at) VALUES
    ('aa0e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001',
     '990e8400-e29b-41d4-a716-446655440001', NOW()),
    ('aa0e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440001',
     '990e8400-e29b-41d4-a716-446655440002', NOW()),
    ('aa0e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440002',
     '990e8400-e29b-41d4-a716-446655440003', NOW()),
    ('aa0e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002',
     '990e8400-e29b-41d4-a716-446655440004', NOW())
ON CONFLICT (employee_id, vacancy_id) DO NOTHING;

-- Verify data
SELECT 'Users:', COUNT(*) FROM users;
SELECT 'Employees:', COUNT(*) FROM employees;
SELECT 'Employers:', COUNT(*) FROM employers;
SELECT 'Resumes:', COUNT(*) FROM resumes;
SELECT 'Vacancies:', COUNT(*) FROM vacancies;
SELECT 'Reactions:', COUNT(*) FROM reactions;

