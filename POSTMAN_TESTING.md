# Тестирование API через Postman

## 🚀 Быстрый старт

### 1. Импорт коллекции

1. Откройте Postman
2. Нажмите **Import**
3. Выберите файл `postman_collection.json`
4. Коллекция "Jobot API" будет добавлена

### 2. Настройка переменных

В коллекции уже настроены переменные:
- `base_url`: `http://localhost:8080`
- `user_id`: `123e4567-e89b-12d3-a456-426614174000`

## 📋 Доступные эндпоинты

### Health Check
```
GET /health
```

**Описание:** Проверка состояния сервиса

**Ответ:**
```json
{
    "status": "ok",
    "service": "jobot"
}
```

### Users API

#### 1. Создание пользователя
```
POST /api/user
```

**Тело запроса:**
```json
{
    "email": "test@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe",
    "role": "employee"
}
```

**Ответ (201 Created):**
```json
{
    "success": true,
    "data": {
        "email": "test@example.com",
        "password": "password123",
        "first_name": "John",
        "last_name": "Doe",
        "role": "employee"
    }
}
```

#### 2. Получение пользователя по ID
```
GET /api/user/{UserID}?user_id={UserID}
```

**Параметры:**
- `UserID` - UUID пользователя в пути
- `user_id` - UUID пользователя в query параметре

**Ответ (200 OK):**
```json
{
    "success": true,
    "data": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "email": "test@example.com",
        "first_name": "John",
        "last_name": "Doe",
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
    }
}
```

#### 3. Обновление пользователя
```
PUT /api/user
```

**Тело запроса:**
```json
{
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "first_name": "Jane",
    "last_name": "Smith",
    "is_active": true
}
```

**Ответ (200 OK):**
```json
{
    "success": true,
    "data": {
        "user_id": "123e4567-e89b-12d3-a456-426614174000",
        "first_name": "Jane",
        "last_name": "Smith",
        "is_active": true
    }
}
```

#### 4. Удаление пользователя
```
DELETE /api/user/{UserID}?user_id={UserID}
```

**Параметры:**
- `UserID` - UUID пользователя в пути
- `user_id` - UUID пользователя в query параметре

**Ответ (200 OK):**
```json
{
    "success": true,
    "data": null
}
```

## 🧪 Пошаговое тестирование

### Шаг 1: Проверка Health Check

1. Выберите запрос **"Health Check"**
2. Нажмите **Send**
3. Ожидаемый ответ: `200 OK` с телом `{"status": "ok", "service": "jobot"}`

### Шаг 2: Создание пользователя

1. Выберите запрос **"Create User"**
2. Убедитесь, что в теле запроса корректные данные
3. Нажмите **Send**
4. Ожидаемый ответ: `201 Created`
5. **Сохраните ID пользователя** из ответа для следующих тестов

### Шаг 3: Получение пользователя

1. Выберите запрос **"Get User by ID"**
2. Замените `{{user_id}}` на реальный ID из шага 2
3. Нажмите **Send**
4. Ожидаемый ответ: `200 OK` с данными пользователя

### Шаг 4: Обновление пользователя

1. Выберите запрос **"Update User"**
2. Замените `{{user_id}}` на реальный ID
3. Измените данные в теле запроса
4. Нажмите **Send**
5. Ожидаемый ответ: `200 OK`

### Шаг 5: Удаление пользователя

1. Выберите запрос **"Delete User"**
2. Замените `{{user_id}}` на реальный ID
3. Нажмите **Send**
4. Ожидаемый ответ: `200 OK`

## 🔧 Настройка Postman

### Переменные окружения

Создайте переменные окружения для разных сред:

**Development:**
```
base_url: http://localhost:8080
user_id: 123e4567-e89b-12d3-a456-426614174000
```

**Production:**
```
base_url: https://api.jobot.com
user_id: 123e4567-e89b-12d3-a456-426614174000
```

### Автоматические тесты

Добавьте тесты в Postman для автоматической проверки:

```javascript
// Проверка статуса ответа
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

// Проверка структуры ответа
pm.test("Response has success field", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('success');
});

// Проверка типа данных
pm.test("Success field is boolean", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.success).to.be.a('boolean');
});
```

## 🐛 Troubleshooting

### Проблема: Connection refused

**Причина:** Приложение не запущено

**Решение:**
```bash
# Запустите приложение
make docker-up

# Или локально
make run
```

### Проблема: 404 Not Found

**Причина:** Неверный URL или маршрут

**Решение:**
- Проверьте `base_url` в переменных
- Убедитесь, что путь к API корректен
- Проверьте, что сервер запущен на правильном порту

### Проблема: 500 Internal Server Error

**Причина:** Ошибка в приложении

**Решение:**
```bash
# Проверьте логи
make docker-logs-app

# Или локально
go run ./cmd/app
```

### Проблема: 400 Bad Request

**Причина:** Неверные данные в запросе

**Решение:**
- Проверьте формат JSON
- Убедитесь, что все обязательные поля заполнены
- Проверьте типы данных

## 📊 Мониторинг

### Логи приложения

```bash
# Просмотр логов в реальном времени
make docker-logs-app

# Последние 100 строк
docker logs --tail 100 jobot_app_container
```

### Метрики

```bash
# Статус контейнеров
make dev-status

# Использование ресурсов
docker stats jobot_app_container
```

## 🎯 Best Practices

### 1. Тестирование

- Всегда начинайте с Health Check
- Тестируйте в правильном порядке (Create → Read → Update → Delete)
- Используйте реальные данные
- Проверяйте все поля ответа

### 2. Отладка

- Включайте логирование в приложении
- Используйте Postman Console для просмотра деталей запроса
- Проверяйте заголовки ответа
- Сохраняйте примеры успешных запросов

### 3. Безопасность

- Не используйте реальные пароли в тестах
- Очищайте тестовые данные после тестирования
- Используйте отдельную среду для тестирования

## 📚 Дополнительные ресурсы

- [Postman Documentation](https://learning.postman.com/)
- [REST API Testing](https://learning.postman.com/docs/writing-scripts/test-scripts/)
- [Environment Variables](https://learning.postman.com/docs/sending-requests/variables/)

---

**Удачного тестирования! 🚀**
