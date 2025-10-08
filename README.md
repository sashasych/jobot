# 🤖 Jobot - Telegram Job Search Bot

Современный Telegram бот для поиска работы и размещения вакансий, построенный на Go с использованием Clean Architecture.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?style=flat&logo=postgresql)](https://www.postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 📋 Описание

Jobot - это полнофункциональный Telegram бот для поиска работы, который позволяет:

- 👨‍💼 **Соискателям**: создавать профиль, загружать резюме, просматривать вакансии и ставить лайки
- 🏢 **Работодателям**: создавать профиль компании, публиковать вакансии, просматривать отклики
- 🔄 **Matching**: система лайков для связи соискателей и работодателей

## 🏗️ Архитектура

Проект построен на базе **Clean Architecture** с четким разделением слоев:

```
jobot/
├── cmd/app/                    # Точка входа приложения
├── internal/                   # Внутренние пакеты
│   ├── api/                   # API слой
│   │   ├── controllers/       # HTTP контроллеры
│   │   ├── converter/         # Конвертеры моделей
│   │   └── models/           # API модели (DTO)
│   ├── application/          # Конфигурация приложения
│   ├── service/              # Бизнес-логика
│   │   ├── user/            # Сервис пользователей
│   │   ├── employee/        # Сервис сотрудников
│   │   ├── employer/        # Сервис работодателей
│   │   ├── resume/          # Сервис резюме
│   │   ├── vacancy/         # Сервис вакансий
│   │   ├── reaction/        # Сервис реакций
│   │   └── models/          # Сервисные модели
│   ├── repository/          # Слой данных
│   │   ├── user/           # Репозиторий пользователей
│   │   ├── employee/       # Репозиторий сотрудников
│   │   ├── employer/       # Репозиторий работодателей
│   │   ├── resume/         # Репозиторий резюме
│   │   ├── vacancy/        # Репозиторий вакансий
│   │   └── reaction/       # Репозиторий реакций
│   └── transport/          # Транспортный слой
│       └── rest/          # REST API (Chi Router)
├── pkg/                    # Переиспользуемые пакеты
│   ├── database/          # Подключение к БД
│   └── logger/            # Логирование (Zap)
├── migrations/            # SQL миграции базы данных
│   ├── 001_create_users_table.sql
│   ├── 002_create_employees_table.sql
│   ├── 003_create_employers_table.sql
│   ├── 004_create_resumes_table.sql
│   ├── 005_create_vacancies_table.sql
│   ├── 006_create_reactions_table.sql
│   ├── README.md         # Документация миграций
│   ├── SCHEMA.md         # Схема базы данных
│   └── test_data.sql     # Тестовые данные
├── scripts/              # Вспомогательные скрипты
│   └── apply_migrations.sh
├── deploy/               # Конфигурация развертывания
│   └── debug/           # Docker Compose для разработки
├── Makefile             # Команды для разработки
├── QUICKSTART.md        # Быстрый старт
└── README.md            # Документация
```

## 🗄️ База данных

### Схема базы данных

```
┌─────────────────────┐
│       users         │  (Пользователи Telegram)
│  - id               │
│  - tg_user_name     │
│  - tg_chat_id       │
│  - role             │
└──────┬──────┬───────┘
       │      │
       ▼      ▼
┌──────────┐ ┌──────────┐
│employees │ │employers │
└────┬─────┘ └────┬─────┘
     │            │
     ▼            ▼
┌─────────┐  ┌──────────┐
│resumes  │  │vacancies │
└─────────┘  └────┬─────┘
                  │
     ┌────────────┘
     ▼
┌──────────┐
│reactions │ (Лайки на вакансии)
└──────────┘
```

Подробная схема: [migrations/SCHEMA.md](migrations/SCHEMA.md)

### Таблицы

- **users** - Пользователи Telegram бота (employee или employer)
- **employees** - Профили соискателей (связь с users)
- **employers** - Профили работодателей (связь с users)
- **resumes** - Резюме сотрудников (хранятся как Telegram file ID)
- **vacancies** - Вакансии работодателей
- **reactions** - Реакции (лайки) сотрудников на вакансии

## 🚀 Быстрый старт

### Предварительные требования

- **Go** 1.21+
- **Docker** & **Docker Compose**
- **PostgreSQL** 15+ (опционально, для локальной разработки)
- **Make** (опционально, для использования Makefile команд)

### 1. Клонирование репозитория

```bash
git clone https://github.com/sashasych/jobot.git
cd jobot
```

### 2. Настройка переменных окружения

```bash
cp env.example .env
# Отредактируйте .env файл при необходимости
```

### 3. Запуск с Docker Compose

```bash
# Запуск всех сервисов (PostgreSQL, pgAdmin)
make docker-up

# Или вручную
docker compose -f deploy/debug/docker-compose.yaml up -d
```

### 4. Применение миграций базы данных

```bash
# Автоматическое применение всех миграций
make db-migrate

# Или вручную
bash scripts/apply_migrations.sh
```

### 5. Запуск приложения

```bash
# Локально
make run

# Или
go run ./cmd/app
```

### 6. Проверка работы

- **API:** http://localhost:8080
- **Health Check:** http://localhost:8080/health
- **pgAdmin:** http://localhost:5050
- **PostgreSQL:** localhost:5435

Подробнее: [QUICKSTART.md](QUICKSTART.md)

## 🛠️ Разработка

### Установка зависимостей

```bash
make deps
# или
go mod download && go mod tidy
```

### Локальная разработка

```bash
# Запуск только базы данных
make dev-setup

# Запуск приложения локально
make run
```

### Команды Makefile

```bash
# Сборка и запуск
make build              # Собрать приложение
make run                # Запустить приложение
make test               # Запустить тесты
make clean              # Очистить артефакты
make deps               # Установить зависимости
make fmt                # Форматировать код
make lint               # Проверить линтером

# Docker команды
make docker-build       # Собрать Docker образ
make docker-up          # Запустить сервисы
make docker-down        # Остановить сервисы
make docker-logs        # Показать логи
make docker-logs-app    # Логи приложения
make docker-logs-db     # Логи базы данных
make docker-rebuild     # Пересобрать и перезапустить
make docker-clean       # Очистить Docker ресурсы

# База данных
make db-migrate         # Применить миграции
make db-migrate-reset   # Сбросить и применить миграции
make db-status          # Статус базы данных
make db-psql            # Подключиться к БД
make db-create          # Создать базу данных
make db-drop            # Удалить базу данных
make db-reset           # Полный сброс (drop, create, migrate)

# Разработка
make dev-setup          # Настроить окружение разработки
make dev-start          # Запустить окружение разработки
make dev-stop           # Остановить окружение
make dev-status         # Статус сервисов

# Помощь
make help               # Показать все команды
```

## 📝 API Endpoints

### Health Check
```http
GET /health
```
Проверка состояния API

---

### 👤 Users (Пользователи)

```http
POST   /api/user              # Создать пользователя
GET    /api/user              # Получить список пользователей
GET    /api/user/{UserID}     # Получить пользователя по ID
PUT    /api/user/{UserID}     # Обновить пользователя
DELETE /api/user/{UserID}     # Удалить пользователя
```

**Пример создания пользователя:**
```json
POST /api/user
{
  "tg_user_name": "john_doe",
  "tg_chat_id": "123456789",
  "is_active": true,
  "is_premium": false,
  "role": "employee"
}
```

---

### 👨‍💼 Employees (Сотрудники)

```http
POST   /api/employee                    # Создать профиль сотрудника
GET    /api/employee/{EmployeeID}       # Получить сотрудника по ID
PUT    /api/employee/{EmployeeID}       # Обновить сотрудника
DELETE /api/employee/{EmployeeID}       # Удалить сотрудника
```

**Пример создания сотрудника:**
```json
POST /api/employee
{
  "user_id": "550e8400-e29b-41d4-a716-446655440001",
  "tags": ["golang", "postgresql", "docker", "backend"]
}
```

---

### 🏢 Employers (Работодатели)

```http
POST   /api/employer                    # Создать профиль работодателя
GET    /api/employer/{EmployerID}       # Получить работодателя по ID
PUT    /api/employer/{EmployerID}       # Обновить работодателя
DELETE /api/employer/{EmployerID}       # Удалить работодателя
```

**Пример создания работодателя:**
```json
POST /api/employer
{
  "user_id": "550e8400-e29b-41d4-a716-446655440002",
  "company_name": "TechCorp Inc",
  "company_description": "Leading technology company",
  "company_website": "https://techcorp.com",
  "company_location": "Москва, Россия",
  "company_size": "51-200"
}
```

---

### 📄 Resumes (Резюме)

```http
POST   /api/resume                      # Загрузить резюме
GET    /api/resume/{ResumeID}           # Получить резюме по ID
PUT    /api/resume/{ResumeID}           # Обновить резюме
DELETE /api/resume/{ResumeID}           # Удалить резюме
```

**Пример загрузки резюме:**
```json
POST /api/resume
{
  "employee_id": "660e8400-e29b-41d4-a716-446655440001",
  "tg_file_id": "BAADAgADZAAD1234567890"
}
```

---

### 💼 Vacancies (Вакансии)

```http
POST   /api/vacancy                           # Создать вакансию
GET    /api/vacancy/{VacancyID}               # Получить вакансию по ID
GET    /api/vacancy                           # Получить список всех вакансий
GET    /api/vacancy/employer/{EmployerID}    # Получить вакансии работодателя
PUT    /api/vacancy/{VacancyID}               # Обновить вакансию
DELETE /api/vacancy/{VacancyID}               # Удалить вакансию
```

**Пример создания вакансии:**
```json
POST /api/vacancy
{
  "employer_id": "770e8400-e29b-41d4-a716-446655440001",
  "tags": ["golang", "kubernetes", "microservices"],
  "title": "Senior Backend Developer",
  "description": "We are looking for experienced Backend Developer...",
  "location": "Москва (можно удалённо)",
  "salary": "250,000 - 350,000 руб/месяц"
}
```

---

### 👍 Reactions (Реакции)

```http
POST   /api/reaction                            # Создать реакцию (лайк)
GET    /api/reaction/{ReactionID}               # Получить реакцию по ID
GET    /api/reaction/employee/{EmployeeID}     # Получить реакции сотрудника
DELETE /api/reaction/{ReactionID}               # Удалить реакцию
```

**Пример создания реакции:**
```json
POST /api/reaction
{
  "employee_id": "660e8400-e29b-41d4-a716-446655440001",
  "vacansie_id": "990e8400-e29b-41d4-a716-446655440001",
  "reaction": "like"
}
```

Подробное тестирование API: [POSTMAN_TESTING.md](POSTMAN_TESTING.md)

## 🗄️ База данных

### Применение миграций

```bash
# Применить все миграции
make db-migrate

# Сбросить и применить заново
make db-migrate-reset

# Подключиться к базе данных
make db-psql
```

### Загрузка тестовых данных

```bash
docker compose -f deploy/debug/docker-compose.yaml exec postgres \
  psql -U postgres -d jobot -f /migrations/test_data.sql
```

### Примеры SQL запросов

```sql
-- Получить все вакансии с информацией о компании
SELECT v.*, e.company_name 
FROM vacancies v
JOIN employers e ON v.employer_id = e.employer_id;

-- Найти вакансии по тегам
SELECT * FROM vacancies
WHERE tags && ARRAY['golang', 'backend']
ORDER BY created_at DESC;

-- Получить реакции сотрудника на вакансии
SELECT v.title, v.salary, r.created_at
FROM reactions r
JOIN vacancies v ON r.vacancy_id = v.vacansie_id
WHERE r.employee_id = '660e8400-e29b-41d4-a716-446655440001'
ORDER BY r.created_at DESC;
```

Подробнее: [migrations/README.md](migrations/README.md) и [migrations/SCHEMA.md](migrations/SCHEMA.md)

## 📊 Сервисы

### 1. Jobot App (Порт 8080)
- **Контейнер:** `jobot_app_container`
- **Образ:** Собирается из Dockerfile
- **Health Check:** http://localhost:8080/health
- **API:** REST API на Chi Router

### 2. PostgreSQL (Порт 5435 → 5432)
- **Контейнер:** `jobotdb_container`
- **База данных:** `jobot`
- **Пользователь:** `asych`
- **Пароль:** `qwerty`
- **Внешний порт:** 5435 (внутренний: 5432)

### 3. pgAdmin (Порт 5050)
- **Контейнер:** `pgadmin_container`
- **Email:** `asych@jobot.com`
- **Пароль:** `qwerty`
- **URL:** http://localhost:5050

## 🔧 Конфигурация

### Переменные окружения (.env)

```bash
# Database
DB_HOST=localhost
DB_PORT=5435
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=jobot
DB_SSLMODE=disable

# Application
APP_NAME=jobot
APP_VERSION=1.0.0
APP_ENVIRONMENT=development
APP_DEBUG=true

# HTTP Server
HTTP_HOST=0.0.0.0
HTTP_PORT=8080

# Telegram Bot (для будущей интеграции)
TG_BOT_TOKEN=your_telegram_bot_token
```

### Настройка подключения к БД в pgAdmin

1. Откройте http://localhost:5050
2. Войдите с учетными данными (asych@jobot.com / qwerty)
3. Добавьте новый сервер:
   - **Name:** Jobot DB
   - **Host:** `postgres` (или `localhost` если снаружи Docker)
   - **Port:** `5432` (внутри Docker) или `5435` (снаружи)
   - **Database:** `jobot`
   - **Username:** `asych`
   - **Password:** `qwerty`

## 🧪 Тестирование

### Запуск тестов

```bash
# Все тесты
make test

# С покрытием
go test -v -cover ./...

# Конкретный пакет
go test -v ./internal/service/user

# С race detector
go test -race ./...
```

### Линтинг

```bash
# Линтинг кода
make lint

# Или вручную
golangci-lint run
```

### Форматирование

```bash
# Форматировать код
make fmt

# Или вручную
go fmt ./...
```

## 📦 Зависимости

### Основные библиотеки

- **[Chi Router](https://github.com/go-chi/chi)** v5 - HTTP роутер и middleware
- **[Zap](https://github.com/uber-go/zap)** - Высокопроизводительное логирование
- **[pgx](https://github.com/jackc/pgx)** v5 - PostgreSQL драйвер
- **[envconfig](https://github.com/kelseyhightower/envconfig)** - Конфигурация из env
- **[uuid](https://github.com/google/uuid)** - Генерация UUID

### Разработка и тестирование

- **[Testify](https://github.com/stretchr/testify)** - Assertions для тестов
- **[golangci-lint](https://github.com/golangci/golangci-lint)** - Линтер

## 🐳 Docker

### Сборка образа

```bash
# С использованием Makefile
make docker-build

# Вручную
docker build -t jobot-app .
```

### Запуск контейнера

```bash
# С docker-compose (рекомендуется)
make docker-up

# Отдельно
docker run -p 8080:8080 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  --network jobot-network \
  jobot-app
```

### Просмотр логов

```bash
# Все сервисы
make docker-logs

# Только приложение
make docker-logs-app

# Только база данных
make docker-logs-db

# Следить за логами в реальном времени
docker compose -f deploy/debug/docker-compose.yaml logs -f
```

## 🔍 Отладка

### Проблемы с подключением к БД

```bash
# Проверить статус сервисов
make dev-status

# Проверить логи базы данных
make docker-logs-db

# Перезапустить PostgreSQL
docker compose -f deploy/debug/docker-compose.yaml restart postgres

# Подключиться к БД напрямую
make db-psql
```

### Проблемы с приложением

```bash
# Просмотр логов
make docker-logs-app

# Пересборка
make docker-rebuild

# Проверка health check
curl http://localhost:8080/health
```

### Проблемы с миграциями

```bash
# Проверить какие таблицы созданы
make db-psql
\dt

# Сбросить и применить миграции заново
make db-migrate-reset

# Применить миграции вручную
bash scripts/apply_migrations.sh
```

## 📚 Документация

### Основная документация
- **[README.md](README.md)** - Этот файл
- **[QUICKSTART.md](QUICKSTART.md)** - Быстрый старт
- **[LICENSE](LICENSE)** - Лицензия MIT

### База данных
- **[migrations/README.md](migrations/README.md)** - Документация по миграциям
- **[migrations/SCHEMA.md](migrations/SCHEMA.md)** - Схема базы данных
- **[migrations/test_data.sql](migrations/test_data.sql)** - Тестовые данные

### API и тестирование
- **[POSTMAN_TESTING.md](POSTMAN_TESTING.md)** - Тестирование API
- **[postman_collection.json](postman_collection.json)** - Postman коллекция

### Внутренняя документация
- **[internal/service/README.md](internal/service/README.md)** - Сервисный слой
- **[internal/repository/user/README.md](internal/repository/user/README.md)** - Репозитории

## 🎯 Roadmap

### Текущая версия (v1.0.0)
- ✅ Clean Architecture
- ✅ RESTful API на Chi Router
- ✅ CRUD операции для всех сущностей
- ✅ PostgreSQL с миграциями
- ✅ Docker и Docker Compose
- ✅ Логирование с Zap
- ✅ Health checks

### Планируется
- [ ] Интеграция с Telegram Bot API
- [ ] Аутентификация и авторизация (JWT)
- [ ] Поиск вакансий с фильтрами
- [ ] Система уведомлений
- [ ] Загрузка файлов (аватары, резюме)
- [ ] Статистика и аналитика
- [ ] WebSocket для real-time обновлений
- [ ] GraphQL API
- [ ] Интеграция с внешними job boards
- [ ] Mobile app (Flutter/React Native)

## 🤝 Вклад в проект

Мы приветствуем ваш вклад! Пожалуйста, следуйте этим шагам:

1. **Fork** репозитория
2. Создайте ветку для новой функции (`git checkout -b feature/amazing-feature`)
3. Внесите изменения
4. Добавьте тесты для новой функциональности
5. Убедитесь что все тесты проходят (`make test`)
6. Отформатируйте код (`make fmt`)
7. Проверьте линтером (`make lint`)
8. Commit изменений (`git commit -m 'Add amazing feature'`)
9. Push в ветку (`git push origin feature/amazing-feature`)
10. Создайте **Pull Request**

### Стандарты кода

- Следуйте [Effective Go](https://golang.org/doc/effective_go)
- Пишите тесты для новой функциональности
- Документируйте публичные функции и типы
- Используйте осмысленные имена переменных
- Избегайте `panic` в production коде

## 📄 Лицензия

Этот проект лицензирован под **MIT License** - см. файл [LICENSE](LICENSE) для деталей.

## 🆘 Поддержка

Если у вас есть вопросы или проблемы:

1. Проверьте существующие [Issues](https://github.com/sashasych/jobot/issues)
2. Прочитайте документацию в `/migrations` и `/internal`
3. Создайте новый Issue с подробным описанием проблемы
4. Включите логи и шаги для воспроизведения

### Часто задаваемые вопросы

**Q: Как применить миграции?**  
A: Используйте `make db-migrate` или см. [migrations/README.md](migrations/README.md)

**Q: Как загрузить тестовые данные?**  
A: См. раздел "Загрузка тестовых данных" выше

**Q: Как подключиться к базе данных?**  
A: Используйте `make db-psql` или подключитесь через pgAdmin

**Q: Порт 5435 уже занят**  
A: Измените порт в `deploy/debug/docker-compose.yaml`

## 👥 Авторы

- **Asych** - [GitHub](https://github.com/sashasych)

## 🙏 Благодарности

- [Go Community](https://golang.org/community) за отличный язык
- [Chi Router](https://github.com/go-chi/chi) за легкий и мощный роутер
- [Uber Zap](https://github.com/uber-go/zap) за быстрое логирование
- Всем контрибьюторам проекта

---

**Удачной разработки! 🚀**

Made with ❤️ using Go
