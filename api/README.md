# 📚 API Documentation

Документация API для проекта Jobot.

## 🎯 Swagger UI

### Быстрый доступ

После запуска приложения, Swagger UI доступен по адресу:

```
http://localhost:8080/api/docs
```

### Возможности Swagger UI

- ✅ **Интерактивная документация** - визуальное представление всех endpoints
- ✅ **Try it out** - возможность тестировать API прямо из браузера
- ✅ **Схемы данных** - детальное описание всех моделей запросов и ответов
- ✅ **Автоматическая валидация** - проверка типов данных и обязательных полей
- ✅ **Примеры запросов** - готовые примеры для каждого endpoint'а

## 📋 Доступные форматы

### 1. Swagger UI (HTML)
```
GET http://localhost:8080/api/docs
```
Интерактивная документация с возможностью тестирования

### 2. OpenAPI YAML
```
GET http://localhost:8080/api/swagger.yaml
```
Спецификация OpenAPI 3.0 в формате YAML

### 3. OpenAPI JSON
```
GET http://localhost:8080/api/swagger.json
```
Спецификация OpenAPI 3.0 в формате JSON

## 🚀 Использование

### Запуск приложения

```bash
# С Docker
make docker-up

# Локально
make run
```

### Открытие документации

1. Откройте браузер
2. Перейдите на http://localhost:8080/api/docs
3. Изучите доступные endpoints
4. Используйте "Try it out" для тестирования

## 📖 Структура API

### Health Check
```
GET /health - Проверка состояния API
```

### Users (Пользователи)
```
POST   /api/user           - Создать пользователя
GET    /api/user/{id}      - Получить пользователя
PUT    /api/user/{id}      - Обновить пользователя
DELETE /api/user/{id}      - Удалить пользователя
```

### Employees (Сотрудники)
```
POST   /api/employee           - Создать профиль сотрудника
GET    /api/employee/{id}      - Получить сотрудника
PUT    /api/employee/{id}      - Обновить сотрудника
DELETE /api/employee/{id}      - Удалить сотрудника
```

### Employers (Работодатели)
```
POST   /api/employer           - Создать профиль работодателя
GET    /api/employer/{id}      - Получить работодателя
PUT    /api/employer/{id}      - Обновить работодателя
DELETE /api/employer/{id}      - Удалить работодателя
```

### Resumes (Резюме)
```
POST   /api/resume           - Загрузить резюме
GET    /api/resume/{id}      - Получить резюме
PUT    /api/resume/{id}      - Обновить резюме
DELETE /api/resume/{id}      - Удалить резюме
```

### Vacancies (Вакансии)
```
POST   /api/vacancy                    - Создать вакансию
GET    /api/vacancy                    - Получить все вакансии
GET    /api/vacancy/{id}               - Получить вакансию
GET    /api/vacancy/employer/{id}     - Получить вакансии работодателя
PUT    /api/vacancy/{id}               - Обновить вакансию
DELETE /api/vacancy/{id}               - Удалить вакансию
```

### Reactions (Реакции)
```
POST   /api/reaction                   - Создать реакцию (лайк)
GET    /api/reaction/{id}              - Получить реакцию
GET    /api/reaction/employee/{id}    - Получить реакции сотрудника
DELETE /api/reaction/{id}              - Удалить реакцию
```

## 🎨 Примеры использования

### Создание пользователя

**Request:**
```bash
curl -X POST http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{
    "tg_user_name": "john_doe",
    "tg_chat_id": "123456789",
    "role": "employee"
  }'
```

**Response:**
```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "tg_user_name": "john_doe",
    "tg_chat_id": "123456789",
    "is_active": true,
    "is_premium": false,
    "role": "employee",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

### Создание вакансии

**Request:**
```bash
curl -X POST http://localhost:8080/api/vacancy \
  -H "Content-Type: application/json" \
  -d '{
    "employer_id": "770e8400-e29b-41d4-a716-446655440001",
    "tags": ["golang", "kubernetes", "senior"],
    "title": "Senior Backend Developer",
    "description": "We are looking for...",
    "location": "Москва (можно удалённо)",
    "salary": "250,000 - 350,000 руб/месяц"
  }'
```

## 🔧 Импорт в другие инструменты

### Postman

1. Откройте Postman
2. File → Import
3. Вставьте URL: `http://localhost:8080/api/swagger.yaml`
4. Нажмите Import

### Insomnia

1. Откройте Insomnia
2. Application → Import/Export → Import Data
3. From URL: `http://localhost:8080/api/swagger.yaml`
4. Import

### HTTPie

```bash
# Установите HTTPie Desktop
# Откройте Settings → Import → OpenAPI
# Вставьте URL: http://localhost:8080/api/swagger.yaml
```

## 📝 Обновление документации

### Ручное обновление

Файлы находятся в директории `api/`:

- `swagger.yaml` - основная спецификация OpenAPI
- `swagger.json` - JSON версия (автоматически генерируется)
- `swagger-ui.html` - HTML страница с Swagger UI

### Редактирование

1. Откройте `api/swagger.yaml`
2. Внесите изменения
3. Перезапустите приложение
4. Обновите страницу http://localhost:8080/api/docs

### Валидация

Используйте онлайн редактор для валидации:

```
https://editor.swagger.io/
```

1. Откройте редактор
2. File → Import URL
3. Вставьте: `http://localhost:8080/api/swagger.yaml`

## 🎯 Best Practices

### 1. Использование Swagger UI

✅ **Правильно:**
- Тестируйте endpoints с реалистичными данными
- Проверяйте схемы ответов
- Используйте примеры из документации

❌ **Неправильно:**
- Тестировать на production сервере
- Использовать реальные учетные данные
- Игнорировать валидацию

### 2. Тестирование API

```bash
# 1. Проверьте health check
curl http://localhost:8080/health

# 2. Создайте пользователя
curl -X POST http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{"tg_chat_id": "123", "role": "employee"}'

# 3. Получите пользователя
curl http://localhost:8080/api/user/{id}
```

### 3. Работа с UUID

Все ID в системе используют формат UUID v4:

```
550e8400-e29b-41d4-a716-446655440001
```

❌ Не используйте:
- Числовые ID
- Короткие строки
- ID из примеров (создавайте свои)

## 🐛 Troubleshooting

### Swagger UI не открывается

**Проблема:**
```
Cannot GET /api/docs
```

**Решение:**
```bash
# Проверьте, что приложение запущено
make dev-status

# Перезапустите приложение
make docker-rebuild

# Проверьте логи
make docker-logs-app
```

### Файлы swagger не найдены

**Проблема:**
```
404 Not Found - swagger.yaml
```

**Решение:**
```bash
# Проверьте наличие файлов
ls -la api/

# Должны быть:
# swagger.yaml
# swagger.json
# swagger-ui.html

# Если файлов нет, пересоздайте их
```

### CORS ошибки

**Проблема:**
```
Access-Control-Allow-Origin error
```

**Решение:**
Убедитесь, что CORS middleware настроен в `server.go`:

```go
r.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
}))
```

## 📚 Дополнительные ресурсы

### OpenAPI / Swagger
- [OpenAPI Specification](https://swagger.io/specification/)
- [Swagger Editor](https://editor.swagger.io/)
- [Swagger UI Documentation](https://swagger.io/tools/swagger-ui/)

### Инструменты
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [HTTPie](https://httpie.io/)
- [curl](https://curl.se/)

### Валидация
- [OpenAPI Validator](https://apitools.dev/swagger-parser/online/)
- [Swagger Validator](https://validator.swagger.io/)

## 🔗 Связанная документация

- [README.md](../README.md) - основная документация проекта
- [POSTMAN_TESTING.md](../POSTMAN_TESTING.md) - тестирование через Postman
- [migrations/SCHEMA.md](../migrations/SCHEMA.md) - схема базы данных

---

**Удачного использования API! 🚀**

