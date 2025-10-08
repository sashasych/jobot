# 🧪 Тестирование API через Postman

Полное руководство по тестированию API Jobot через Postman.

## 📦 Быстрый старт

### 1. Импорт коллекции

1. Откройте **Postman**
2. Нажмите **Import** (Ctrl+O)
3. Выберите файл `postman_collection.json` из корня проекта
4. Коллекция **"Jobot API v2.0"** будет добавлена

### 2. Настройка переменных

Коллекция уже содержит предустановленные переменные:

| Переменная | Значение по умолчанию | Описание |
|------------|----------------------|----------|
| `base_url` | `http://localhost:8080` | URL API сервера |
| `user_id` | UUID | ID пользователя-сотрудника |
| `employer_user_id` | UUID | ID пользователя-работодателя |
| `employee_id` | UUID | ID профиля сотрудника |
| `employer_id` | UUID | ID профиля работодателя |
| `resume_id` | UUID | ID резюме |
| `vacancy_id` | UUID | ID вакансии |
| `reaction_id` | UUID | ID реакции |

### 3. Автоматическое сохранение ID

Коллекция автоматически сохраняет ID из ответов сервера в переменные. После создания любой сущности её ID будет доступен для использования в следующих запросах.

## 📋 API Endpoints

### 🏥 Health Check

```http
GET /health
```

**Описание:** Проверка состояния API сервиса

**Ответ (200 OK):**
```json
{
  "status": "ok",
  "service": "jobot"
}
```

---

## 👤 Users API

### 1. Create User (Создать пользователя)

```http
POST /api/user
```

**Тело запроса:**
```json
{
  "tg_user_name": "john_doe",
  "tg_chat_id": "123456789",
  "is_active": true,
  "is_premium": false,
  "role": "employee"
}
```

**Поля:**
- `tg_user_name` (string, optional) - имя пользователя в Telegram
- `tg_chat_id` (string, required) - уникальный ID чата Telegram
- `is_active` (boolean) - активен ли аккаунт
- `is_premium` (boolean) - премиум статус
- `role` (string) - роль: `"employee"` или `"employer"`

**Ответ (201 Created):**
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
  },
  "message": ""
}
```

### 2. Get User by ID (Получить пользователя)

```http
GET /api/user/{{user_id}}
```

**Ответ (200 OK):**
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
  },
  "message": ""
}
```

### 3. Update User (Обновить пользователя)

```http
PUT /api/user/{{user_id}}
```

**Тело запроса:**
```json
{
  "tg_user_name": "john_updated",
  "is_active": true,
  "is_premium": true
}
```

**Ответ (200 OK):**
```json
{
  "data": {
    "tg_user_name": "john_updated",
    "is_active": true,
    "is_premium": true
  },
  "message": ""
}
```

### 4. Delete User (Удалить пользователя)

```http
DELETE /api/user/{{user_id}}
```

**Ответ (200 OK):**
```json
{
  "data": null,
  "message": ""
}
```

---

## 👨‍💼 Employees API

### 1. Create Employee (Создать профиль сотрудника)

```http
POST /api/employee
```

**Тело запроса:**
```json
{
  "user_id": "{{user_id}}",
  "tags": ["golang", "postgresql", "docker", "backend", "microservices"]
}
```

**Поля:**
- `user_id` (UUID, required) - ID пользователя
- `tags` (array of strings, required) - навыки, интересы, предпочтения

**Ответ (201 Created):**
```json
{
  "data": {
    "employee_id": "660e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440001",
    "tags": ["golang", "postgresql", "docker", "backend", "microservices"],
    "created_at": "2024-01-15T10:35:00Z",
    "updated_at": "2024-01-15T10:35:00Z"
  },
  "message": ""
}
```

### 2. Get Employee by ID

```http
GET /api/employee/{{employee_id}}
```

### 3. Update Employee

```http
PUT /api/employee/{{employee_id}}
```

**Тело запроса:**
```json
{
  "tags": ["golang", "kubernetes", "ci/cd", "devops"]
}
```

### 4. Delete Employee

```http
DELETE /api/employee/{{employee_id}}
```

---

## 🏢 Employers API

### 1. Create Employer (Создать профиль работодателя)

```http
POST /api/employer
```

**Тело запроса:**
```json
{
  "user_id": "{{employer_user_id}}",
  "company_name": "TechCorp Inc",
  "company_description": "Leading technology company specializing in AI and ML",
  "company_website": "https://techcorp.example.com",
  "company_location": "Москва, Россия",
  "company_size": "51-200"
}
```

**Поля:**
- `user_id` (UUID, required) - ID пользователя
- `company_name` (string, required) - название компании
- `company_description` (string, required) - описание компании
- `company_website` (string, optional) - сайт компании
- `company_location` (string, required) - местоположение
- `company_size` (string, required) - размер компании

**Ответ (201 Created):**
```json
{
  "data": {
    "employer_id": "770e8400-e29b-41d4-a716-446655440001",
    "user_id": "550e8400-e29b-41d4-a716-446655440002",
    "company_name": "TechCorp Inc",
    "company_description": "Leading technology company specializing in AI and ML",
    "company_website": "https://techcorp.example.com",
    "company_location": "Москва, Россия",
    "company_size": "51-200",
    "created_at": "2024-01-15T10:40:00Z",
    "updated_at": "2024-01-15T10:40:00Z"
  },
  "message": ""
}
```

### 2. Get Employer by ID

```http
GET /api/employer/{{employer_id}}
```

### 3. Update Employer

```http
PUT /api/employer/{{employer_id}}
```

**Тело запроса:**
```json
{
  "company_name": "TechCorp International",
  "company_description": "Global leader in AI technology",
  "company_size": "201-500"
}
```

### 4. Delete Employer

```http
DELETE /api/employer/{{employer_id}}
```

---

## 📄 Resumes API

### 1. Create Resume (Загрузить резюме)

```http
POST /api/resume
```

**Тело запроса:**
```json
{
  "employee_id": "{{employee_id}}",
  "tg_file_id": "BAADAgADZAAD1234567890"
}
```

**Поля:**
- `employee_id` (UUID, required) - ID сотрудника
- `tg_file_id` (string, required) - Telegram file ID документа

**Ответ (201 Created):**
```json
{
  "data": {
    "resume_id": "880e8400-e29b-41d4-a716-446655440001",
    "employee_id": "660e8400-e29b-41d4-a716-446655440001",
    "tg_file_id": "BAADAgADZAAD1234567890",
    "created_at": "2024-01-15T10:45:00Z",
    "updated_at": "2024-01-15T10:45:00Z"
  },
  "message": ""
}
```

### 2. Get Resume by ID

```http
GET /api/resume/{{resume_id}}
```

### 3. Update Resume (Обновить резюме)

```http
PUT /api/resume/{{resume_id}}
```

**Тело запроса:**
```json
{
  "tg_file_id": "BAADAgADaAAD0987654321"
}
```

### 4. Delete Resume

```http
DELETE /api/resume/{{resume_id}}
```

---

## 💼 Vacancies API

### 1. Create Vacancy (Создать вакансию)

```http
POST /api/vacancy
```

**Тело запроса:**
```json
{
  "employer_id": "{{employer_id}}",
  "tags": ["golang", "kubernetes", "microservices", "senior"],
  "title": "Senior Backend Developer (Go)",
  "description": "We are looking for an experienced Backend Developer to join our team. Must have 5+ years of experience with Go, Kubernetes, and microservices architecture. Responsibilities include designing and implementing scalable backend services.",
  "location": "Москва (можно удалённо)",
  "salary": "250,000 - 350,000 руб/месяц"
}
```

**Поля:**
- `employer_id` (UUID, required) - ID работодателя
- `tags` (array, required) - теги, навыки
- `title` (string, required) - название вакансии
- `description` (string, required) - описание и требования
- `location` (string, required) - местоположение работы
- `salary` (string, required) - информация о зарплате

**Ответ (201 Created):**
```json
{
  "data": {
    "vacansie_id": "990e8400-e29b-41d4-a716-446655440001",
    "employer_id": "770e8400-e29b-41d4-a716-446655440001",
    "tags": ["golang", "kubernetes", "microservices", "senior"],
    "title": "Senior Backend Developer (Go)",
    "description": "We are looking for...",
    "location": "Москва (можно удалённо)",
    "salary": "250,000 - 350,000 руб/месяц",
    "created_at": "2024-01-15T10:50:00Z",
    "updated_at": "2024-01-15T10:50:00Z"
  },
  "message": ""
}
```

### 2. Get Vacancy by ID

```http
GET /api/vacancy/{{vacancy_id}}
```

### 3. Get All Vacancies (Список всех вакансий)

```http
GET /api/vacancy
```

**Ответ (200 OK):**
```json
{
  "data": {
    "vacansies": [
      {
        "vacansie_id": "990e8400-e29b-41d4-a716-446655440001",
        "employer_id": "770e8400-e29b-41d4-a716-446655440001",
        "tags": ["golang", "kubernetes"],
        "title": "Senior Backend Developer",
        "description": "...",
        "location": "Москва",
        "salary": "250,000 - 350,000 руб",
        "created_at": "2024-01-15T10:50:00Z",
        "updated_at": "2024-01-15T10:50:00Z"
      }
    ]
  },
  "message": ""
}
```

### 4. Get Employer Vacancies (Вакансии работодателя)

```http
GET /api/vacancy/employer/{{employer_id}}
```

**Ответ (200 OK):**
```json
{
  "data": {
    "vacansies": [...],
    "employer_id": "770e8400-e29b-41d4-a716-446655440001"
  },
  "message": ""
}
```

### 5. Update Vacancy

```http
PUT /api/vacancy/{{vacancy_id}}
```

**Тело запроса:**
```json
{
  "title": "Senior Backend Developer (Go) - Updated",
  "salary": "300,000 - 400,000 руб/месяц",
  "tags": ["golang", "kubernetes", "aws", "terraform"]
}
```

### 6. Delete Vacancy

```http
DELETE /api/vacancy/{{vacancy_id}}
```

---

## 👍 Reactions API

### 1. Create Reaction (Поставить лайк на вакансию)

```http
POST /api/reaction
```

**Тело запроса:**
```json
{
  "employee_id": "{{employee_id}}",
  "vacansie_id": "{{vacancy_id}}",
  "reaction": "like"
}
```

**Поля:**
- `employee_id` (UUID, required) - ID сотрудника
- `vacansie_id` (UUID, required) - ID вакансии
- `reaction` (string, required) - тип реакции: `"like"` или `"dislike"`

**Ответ (201 Created):**
```json
{
  "data": {
    "reaction_id": "aa0e8400-e29b-41d4-a716-446655440001",
    "employee_id": "660e8400-e29b-41d4-a716-446655440001",
    "vacansie_id": "990e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-01-15T10:55:00Z"
  },
  "message": ""
}
```

### 2. Get Reaction by ID

```http
GET /api/reaction/{{reaction_id}}
```

### 3. Get Employee Reactions (Все реакции сотрудника)

```http
GET /api/reaction/employee/{{employee_id}}
```

**Ответ (200 OK):**
```json
{
  "data": {
    "reactions_ids": [
      "aa0e8400-e29b-41d4-a716-446655440001",
      "aa0e8400-e29b-41d4-a716-446655440002"
    ],
    "employee_id": "660e8400-e29b-41d4-a716-446655440001"
  },
  "message": ""
}
```

### 4. Delete Reaction (Отменить лайк)

```http
DELETE /api/reaction/{{reaction_id}}
```

---

## 🧪 Тестовые сценарии

### Сценарий 1: Полный флоу сотрудника

**Цель:** Создать пользователя-сотрудника, загрузить резюме и поставить лайк на вакансию

1. **Health Check** - проверить, что API работает
2. **Create User** (role: "employee") - создать пользователя
3. **Create Employee** - создать профиль сотрудника
4. **Create Resume** - загрузить резюме
5. **Get All Vacancies** - посмотреть доступные вакансии
6. **Create Reaction** - поставить лайк на вакансию
7. **Get Employee Reactions** - посмотреть свои лайки

### Сценарий 2: Полный флоу работодателя

**Цель:** Создать пользователя-работодателя и опубликовать вакансию

1. **Health Check** - проверить, что API работает
2. **Create User** (role: "employer") - создать пользователя
3. **Create Employer** - создать профиль компании
4. **Create Vacancy** - опубликовать вакансию
5. **Get Employer Vacancies** - посмотреть свои вакансии
6. **Update Vacancy** - обновить вакансию
7. **Get Vacancy by ID** - проверить изменения

### Сценарий 3: Matching (Соискатель ↔ Вакансия)

**Цель:** Проверить систему лайков

1. Создать сотрудника (см. Сценарий 1, шаги 1-3)
2. Создать работодателя и вакансию (см. Сценарий 2, шаги 2-4)
3. **Create Reaction** - сотрудник лайкает вакансию
4. **Get Employee Reactions** - проверить список лайков
5. **Delete Reaction** - отменить лайк
6. **Get Employee Reactions** - убедиться, что лайк удален

---

## 🔧 Настройка Postman

### Environments (Окружения)

Создайте разные окружения для разных сред:

**Development:**
```json
{
  "base_url": "http://localhost:8080"
}
```

**Staging:**
```json
{
  "base_url": "https://staging-api.jobot.com"
}
```

**Production:**
```json
{
  "base_url": "https://api.jobot.com"
}
```

### Pre-request Scripts

Добавьте в коллекцию для генерации тестовых данных:

```javascript
// Генерация случайного Telegram chat ID
pm.collectionVariables.set('random_chat_id', Math.floor(Math.random() * 1000000000).toString());

// Генерация случайного username
pm.collectionVariables.set('random_username', 'user_' + Math.random().toString(36).substring(7));
```

### Tests (Автоматические тесты)

Уже включены в коллекцию:

```javascript
// Автоматическое сохранение ID из ответов
if (pm.response.code === 201 || pm.response.code === 200) {
    try {
        const jsonData = pm.response.json();
        if (jsonData.data) {
            // Сохранить различные ID
            if (jsonData.data.id) {
                pm.collectionVariables.set('user_id', jsonData.data.id);
            }
            // ... и другие ID
        }
    } catch(e) {
        console.log('Could not parse response');
    }
}
```

Добавьте дополнительные тесты:

```javascript
// Проверка статуса
pm.test("Status code is successful", function () {
    pm.expect(pm.response.code).to.be.oneOf([200, 201]);
});

// Проверка формата ответа
pm.test("Response has data field", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('data');
});

// Проверка UUID
pm.test("ID is valid UUID", function () {
    var jsonData = pm.response.json();
    if (jsonData.data && jsonData.data.id) {
        pm.expect(jsonData.data.id).to.match(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/);
    }
});
```

---

## 🐛 Troubleshooting

### ❌ Connection Refused

**Симптомы:**
```
Error: connect ECONNREFUSED 127.0.0.1:8080
```

**Решение:**
```bash
# Проверить, запущено ли приложение
make dev-status

# Запустить приложение
make docker-up
# или
make run
```

### ❌ 404 Not Found

**Симптомы:**
```json
{
  "error": "Not Found",
  "message": "404 page not found"
}
```

**Решение:**
- Проверьте правильность URL
- Убедитесь, что используется правильная версия API
- Проверьте, что endpoint существует

### ❌ 400 Bad Request

**Симптомы:**
```json
{
  "error": "Bad Request",
  "message": "invalid user id"
}
```

**Решение:**
- Проверьте формат JSON
- Убедитесь, что все обязательные поля заполнены
- Проверьте типы данных (UUID должен быть строкой)
- Проверьте валидацию полей

### ❌ 500 Internal Server Error

**Симптомы:**
```json
{
  "error": "Internal Server Error",
  "message": "failed to create user"
}
```

**Решение:**
```bash
# Проверить логи приложения
make docker-logs-app

# Проверить статус базы данных
make db-status

# Подключиться к БД и проверить таблицы
make db-psql
\dt
```

### ❌ Database Connection Error

**Решение:**
```bash
# Проверить, запущена ли база данных
docker ps | grep postgres

# Перезапустить базу данных
docker compose -f deploy/debug/docker-compose.yaml restart postgres

# Применить миграции
make db-migrate
```

---

## 📊 Мониторинг и отладка

### Логи в реальном времени

```bash
# Все логи
make docker-logs

# Только приложение
make docker-logs-app

# Только база данных
make docker-logs-db

# Последние 50 строк
docker logs --tail 50 -f jobot_app_container
```

### Проверка базы данных

```bash
# Подключиться к БД
make db-psql

# SQL запросы для проверки
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM employees;
SELECT COUNT(*) FROM employers;
SELECT COUNT(*) FROM vacancies;
SELECT COUNT(*) FROM reactions;

# Посмотреть последние записи
SELECT * FROM users ORDER BY created_at DESC LIMIT 5;
```

### Postman Console

1. Откройте **View → Show Postman Console** (Alt+Ctrl+C)
2. Смотрите детали каждого запроса:
   - Headers
   - Request body
   - Response body
   - Status code
   - Response time

---

## 🎯 Best Practices

### 1. Порядок тестирования

✅ **Правильно:**
1. Health Check
2. Create User
3. Create Employee/Employer
4. Create Resume/Vacancy
5. Create Reaction
6. Get operations
7. Update operations
8. Delete operations (в обратном порядке создания)

❌ **Неправильно:**
- Создание Employee до User
- Удаление User до удаления зависимых сущностей
- Тестирование Update/Delete до Create

### 2. Работа с переменными

✅ **Правильно:**
- Использовать `{{variable_name}}`
- Полагаться на автосохранение ID
- Создавать отдельные окружения для dev/staging/prod

❌ **Неправильно:**
- Хардкодить UUID в запросах
- Использовать одни и те же данные во всех окружениях

### 3. Данные для тестирования

✅ **Правильно:**
- Использовать реалистичные данные
- Тестировать граничные случаи
- Очищать тестовые данные после тестов

❌ **Неправильно:**
- Использовать "test", "123" в полях
- Оставлять мусорные данные в БД
- Использовать production данные для тестов

### 4. Валидация ответов

✅ **Правильно:**
```javascript
pm.test("Response structure is correct", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('data');
    pm.expect(jsonData.data).to.have.property('id');
    pm.expect(jsonData.data.id).to.be.a('string');
});
```

❌ **Неправильно:**
- Не проверять ответы
- Проверять только status code
- Игнорировать ошибки

---

## 📚 Дополнительные ресурсы

- [Postman Learning Center](https://learning.postman.com/)
- [REST API Testing Guide](https://www.postman.com/api-platform/api-testing/)
- [Postman Variables](https://learning.postman.com/docs/sending-requests/variables/)
- [Writing Tests in Postman](https://learning.postman.com/docs/writing-scripts/test-scripts/)

---

## 📞 Поддержка

Если возникли проблемы:

1. Проверьте [Troubleshooting](#-troubleshooting) секцию
2. Посмотрите логи: `make docker-logs-app`
3. Проверьте статус сервисов: `make dev-status`
4. Создайте Issue на GitHub с описанием проблемы

---

**Удачного тестирования! 🚀**

Made with ❤️ for Jobot

