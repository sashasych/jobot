# 📖 Swagger UI Guide - Jobot API

Руководство по использованию Swagger UI для тестирования и изучения API Jobot.

## 🚀 Быстрый старт

### 1. Запуск приложения

```bash
# С Docker
make docker-up

# Локально
make run
```

### 2. Открытие Swagger UI

Откройте в браузере:
```
http://localhost:8080/api/docs
```

### 3. Изучение документации

На странице вы увидите:
- 📋 Список всех endpoints, сгруппированных по категориям
- 📝 Описание каждого endpoint'а
- 🎯 Параметры запросов
- 📊 Схемы данных
- 💡 Примеры использования

## 🎯 Основные возможности

### 1. Просмотр endpoints

Swagger UI группирует endpoints по тегам:

- **health** - Health check
- **users** - Пользователи
- **employees** - Сотрудники
- **employers** - Работодатели
- **resumes** - Резюме
- **vacancies** - Вакансии
- **reactions** - Реакции

### 2. Тестирование API

#### Пример: Создание пользователя

**Шаг 1:** Раскройте секцию `users`

**Шаг 2:** Найдите `POST /api/users` - "Создать пользователя"

**Шаг 3:** Нажмите **"Try it out"**

**Шаг 4:** Отредактируйте JSON в поле Request body:
```json
{
  "tg_user_name": "test_user",
  "tg_chat_id": "999888777",
  "is_active": true,
  "is_premium": false,
  "role": "employee"
}
```

**Шаг 5:** Нажмите **"Execute"**

**Шаг 6:** Просмотрите результат:
- **Response code**: 201 Created
- **Response body**: данные созданного пользователя
- **Response headers**: заголовки ответа

**Шаг 7:** Скопируйте `id` из ответа для последующих запросов

### 3. Схемы данных

#### Просмотр схем

1. Прокрутите вниз до секции **"Schemas"**
2. Найдите нужную модель (например, `UserCreateRequest`)
3. Раскройте её для просмотра всех полей
4. Изучите типы данных и обязательные поля

#### Пример схемы: UserCreateRequest

```yaml
UserCreateRequest:
  type: object
  required:
    - tg_chat_id
    - role
  properties:
    tg_user_name:
      type: string
      example: "john_doe"
    tg_chat_id:
      type: string
      example: "123456789"
    role:
      type: string
      enum: [employee, employer]
```

**Обязательные поля:** `tg_chat_id`, `role`  
**Опциональные поля:** `tg_user_name`, `is_active`, `is_premium`

## 📝 Тестовые сценарии

### Сценарий 1: Регистрация соискателя

```
1. POST /api/users (role: employee)
   → Скопировать user_id из ответа

2. POST /api/employees (используя user_id)
   → Скопировать employee_id из ответа

3. POST /api/resumes (используя employee_id)
   → Скопировать resume_id из ответа

4. GET /api/employees/{employee_id}
   → Проверить, что профиль создан
```

### Сценарий 2: Публикация вакансии

```
1. POST /api/users (role: employer)
   → Скопировать user_id

2. POST /api/employers (используя user_id)
   → Скопировать employer_id

3. POST /api/vacancies (используя employer_id)
   → Скопировать vacancy_id

4. GET /api/employers/{employer_id}/vacansies
   → Увидеть список вакансий компании
```

### Сценарий 3: Matching (Лайки)

```
1. GET /api/vacancies
   → Получить список всех вакансий
   → Выбрать vacancy_id

2. POST /api/reactions
   {
     "employee_id": "...",
     "vacansie_id": "...",
     "reaction": "like"
   }

3. GET /api/employees/{employee_id}/reactions
   → Увидеть все лайки сотрудника
```

## 🔧 Продвинутые возможности

### 1. Фильтрация endpoints

В Swagger UI есть поиск (обычно сверху):
- Введите название endpoint'а
- Или название модели
- Или тег (users, vacancies, и т.д.)

### 2. Изменение сервера

В верхней части Swagger UI:
1. Найдите выпадающий список **"Servers"**
2. Выберите нужный сервер:
   - `http://localhost:8080` - Development
   - `https://api.jobot.com` - Production

### 3. Сохранение параметров

Swagger UI запоминает:
- ✅ Последние использованные значения
- ✅ Выбранный сервер
- ✅ Раскрытые секции

### 4. Скачивание спецификации

```bash
# YAML
curl http://localhost:8080/api/swagger.yaml > jobot-api.yaml

# JSON
curl http://localhost:8080/api/swagger.json > jobot-api.json
```

### 5. Импорт в другие инструменты

#### Postman
```
File → Import → Link
http://localhost:8080/api/swagger.yaml
```

#### Insomnia
```
Application → Import/Export → Import Data → From URL
http://localhost:8080/api/swagger.yaml
```

#### VS Code (REST Client)
Используйте расширение "REST Client" и импортируйте OpenAPI спецификацию

## 💡 Советы и трюки

### 1. Быстрое тестирование

**Используйте curl для быстрых тестов:**

```bash
# Health check
curl http://localhost:8080/health

# Создание пользователя
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"tg_chat_id": "123", "role": "employee"}'

# Получение вакансий
curl http://localhost:8080/api/vacancies
```

### 2. Работа с UUID

Все ID в системе - это UUID v4. Генерация UUID:

**В браузере (Console):**
```javascript
crypto.randomUUID()
// → "550e8400-e29b-41d4-a716-446655440001"
```

**В командной строке:**
```bash
uuidgen
# или
python3 -c "import uuid; print(uuid.uuid4())"
```

**Онлайн:**
```
https://www.uuidgenerator.net/
```

### 3. Проверка валидации

Swagger UI автоматически валидирует:
- ✅ Обязательные поля
- ✅ Типы данных
- ✅ Форматы (UUID, email, и т.д.)
- ✅ Enum значения

Если данные невалидны, вы увидите предупреждение перед отправкой.

### 4. Экспорт результатов

**Копирование curl команды:**
1. Выполните запрос в Swagger UI
2. Найдите секцию "Request" в результатах
3. Нажмите на кнопку "Copy as cURL"
4. Используйте команду в терминале

## 🐛 Troubleshooting

### Swagger UI не загружается

**Проблема:** Белый экран или "Failed to load API definition"

**Решение:**
```bash
# 1. Проверьте, что приложение запущено
make dev-status

# 2. Проверьте наличие файлов
ls -la api/

# 3. Проверьте логи
make docker-logs-app

# 4. Перезапустите приложение
make docker-rebuild
```

### Ошибка "Failed to fetch"

**Проблема:** При нажатии "Try it out" → "Execute"

**Решение:**
```bash
# 1. Проверьте CORS настройки в server.go
# 2. Проверьте, что API доступен
curl http://localhost:8080/health

# 3. Проверьте логи браузера (F12 → Console)
# 4. Убедитесь, что используется правильный server URL
```

### swagger.yaml не найден

**Проблема:** 404 при загрузке спецификации

**Решение:**
```bash
# Проверьте пути в server.go
# Файлы должны быть в директории api/

# Проверьте права доступа
chmod 644 api/swagger.yaml
chmod 644 api/swagger-ui.html
```

### Старая версия документации

**Проблема:** Изменения в swagger.yaml не отображаются

**Решение:**
```bash
# 1. Очистите кэш браузера (Ctrl+Shift+R)
# 2. Перезапустите приложение
make docker-rebuild

# 3. Проверьте, что файл обновлён
cat api/swagger.yaml | head -20
```

## 📚 Дополнительные ресурсы

### Документация
- [OpenAPI Specification 3.0](https://swagger.io/specification/)
- [Swagger UI Documentation](https://swagger.io/tools/swagger-ui/)
- [Swagger Editor](https://editor.swagger.io/)

### Онлайн инструменты
- [Swagger Editor](https://editor.swagger.io/) - редактор и валидатор
- [Swagger Inspector](https://inspector.swagger.io/) - тестирование API
- [Swagger Hub](https://app.swaggerhub.com/) - хостинг документации

### Альтернативы
- [Redoc](https://github.com/Redocly/redoc) - альтернативный UI для OpenAPI
- [RapiDoc](https://mrin9.github.io/RapiDoc/) - ещё один UI вариант
- [Stoplight](https://stoplight.io/) - платформа для API документации

## 🎓 Обучение

### Как читать Swagger документацию

**Зелёный цвет** - GET запросы (чтение данных)  
**Оранжевый цвет** - POST запросы (создание)  
**Синий цвет** - PUT запросы (обновление)  
**Красный цвет** - DELETE запросы (удаление)

### Понимание схем

```yaml
properties:
  field_name:
    type: string          # Тип данных
    format: uuid          # Формат (если есть)
    example: "value"      # Пример значения
    description: "..."    # Описание поля
```

**Типы данных:**
- `string` - строка
- `integer` - целое число
- `boolean` - true/false
- `array` - массив
- `object` - объект
- `uuid` - UUID формат

### HTTP коды ответов

- **200 OK** - успешный запрос (GET, PUT, DELETE)
- **201 Created** - ресурс создан (POST)
- **400 Bad Request** - неверные данные
- **404 Not Found** - ресурс не найден
- **500 Internal Server Error** - ошибка сервера

## 🔗 Связанная документация

- [README.md](README.md) - основная документация
- [POSTMAN_TESTING.md](POSTMAN_TESTING.md) - тестирование через Postman
- [api/README.md](api/README.md) - документация API
- [QUICKSTART.md](QUICKSTART.md) - быстрый старт

## 💬 Поддержка

Если у вас возникли вопросы:

1. Проверьте [Troubleshooting](#-troubleshooting)
2. Изучите [примеры тестовых сценариев](#-тестовые-сценарии)
3. Посмотрите логи: `make docker-logs-app`
4. Создайте Issue на GitHub

---

**Счастливого тестирования! 🎉**

