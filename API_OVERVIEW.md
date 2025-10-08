# 🎯 API Overview - Jobot

Краткий обзор всех возможностей API Jobot.

## 🌐 Быстрые ссылки

| Ресурс | URL | Описание |
|--------|-----|----------|
| 🏥 Health Check | http://localhost:8080/health | Проверка состояния API |
| 📖 Swagger UI | http://localhost:8080/api/docs | Интерактивная документация |
| 📄 OpenAPI YAML | http://localhost:8080/api/swagger.yaml | Спецификация в YAML |
| 📋 OpenAPI JSON | http://localhost:8080/api/swagger.json | Спецификация в JSON |
| 🗄️ pgAdmin | http://localhost:5050 | Управление БД |

## 📊 Статистика API

- **Всего endpoints**: 30+
- **Категорий**: 7 (health, users, employees, employers, resumes, vacancies, reactions)
- **Методы**: GET, POST, PUT, DELETE
- **Формат**: REST API, JSON
- **Документация**: OpenAPI 3.0

## 🔗 Endpoints по категориям

### 👤 Users (4 endpoints)
```
POST   /api/user           ← Создать пользователя
GET    /api/user/{id}      ← Получить пользователя
PUT    /api/user/{id}      ← Обновить пользователя
DELETE /api/user/{id}      ← Удалить пользователя
```

### 👨‍💼 Employees (4 endpoints)
```
POST   /api/employee       ← Создать профиль сотрудника
GET    /api/employee/{id}  ← Получить сотрудника
PUT    /api/employee/{id}  ← Обновить сотрудника
DELETE /api/employee/{id}  ← Удалить сотрудника
```

### 🏢 Employers (4 endpoints)
```
POST   /api/employer       ← Создать профиль работодателя
GET    /api/employer/{id}  ← Получить работодателя
PUT    /api/employer/{id}  ← Обновить работодателя
DELETE /api/employer/{id}  ← Удалить работодателя
```

### 📄 Resumes (4 endpoints)
```
POST   /api/resume         ← Загрузить резюме
GET    /api/resume/{id}    ← Получить резюме
PUT    /api/resume/{id}    ← Обновить резюме
DELETE /api/resume/{id}    ← Удалить резюме
```

### 💼 Vacancies (6 endpoints)
```
POST   /api/vacancy                  ← Создать вакансию
GET    /api/vacancy                  ← Все вакансии
GET    /api/vacancy/{id}             ← Получить вакансию
GET    /api/vacancy/employer/{id}   ← Вакансии работодателя
PUT    /api/vacancy/{id}             ← Обновить вакансию
DELETE /api/vacancy/{id}             ← Удалить вакансию
```

### 👍 Reactions (4 endpoints)
```
POST   /api/reaction                 ← Создать реакцию (лайк)
GET    /api/reaction/{id}            ← Получить реакцию
GET    /api/reaction/employee/{id}  ← Реакции сотрудника
DELETE /api/reaction/{id}            ← Удалить реакцию
```

## 🎨 Быстрый тест

### 1. Проверка работоспособности

```bash
curl http://localhost:8080/health
```

**Ожидаемый ответ:**
```json
{"status": "ok", "service": "jobot"}
```

### 2. Создание пользователя

```bash
curl -X POST http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{
    "tg_chat_id": "123456789",
    "tg_user_name": "test_user",
    "role": "employee"
  }'
```

### 3. Открытие Swagger UI

```bash
# Автоматически
make api-docs-open

# Или вручную
http://localhost:8080/api/docs
```

## 📚 Документация

### Основные файлы

| Файл | Описание |
|------|----------|
| [README.md](README.md) | Основная документация проекта |
| [QUICKSTART.md](QUICKSTART.md) | Быстрый старт |
| [SWAGGER_GUIDE.md](SWAGGER_GUIDE.md) | Руководство по Swagger UI |
| [api/README.md](api/README.md) | Документация API |
| [POSTMAN_TESTING.md](POSTMAN_TESTING.md) | Тестирование через Postman |

### Swagger файлы

| Файл | Назначение |
|------|-----------|
| [api/swagger.yaml](api/swagger.yaml) | OpenAPI 3.0 спецификация (YAML) |
| [api/swagger.json](api/swagger.json) | OpenAPI 3.0 спецификация (JSON) |
| [api/swagger-ui.html](api/swagger-ui.html) | Swagger UI HTML страница |

### База данных

| Файл | Описание |
|------|----------|
| [migrations/README.md](migrations/README.md) | Документация миграций |
| [migrations/SCHEMA.md](migrations/SCHEMA.md) | Схема базы данных |
| [migrations/test_data.sql](migrations/test_data.sql) | Тестовые данные |

## 🛠️ Инструменты

### Для тестирования API

1. **Swagger UI** (встроенный) ⭐ Рекомендуется
   - URL: http://localhost:8080/api/docs
   - Интерактивное тестирование
   - Визуальная документация

2. **Postman**
   - Импорт: `postman_collection.json`
   - Или: http://localhost:8080/api/swagger.yaml

3. **curl**
   - Быстрые тесты из командной строки
   - Примеры в документации

4. **HTTPie**
   - Более дружественная альтернатива curl
   ```bash
   http POST localhost:8080/api/user tg_chat_id=123 role=employee
   ```

### Для работы с БД

1. **pgAdmin** (http://localhost:5050)
2. **psql** (`make db-psql`)
3. **DBeaver** (Desktop клиент)
4. **DataGrip** (JetBrains IDE)

## 🚀 Команды для начала работы

```bash
# Запустить всё (БД + Приложение)
make docker-up && make db-migrate

# Открыть Swagger UI
make api-docs-open

# Загрузить тестовые данные
docker compose -f deploy/debug/docker-compose.yaml exec postgres \
  psql -U postgres -d jobot -f /migrations/test_data.sql

# Посмотреть логи
make docker-logs-app

# Подключиться к БД
make db-psql
```

## 📖 Примеры использования

### Сценарий: Создание сотрудника и лайк на вакансию

```bash
# 1. Создать пользователя
USER_ID=$(curl -X POST http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{"tg_chat_id": "111", "role": "employee"}' | jq -r '.data.id')

# 2. Создать профиль сотрудника
EMPLOYEE_ID=$(curl -X POST http://localhost:8080/api/employee \
  -H "Content-Type: application/json" \
  -d "{\"user_id\": \"$USER_ID\", \"tags\": [\"golang\"]}" | jq -r '.data.employee_id')

# 3. Получить список вакансий
curl http://localhost:8080/api/vacancy

# 4. Поставить лайк на вакансию
curl -X POST http://localhost:8080/api/reaction \
  -H "Content-Type: application/json" \
  -d "{\"employee_id\": \"$EMPLOYEE_ID\", \"vacansie_id\": \"VACANCY_ID\", \"reaction\": \"like\"}"

# 5. Посмотреть свои лайки
curl http://localhost:8080/api/reaction/employee/$EMPLOYEE_ID
```

## 🎯 Рекомендуемый workflow

### Для разработчиков

1. ✅ Запустить приложение: `make run`
2. ✅ Открыть Swagger UI: `make api-docs-open`
3. ✅ Изучить endpoints в интерактивном режиме
4. ✅ Протестировать API через Swagger UI
5. ✅ Написать автоматические тесты
6. ✅ Обновить документацию при изменениях

### Для тестировщиков

1. ✅ Импортировать Postman коллекцию
2. ✅ Изучить Swagger документацию
3. ✅ Выполнить тестовые сценарии
4. ✅ Проверить граничные случаи
5. ✅ Задокументировать найденные баги

### Для пользователей API

1. ✅ Открыть Swagger UI: http://localhost:8080/api/docs
2. ✅ Изучить доступные endpoints
3. ✅ Посмотреть примеры запросов
4. ✅ Использовать в своём коде

## 📞 Поддержка

- 📖 Документация: [README.md](README.md)
- 🐛 Issues: https://github.com/sashasych/jobot/issues
- 💬 Вопросы: см. FAQ в основной документации

---

**Вся информация по API всегда доступна в Swagger UI! 📖**

Запустите: `make api-docs-open`

