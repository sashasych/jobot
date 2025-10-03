# Jobot Development Environment

Этот каталог содержит Docker Compose конфигурацию для разработки приложения Jobot.

## 🚀 Быстрый старт

### 1. Запуск всех сервисов

```bash
# Из корня проекта
make docker-up

# Или напрямую
docker compose -f deploy/debug/docker-compose.yaml up -d
```

### 2. Проверка статуса

```bash
make dev-status
# или
docker compose -f deploy/debug/docker-compose.yaml ps
```

### 3. Просмотр логов

```bash
# Все сервисы
make docker-logs

# Только приложение
make docker-logs-app

# Только база данных
make docker-logs-db
```

## 📊 Сервисы

| Сервис | Порт | Описание |
|--------|------|----------|
| **jobot-app** | 8080 | Основное приложение |
| **postgres** | 5435 | База данных PostgreSQL |
| **pgadmin** | 5050 | Веб-интерфейс для БД |

## 🔧 Конфигурация

### Переменные окружения для jobot-app

```yaml
DB_HOST: postgres          # Хост базы данных
DB_PORT: 5432             # Порт базы данных
DB_NAME: jobot            # Имя базы данных
DB_USER: asych            # Пользователь БД
DB_PASSWORD: qwerty       # Пароль БД
APP_PORT: 8080            # Порт приложения
LOG_LEVEL: debug          # Уровень логирования
```

### Подключение к базе данных

**PostgreSQL:**
- **Host:** `localhost`
- **Port:** `5435`
- **Database:** `jobot`
- **Username:** `asych`
- **Password:** `qwerty`

**pgAdmin:**
- **URL:** http://localhost:5050
- **Email:** `asych@jobot.com`
- **Password:** `qwerty`

## 🛠️ Команды разработки

### Основные команды

```bash
# Запуск всех сервисов
make docker-up

# Остановка всех сервисов
make docker-down

# Пересборка и перезапуск
make docker-rebuild

# Остановка с удалением volumes
make docker-down-volumes
```

### Отладка

```bash
# Просмотр логов приложения
make docker-logs-app

# Просмотр логов базы данных
make docker-logs-db

# Проверка health check
curl http://localhost:8080/health
```

### Очистка

```bash
# Очистка Docker ресурсов
make docker-clean

# Удаление volumes
docker compose -f deploy/debug/docker-compose.yaml down -v
```

## 🐳 Docker образы

### jobot-app
- **Сборка:** Из Dockerfile в корне проекта
- **Базовый образ:** `golang:1.25-alpine` (builder) + `alpine:latest` (runtime)
- **Размер:** ~15-20MB (оптимизированный)
- **Пользователь:** `appuser` (non-root)

### postgres
- **Образ:** `postgres:15-alpine`
- **Версия:** PostgreSQL 15
- **Размер:** ~200MB

### pgadmin
- **Образ:** `dpage/pgadmin4:7.2`
- **Версия:** pgAdmin 4.7.2
- **Размер:** ~500MB

## 🔍 Troubleshooting

### Проблема: Приложение не запускается

```bash
# Проверьте логи
make docker-logs-app

# Проверьте, что база данных готова
make docker-logs-db

# Перезапустите сервисы
make docker-rebuild
```

### Проблема: Не удается подключиться к БД

```bash
# Проверьте статус PostgreSQL
docker compose -f deploy/debug/docker-compose.yaml ps postgres

# Проверьте логи PostgreSQL
make docker-logs-db

# Перезапустите PostgreSQL
docker compose -f deploy/debug/docker-compose.yaml restart postgres
```

### Проблема: Порт уже используется

```bash
# Проверьте, что порты свободны
netstat -tulpn | grep -E ':(8080|5435|5050)'

# Остановите конфликтующие сервисы
sudo systemctl stop postgresql  # Если локальный PostgreSQL
```

### Проблема: Недостаточно памяти

```bash
# Проверьте использование ресурсов
docker stats

# Увеличьте лимиты в docker-compose.yaml
deploy:
  resources:
    limits:
      memory: 1G  # Увеличьте лимит
```

## 📝 Полезные команды

### Работа с базой данных

```bash
# Подключение к PostgreSQL
docker exec -it jobotdb_container psql -U asych -d jobot

# Создание бэкапа
docker exec jobotdb_container pg_dump -U asych jobot > backup.sql

# Восстановление из бэкапа
docker exec -i jobotdb_container psql -U asych -d jobot < backup.sql
```

### Работа с приложением

```bash
# Подключение к контейнеру приложения
docker exec -it jobot_app_container sh

# Просмотр переменных окружения
docker exec jobot_app_container env

# Перезапуск приложения
docker compose -f deploy/debug/docker-compose.yaml restart jobot-app
```

## 🎯 Best Practices

### 1. Разработка
- Используйте `make dev-setup` для быстрого старта
- Проверяйте логи при возникновении проблем
- Используйте health checks для мониторинга

### 2. Тестирование
- Запускайте тесты локально: `make test`
- Используйте отдельную БД для тестов
- Очищайте данные после тестов

### 3. Отладка
- Включайте debug логирование
- Используйте `docker exec` для отладки
- Проверяйте переменные окружения

## 📚 Дополнительные ресурсы

- [Docker Compose документация](https://docs.docker.com/compose/)
- [PostgreSQL документация](https://www.postgresql.org/docs/)
- [pgAdmin документация](https://www.pgadmin.org/docs/)
- [Go Docker best practices](https://docs.docker.com/language/golang/)

---

**Удачной разработки! 🚀**
