# Database Migrations

Этот каталог содержит SQL миграции для базы данных приложения jobot - Telegram бота для поиска работы.

## Обзор структуры базы данных

```
users (базовая таблица пользователей)
├── employees (сотрудники/соискатели)
│   ├── resumes (резюме)
│   └── reactions (реакции на вакансии)
│
└── employers (работодатели/компании)
    └── vacancies (вакансии)
        └── reactions (реакции на вакансии)
```

## Список миграций

### 001_create_users_table.sql
Создает базовую таблицу пользователей для Telegram бота.

**Таблица:** `users`

**Поля:**
- `id` (UUID) - первичный ключ
- `tg_user_name` (VARCHAR) - имя пользователя в Telegram (опционально)
- `tg_chat_id` (VARCHAR) - уникальный ID чата Telegram
- `is_active` (BOOLEAN) - активен ли аккаунт
- `is_premium` (BOOLEAN) - статус премиум подписки
- `role` (VARCHAR) - роль пользователя ('employee' или 'employer')
- `created_at` (TIMESTAMP) - дата создания
- `updated_at` (TIMESTAMP) - дата обновления

**Индексы:**
- `idx_users_tg_chat_id` - для быстрого поиска по chat ID
- `idx_users_role` - для фильтрации по роли
- `idx_users_is_active` - для фильтрации активных пользователей
- `idx_users_created_at` - для сортировки по дате

### 002_create_employees_table.sql
Создает таблицу сотрудников (соискателей).

**Таблица:** `employees`

**Поля:**
- `employee_id` (UUID) - первичный ключ
- `user_id` (UUID) - внешний ключ на users (UNIQUE)
- `tags` (TEXT[]) - навыки, интересы, предпочтения
- `created_at` (TIMESTAMP) - дата создания
- `updated_at` (TIMESTAMP) - дата обновления

**Индексы:**
- `idx_employees_user_id` - для связи с пользователем
- `idx_employees_tags` (GIN) - для полнотекстового поиска по тегам
- `idx_employees_created_at` - для сортировки

### 003_create_employers_table.sql
Создает таблицу работодателей (компаний/рекрутеров).

**Таблица:** `employers`

**Поля:**
- `employer_id` (UUID) - первичный ключ
- `user_id` (UUID) - внешний ключ на users (UNIQUE)
- `company_name` (VARCHAR) - название компании
- `company_description` (TEXT) - описание компании
- `company_website` (VARCHAR) - сайт компании
- `company_location` (VARCHAR) - местоположение
- `company_size` (VARCHAR) - размер компании
- `created_at` (TIMESTAMP) - дата создания
- `updated_at` (TIMESTAMP) - дата обновления

**Индексы:**
- `idx_employers_user_id` - для связи с пользователем
- `idx_employers_company_name` - для поиска по названию
- `idx_employers_company_location` - для поиска по локации
- `idx_employers_created_at` - для сортировки

### 004_create_resumes_table.sql
Создает таблицу резюме.

**Таблица:** `resumes`

**Поля:**
- `resume_id` (UUID) - первичный ключ
- `employee_id` (UUID) - внешний ключ на employees (UNIQUE)
- `tg_file_id` (VARCHAR) - ID файла в Telegram
- `created_at` (TIMESTAMP) - дата загрузки
- `updated_at` (TIMESTAMP) - дата обновления

**Индексы:**
- `idx_resumes_employee_id` - для связи с сотрудником
- `idx_resumes_created_at` - для сортировки

### 005_create_vacancies_table.sql
Создает таблицу вакансий.

**Таблица:** `vacancies`

**Поля:**
- `vacansie_id` (UUID) - первичный ключ
- `employer_id` (UUID) - внешний ключ на employers
- `tags` (TEXT[]) - теги, навыки, категории
- `title` (VARCHAR) - название вакансии
- `description` (TEXT) - описание и требования
- `location` (VARCHAR) - местоположение работы
- `salary` (VARCHAR) - информация о зарплате
- `created_at` (TIMESTAMP) - дата публикации
- `updated_at` (TIMESTAMP) - дата обновления

**Индексы:**
- `idx_vacancies_employer_id` - для связи с работодателем
- `idx_vacancies_tags` (GIN) - для полнотекстового поиска
- `idx_vacancies_title` - для поиска по названию
- `idx_vacancies_location` - для поиска по локации
- `idx_vacancies_created_at` - для сортировки

### 006_create_reactions_table.sql
Создает таблицу реакций (лайки/дизлайки сотрудников на вакансии).

**Таблица:** `reactions`

**Поля:**
- `id` (UUID) - первичный ключ
- `employee_id` (UUID) - внешний ключ на employees
- `vacancy_id` (UUID) - внешний ключ на vacancies
- `created_at` (TIMESTAMP) - дата реакции

**Ограничения:**
- UNIQUE(employee_id, vacancy_id) - один сотрудник может поставить только одну реакцию на вакансию

**Индексы:**
- `idx_reactions_employee_id` - для поиска реакций сотрудника
- `idx_reactions_vacancy_id` - для поиска реакций на вакансию
- `idx_reactions_created_at` - для сортировки

## Применение миграций

### Вручную через psql

```bash
# Убедитесь, что база данных создана
createdb jobot

# Применить миграции по порядку
psql -U postgres -d jobot -f migrations/001_create_users_table.sql
psql -U postgres -d jobot -f migrations/002_create_employees_table.sql
psql -U postgres -d jobot -f migrations/003_create_employers_table.sql
psql -U postgres -d jobot -f migrations/004_create_resumes_table.sql
psql -U postgres -d jobot -f migrations/005_create_vacancies_table.sql
psql -U postgres -d jobot -f migrations/006_create_reactions_table.sql
```

### Применить все миграции одной командой

```bash
# Linux/Mac
for f in migrations/*.sql; do psql -U postgres -d jobot -f "$f"; done

# Или через цикл с явным порядком
for i in 001 002 003 004 005 006; do
    psql -U postgres -d jobot -f "migrations/${i}_*.sql"
done
```

### С использованием переменных окружения

```bash
export PGHOST=localhost
export PGPORT=5432
export PGUSER=postgres
export PGPASSWORD=your_password
export PGDATABASE=jobot

# Применить миграцию
psql -f migrations/001_create_users_table.sql
```

### Docker Compose

```bash
# Если используете docker-compose
docker-compose exec postgres psql -U postgres -d jobot -f /migrations/001_create_users_table.sql
```

## Важные замечания

1. **Порядок применения**: Миграции должны применяться строго по порядку (001, 002, 003, ...) из-за зависимостей внешних ключей.

2. **Foreign Keys с CASCADE**: Все внешние ключи используют `ON DELETE CASCADE`, что означает:
   - При удалении пользователя удаляются связанные employee/employer записи
   - При удалении employee удаляются его резюме и реакции
   - При удалении employer удаляются его вакансии
   - При удалении vacancy удаляются связанные реакции

3. **UUID по умолчанию**: Используется `gen_random_uuid()` для автоматической генерации UUID (требует расширение pgcrypto в старых версиях PostgreSQL).

4. **GIN индексы**: Используются для полнотекстового поиска по массивам (tags).

5. **UNIQUE ограничения**:
   - `users.tg_chat_id` - один чат = один пользователь
   - `employees.user_id` - один пользователь = один employee профиль
   - `employers.user_id` - один пользователь = один employer профиль
   - `resumes.employee_id` - один сотрудник = одно резюме
   - `reactions(employee_id, vacancy_id)` - одна реакция на вакансию

## Проверка применения миграций

```sql
-- Посмотреть все таблицы
\dt

-- Посмотреть структуру конкретной таблицы
\d users
\d employees
\d employers
\d resumes
\d vacancies
\d reactions

-- Посмотреть все индексы
\di

-- Проверить внешние ключи
SELECT
    tc.table_name, 
    kcu.column_name,
    ccu.table_name AS foreign_table_name,
    ccu.column_name AS foreign_column_name 
FROM information_schema.table_constraints AS tc 
JOIN information_schema.key_column_usage AS kcu
    ON tc.constraint_name = kcu.constraint_name
    AND tc.table_schema = kcu.table_schema
JOIN information_schema.constraint_column_usage AS ccu
    ON ccu.constraint_name = tc.constraint_name
    AND ccu.table_schema = tc.table_schema
WHERE tc.constraint_type = 'FOREIGN KEY' 
    AND tc.table_schema='public';

-- Посмотреть размер таблиц
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size,
    pg_total_relation_size(schemaname||'.'||tablename) AS size_bytes
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY size_bytes DESC;
```

## Откат миграций

Для отката миграций можно использовать DROP TABLE в обратном порядке:

```sql
-- Откат в обратном порядке
DROP TABLE IF EXISTS reactions CASCADE;
DROP TABLE IF EXISTS vacancies CASCADE;
DROP TABLE IF EXISTS resumes CASCADE;
DROP TABLE IF EXISTS employers CASCADE;
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS users CASCADE;
```

## Тестовые данные

После применения миграций можно добавить тестовые данные:

```sql
-- Создать тестового пользователя-сотрудника
INSERT INTO users (id, tg_user_name, tg_chat_id, role) 
VALUES (gen_random_uuid(), 'test_employee', '123456789', 'employee');

-- Создать тестового пользователя-работодателя
INSERT INTO users (id, tg_user_name, tg_chat_id, role) 
VALUES (gen_random_uuid(), 'test_employer', '987654321', 'employer');
```

## Дополнительные команды PostgreSQL

```sql
-- Удалить все данные из таблицы (сохранить структуру)
TRUNCATE TABLE reactions CASCADE;
TRUNCATE TABLE vacancies CASCADE;
TRUNCATE TABLE resumes CASCADE;
TRUNCATE TABLE employees CASCADE;
TRUNCATE TABLE employers CASCADE;
TRUNCATE TABLE users CASCADE;

-- Сбросить счетчики последовательностей (если используются)
-- (В нашем случае не нужно, т.к. используем UUID)

-- Пересоздать индексы
REINDEX TABLE users;
REINDEX TABLE employees;
REINDEX TABLE employers;
REINDEX TABLE resumes;
REINDEX TABLE vacancies;
REINDEX TABLE reactions;

-- Обновить статистику для оптимизатора запросов
ANALYZE users;
ANALYZE employees;
ANALYZE employers;
ANALYZE resumes;
ANALYZE vacancies;
ANALYZE reactions;
```

## Резервное копирование

```bash
# Создать backup базы данных
pg_dump -U postgres jobot > backup_$(date +%Y%m%d_%H%M%S).sql

# Создать backup только схемы
pg_dump -U postgres --schema-only jobot > schema_backup.sql

# Создать backup только данных
pg_dump -U postgres --data-only jobot > data_backup.sql

# Восстановить из backup
psql -U postgres jobot < backup.sql
```

## Рекомендации

1. **Всегда делайте backup** перед применением миграций в production
2. **Тестируйте миграции** на dev/staging окружении
3. **Используйте транзакции** для атомарного применения миграций
4. **Мониторьте производительность** после применения новых индексов
5. **Документируйте изменения** в этом README

