# Jobot - Job Board Application

Современное приложение для поиска работы и размещения вакансий, построенное на Go с использованием Clean Architecture.

## 🏗️ Архитектура

```
jobot/
├── cmd/app/                 # Точка входа приложения
├── internal/                # Внутренние пакеты
│   ├── api/                # API слой (HTTP handlers)
│   ├── application/        # Слой приложения
│   ├── service/            # Бизнес-логика
│   └── transport/          # Транспортный слой (REST)
├── pkg/                    # Переиспользуемые пакеты
│   └── logger/             # Логирование
├── deploy/debug/           # Docker Compose для разработки
└── migrations/             # Миграции базы данных
```

## 🚀 Быстрый старт

### Предварительные требования

- Go 1.25+
- Docker & Docker Compose
- Make (опционально)

### 1. Клонирование репозитория

```bash
git clone https://github.com/sashasych/jobot.git
cd jobot
```

### 2. Установка зависимостей

```bash
make deps
# или
go mod download
```

### 3. Запуск с Docker Compose

```bash
# Запуск всех сервисов (PostgreSQL, pgAdmin, Jobot App)
make docker-up

# Или вручную
docker compose -f deploy/debug/docker-compose.yaml up -d
```

### 4. Проверка работы

- **Приложение:** http://localhost:8080
- **pgAdmin:** http://localhost:5050
- **PostgreSQL:** localhost:5435

## 🛠️ Разработка

### Локальная разработка

```bash
# Запуск только базы данных
make dev-setup

# Запуск приложения локально
make run

# Или
go run ./cmd/app
```

### Полезные команды

```bash
# Просмотр логов
make docker-logs-app

# Пересборка и перезапуск
make docker-rebuild

# Остановка всех сервисов
make docker-down

# Очистка Docker ресурсов
make docker-clean
```

## 📊 Сервисы

### 1. Jobot App (Порт 8080)
- **Контейнер:** `jobot_app_container`
- **Изображение:** Собирается из Dockerfile
- **Health Check:** http://localhost:8080/health

### 2. PostgreSQL (Порт 5435)
- **Контейнер:** `jobotdb_container`
- **База данных:** `jobot`
- **Пользователь:** `asych`
- **Пароль:** `qwerty`

### 3. pgAdmin (Порт 5050)
- **Контейнер:** `pgadmin_container`
- **Email:** `asych@jobot.com`
- **Пароль:** `qwerty`

## 🔧 Конфигурация

### Переменные окружения

```yaml
# База данных
DB_HOST: postgres
DB_PORT: 5432
DB_NAME: jobot
DB_USER: asych
DB_PASSWORD: qwerty

# Приложение
APP_PORT: 8080
LOG_LEVEL: debug
```

### Настройка подключения к БД

В pgAdmin добавьте новый сервер:
- **Host:** `postgres`
- **Port:** `5432`
- **Database:** `jobot`
- **Username:** `asych`
- **Password:** `qwerty`

## 🧪 Тестирование

```bash
# Запуск тестов
make test

# Запуск с покрытием
go test -v -cover ./...

# Запуск конкретного теста
go test -v ./internal/service/user
```

## 📝 API Endpoints

### Health Check
```
GET /health
```

### Users
```
POST   /api/user
GET    /api/user
GET    /api/user/{UserID}
PUT    /api/user/{UserID}
DELETE /api/user/{UserID}
```

### Resumes
```
POST   /api/resumes
GET    /api/resumes/user/{EmployeeID}
GET    /api/resumes/{ResumeID}
PUT    /api/resumes/{ResumeID}
DELETE /api/resumes/{ResumeID}
```

## 🐳 Docker

### Сборка образа

```bash
make docker-build
```

### Запуск контейнера

```bash
docker run -p 8080:8080 jobot-app
```

### Просмотр логов

```bash
# Все сервисы
make docker-logs

# Только приложение
make docker-logs-app

# Только база данных
make docker-logs-db
```

## 📦 Зависимости

### Основные
- **Chi Router** - HTTP роутер
- **Zap** - Структурированное логирование
- **pgx** - PostgreSQL драйвер
- **envconfig** - Конфигурация из переменных окружения

### Разработка
- **Testify** - Тестирование
- **golangci-lint** - Линтинг

## 🔍 Отладка

### Проблемы с подключением к БД

```bash
# Проверка статуса сервисов
make dev-status

# Просмотр логов PostgreSQL
make docker-logs-db

# Перезапуск базы данных
docker compose -f deploy/debug/docker-compose.yaml restart postgres
```

### Проблемы с приложением

```bash
# Просмотр логов приложения
make docker-logs-app

# Пересборка приложения
make docker-rebuild

# Проверка health check
curl http://localhost:8080/health
```

## 📚 Документация

- [Архитектура контроллеров](internal/transport/rest/controllers/ARCHITECTURE.md)
- [Chi Router Setup](internal/transport/rest/CHI_ROUTER.md)
- [Миграции БД](migrations/README.md)

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции
3. Внесите изменения
4. Добавьте тесты
5. Создайте Pull Request

## 📄 Лицензия

Этот проект лицензирован под MIT License - см. файл [LICENSE](LICENSE) для деталей.

## 🆘 Поддержка

Если у вас есть вопросы или проблемы:

1. Проверьте [Issues](https://github.com/sashasych/jobot/issues)
2. Создайте новый Issue
3. Опишите проблему подробно

---

**Удачной разработки! 🚀**
