# Quick Start Guide - Jobot

Быстрое руководство по запуску приложения Jobot.

## Предварительные требования

- Docker и Docker Compose
- PostgreSQL (если запускаете локально без Docker)
- Go 1.21+ (для разработки)
- Make (опционально, для использования Makefile команд)

## Быстрый старт с Docker

### 1. Клонируйте репозиторий
```bash
git clone <repository-url>
cd jobot
```

### 2. Создайте файл .env
```bash
cp env.example .env
# Отредактируйте .env при необходимости
```

### 3. Запустите базу данных
```bash
make docker-up
# или
docker compose -f deploy/debug/docker-compose.yaml up -d
```

### 4. Примените миграции
```bash
make db-migrate
# или
bash scripts/apply_migrations.sh
```

### 5. Запустите приложение
```bash
make run
# или
go run ./cmd/app
```

## Структура проекта

```
jobot/
├── cmd/app/              # Точка входа приложения
├── internal/             # Внутренний код
│   ├── api/             # API слой (контроллеры, модели)
│   ├── service/         # Бизнес-логика
│   ├── repository/      # Работа с БД
│   ├── application/     # Конфигурация приложения
│   └── transport/       # HTTP сервер
├── migrations/          # SQL миграции
├── scripts/            # Вспомогательные скрипты
└── deploy/             # Конфигурация развертывания
```

## База данных

### Применить миграции
```bash
make db-migrate
```

### Подключиться к БД через psql
```bash
make db-psql
```

### Сбросить БД (удалить все данные)
```bash
make db-reset
```

### Загрузить тестовые данные
```bash
docker compose -f deploy/debug/docker-compose.yaml exec postgres \
  psql -U postgres -d jobot -f /migrations/test_data.sql
```

## Структура базы данных

```
users → employees → resumes
                 └→ reactions → vacancies
     └→ employers ─────────────┘
```

### Таблицы:
- **users** - пользователи Telegram бота
- **employees** - профили соискателей
- **employers** - профили работодателей
- **resumes** - резюме (Telegram file ID)
- **vacancies** - вакансии
- **reactions** - реакции на вакансии

Подробнее см. [migrations/SCHEMA.md](migrations/SCHEMA.md)

## Разработка

### Установить зависимости
```bash
make deps
```

### Запустить тесты
```bash
make test
```

### Форматировать код
```bash
make fmt
```

### Проверить линтером
```bash
make lint
```

### Собрать приложение
```bash
make build
```

## Docker команды

### Запустить все сервисы
```bash
make docker-up
```

### Остановить сервисы
```bash
make docker-down
```

### Посмотреть логи
```bash
make docker-logs
# или для конкретного сервиса
make docker-logs-app
make docker-logs-db
```

### Пересобрать образы
```bash
make docker-rebuild
```

## Полезные команды

### Посмотреть все доступные команды
```bash
make help
```

### Проверить статус сервисов
```bash
make dev-status
```

### Подключиться к базе данных
```bash
# Через Docker
make db-psql

# Локально
psql -h localhost -p 5435 -U postgres -d jobot
```

### Посмотреть таблицы в БД
```sql
\dt                    -- список таблиц
\d users              -- структура таблицы users
\di                   -- список индексов
```

### Примеры SQL запросов
```sql
-- Все пользователи
SELECT * FROM users;

-- Все вакансии с информацией о компании
SELECT v.*, e.company_name 
FROM vacancies v
JOIN employers e ON v.employer_id = e.employer_id;

-- Вакансии с тегами golang
SELECT * FROM vacancies
WHERE 'golang' = ANY(tags);
```

## Переменные окружения

Основные переменные в файле `.env`:

```bash
# Database
DB_HOST=localhost
DB_PORT=5435
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=jobot

# Application
APP_NAME=jobot
APP_VERSION=1.0.0
APP_ENVIRONMENT=development
APP_DEBUG=true

# HTTP Server
HTTP_HOST=0.0.0.0
HTTP_PORT=8080

# Telegram Bot
TG_BOT_TOKEN=your_telegram_bot_token
```

## Troubleshooting

### База данных не запускается
```bash
# Проверить логи
make docker-logs-db

# Пересоздать контейнер
make docker-down-volumes
make docker-up
```

### Миграции не применяются
```bash
# Проверить подключение
make db-status

# Проверить существует ли БД
docker compose -f deploy/debug/docker-compose.yaml exec postgres \
  psql -U postgres -l

# Создать БД вручную
make db-create
```

### Порт уже занят
```bash
# Изменить порт в docker-compose.yaml
# или в .env файле
```

## Следующие шаги

1. Изучите [документацию миграций](migrations/README.md)
2. Ознакомьтесь со [схемой базы данных](migrations/SCHEMA.md)
3. Посмотрите [примеры API запросов](POSTMAN_TESTING.md)
4. Изучите структуру кода в `internal/`

## Дополнительная информация

- [README.md](README.md) - основная документация
- [migrations/README.md](migrations/README.md) - документация по миграциям
- [migrations/SCHEMA.md](migrations/SCHEMA.md) - схема базы данных
- [POSTMAN_TESTING.md](POSTMAN_TESTING.md) - тестирование API

## Поддержка

При возникновении проблем:
1. Проверьте логи: `make docker-logs`
2. Проверьте статус сервисов: `make dev-status`
3. Посмотрите документацию в `/migrations`

