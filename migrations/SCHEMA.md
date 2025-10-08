# Database Schema - Jobot

Эта документация описывает схему базы данных для приложения Jobot - Telegram бота для поиска работы.

## Диаграмма связей

```
┌─────────────────────────────────────┐
│            users                    │
├─────────────────────────────────────┤
│ id            UUID PK               │
│ tg_user_name  VARCHAR(255)          │
│ tg_chat_id    VARCHAR(255) UNIQUE   │
│ is_active     BOOLEAN               │
│ is_premium    BOOLEAN               │
│ role          VARCHAR(50)           │
│ created_at    TIMESTAMP             │
│ updated_at    TIMESTAMP             │
└──────────┬────────────┬─────────────┘
           │            │
           │            │
  ┌────────┘            └────────┐
  │                              │
  ▼                              ▼
┌─────────────────────────┐   ┌──────────────────────────────┐
│      employees          │   │         employers            │
├─────────────────────────┤   ├──────────────────────────────┤
│ employee_id  UUID PK    │   │ employer_id     UUID PK      │
│ user_id      UUID FK    │   │ user_id         UUID FK      │
│ tags         TEXT[]     │   │ company_name    VARCHAR(255) │
│ created_at   TIMESTAMP  │   │ company_desc... TEXT         │
│ updated_at   TIMESTAMP  │   │ company_website VARCHAR(255) │
└──────┬────────┬─────────┘   │ company_loc...  VARCHAR(255) │
       │        │              │ company_size    VARCHAR(100) │
       │        │              │ created_at      TIMESTAMP    │
       │        │              │ updated_at      TIMESTAMP    │
       │        │              └──────────┬───────────────────┘
       │        │                         │
       │        │                         │
       │        │                         ▼
       │        │              ┌──────────────────────────┐
       │        │              │       vacancies          │
       │        │              ├──────────────────────────┤
       │        │              │ vacansie_id  UUID PK     │
       │        │              │ employer_id  UUID FK     │
       │        │              │ tags         TEXT[]      │
       │        │              │ title        VARCHAR(255)│
       │        │              │ description  TEXT        │
       │        │              │ location     VARCHAR(255)│
       │        │              │ salary       VARCHAR(255)│
       │        │              │ created_at   TIMESTAMP   │
       │        │              │ updated_at   TIMESTAMP   │
       │        │              └──────────┬───────────────┘
       │        │                         │
       │        │                         │
       │        └──────────┐              │
       │                   │              │
       ▼                   ▼              │
┌─────────────────────┐ ┌────────────────┴──────────┐
│      resumes        │ │       reactions           │
├─────────────────────┤ ├───────────────────────────┤
│ resume_id   UUID PK │ │ id          UUID PK       │
│ employee_id UUID FK │ │ employee_id UUID FK       │
│ tg_file_id  VARCHAR │ │ vacancy_id  UUID FK       │
│ created_at  TIMEST. │ │ created_at  TIMESTAMP     │
│ updated_at  TIMEST. │ └───────────────────────────┘
└─────────────────────┘   UNIQUE (employee_id, vacancy_id)
  UNIQUE (employee_id)
```

## Таблицы и отношения

### 1. users (Пользователи)
**Описание**: Базовая таблица для всех пользователей Telegram бота

**Отношения**:
- One-to-One с `employees` (один пользователь → один профиль сотрудника)
- One-to-One с `employers` (один пользователь → один профиль работодателя)

**Ограничения**:
- `tg_chat_id` - уникальный
- `role` - проверка IN ('employee', 'employer')

### 2. employees (Сотрудники)
**Описание**: Профили соискателей работы

**Отношения**:
- Many-to-One с `users` (FK: user_id)
- One-to-One с `resumes`
- One-to-Many с `reactions`

**Ограничения**:
- `user_id` - уникальный (один пользователь = один профиль сотрудника)
- ON DELETE CASCADE - удаление пользователя удаляет сотрудника

### 3. employers (Работодатели)
**Описание**: Профили компаний и рекрутеров

**Отношения**:
- Many-to-One с `users` (FK: user_id)
- One-to-Many с `vacancies`

**Ограничения**:
- `user_id` - уникальный (один пользователь = один профиль работодателя)
- ON DELETE CASCADE - удаление пользователя удаляет работодателя

### 4. resumes (Резюме)
**Описание**: Резюме сотрудников, хранящиеся как файлы в Telegram

**Отношения**:
- Many-to-One с `employees` (FK: employee_id)

**Ограничения**:
- `employee_id` - уникальный (один сотрудник = одно резюме)
- ON DELETE CASCADE - удаление сотрудника удаляет резюме

**Особенности**:
- `tg_file_id` хранит Telegram file ID для доступа к файлу

### 5. vacancies (Вакансии)
**Описание**: Объявления о работе, созданные работодателями

**Отношения**:
- Many-to-One с `employers` (FK: employer_id)
- One-to-Many с `reactions`

**Ограничения**:
- ON DELETE CASCADE - удаление работодателя удаляет все вакансии

**Особенности**:
- `tags` - массив для хранения навыков/тегов (GIN индекс для поиска)

### 6. reactions (Реакции)
**Описание**: Взаимодействие сотрудников с вакансиями (лайки/дизлайки)

**Отношения**:
- Many-to-One с `employees` (FK: employee_id)
- Many-to-One с `vacancies` (FK: vacancy_id)

**Ограничения**:
- UNIQUE(employee_id, vacancy_id) - один сотрудник может поставить только одну реакцию на вакансию
- ON DELETE CASCADE - удаление сотрудника или вакансии удаляет реакции

## Индексы

### users
- `idx_users_tg_chat_id` - поиск по Telegram chat ID
- `idx_users_role` - фильтрация по роли
- `idx_users_is_active` - фильтрация активных пользователей
- `idx_users_created_at` - сортировка по дате создания

### employees
- `idx_employees_user_id` - связь с пользователем
- `idx_employees_tags` (GIN) - полнотекстовый поиск по тегам
- `idx_employees_created_at` - сортировка

### employers
- `idx_employers_user_id` - связь с пользователем
- `idx_employers_company_name` - поиск по названию компании
- `idx_employers_company_location` - поиск по локации
- `idx_employers_created_at` - сортировка

### resumes
- `idx_resumes_employee_id` - связь с сотрудником
- `idx_resumes_created_at` - сортировка

### vacancies
- `idx_vacancies_employer_id` - связь с работодателем
- `idx_vacancies_tags` (GIN) - полнотекстовый поиск по тегам
- `idx_vacancies_title` - поиск по названию
- `idx_vacancies_location` - поиск по локации
- `idx_vacancies_created_at` - сортировка

### reactions
- `idx_reactions_employee_id` - поиск реакций сотрудника
- `idx_reactions_vacancy_id` - поиск реакций на вакансию
- `idx_reactions_created_at` - сортировка

## Типы данных

### UUID
Все ID используют UUID для:
- Уникальности в распределенных системах
- Безопасности (невозможность угадать ID)
- Совместимости с microservices архитектурой

### TEXT[]
Массивы текста используются для:
- `tags` в employees и vacancies
- Гибкое хранение множественных значений
- Эффективный поиск через GIN индексы

### TIMESTAMP
Все временные метки используют:
- `NOW()` как значение по умолчанию
- Автоматическое отслеживание создания и обновления

## Cascade операции

Все внешние ключи используют `ON DELETE CASCADE`:

```
users (удаление)
  └─> employees (каскадное удаление)
        ├─> resumes (каскадное удаление)
        └─> reactions (каскадное удаление)
  └─> employers (каскадное удаление)
        └─> vacancies (каскадное удаление)
              └─> reactions (каскадное удаление)
```

Это означает:
- Удаление пользователя автоматически удалит все связанные данные
- Целостность данных всегда сохраняется
- Нет orphan records (потерянных записей без родителя)

## Примеры запросов

### Найти все вакансии компании
```sql
SELECT v.* 
FROM vacancies v
JOIN employers e ON v.employer_id = e.employer_id
WHERE e.company_name = 'TechCorp Inc';
```

### Получить резюме сотрудника по Telegram ID
```sql
SELECT r.* 
FROM resumes r
JOIN employees e ON r.employee_id = e.employee_id
JOIN users u ON e.user_id = u.id
WHERE u.tg_chat_id = '123456789';
```

### Найти вакансии по тегам
```sql
SELECT * FROM vacancies
WHERE tags && ARRAY['golang', 'backend']
ORDER BY created_at DESC;
```

### Получить все реакции сотрудника
```sql
SELECT v.title, v.salary, r.created_at
FROM reactions r
JOIN vacancies v ON r.vacancy_id = v.vacansie_id
JOIN employees e ON r.employee_id = e.employee_id
WHERE e.employee_id = '660e8400-e29b-41d4-a716-446655440001'
ORDER BY r.created_at DESC;
```

### Статистика по вакансиям работодателя
```sql
SELECT 
    e.company_name,
    COUNT(v.vacansie_id) as total_vacancies,
    COUNT(DISTINCT r.employee_id) as total_reactions
FROM employers e
LEFT JOIN vacancies v ON e.employer_id = v.employer_id
LEFT JOIN reactions r ON v.vacansie_id = r.vacancy_id
GROUP BY e.employer_id, e.company_name;
```

## Миграции

Миграции должны применяться в следующем порядке:

1. `001_create_users_table.sql` - базовая таблица
2. `002_create_employees_table.sql` - зависит от users
3. `003_create_employers_table.sql` - зависит от users
4. `004_create_resumes_table.sql` - зависит от employees
5. `005_create_vacancies_table.sql` - зависит от employers
6. `006_create_reactions_table.sql` - зависит от employees и vacancies

## Версионирование

Текущая версия схемы: **v1.0.0**

История изменений:
- v1.0.0 (2024) - Первоначальная схема базы данных

