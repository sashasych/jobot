# 🚀 Как использовать API Jobot

Самый простой способ начать работу с API.

## ⚡ Быстрый старт (3 шага)

### Шаг 1: Запустите приложение

```bash
make docker-up && make db-migrate && make run
```

### Шаг 2: Откройте Swagger UI

```bash
make api-docs-open
```

Или перейдите в браузере на:
```
http://localhost:8080/api/docs
```

### Шаг 3: Начните тестировать!

На странице Swagger UI:
1. Выберите любой endpoint (например, `POST /api/users`)
2. Нажмите **"Try it out"**
3. Заполните данные
4. Нажмите **"Execute"**
5. Посмотрите результат

## 🎯 Типичные сценарии

### Сценарий 1: Я соискатель

```
1. Создать аккаунт     → POST /api/users (role: employee)
2. Создать профиль     → POST /api/employees
3. Загрузить резюме    → POST /api/resumes
4. Посмотреть вакансии → GET /api/vacancies
5. Поставить лайк      → POST /api/reactions
```

**Попробуйте в Swagger UI:**
http://localhost:8080/api/docs

### Сценарий 2: Я работодатель

```
1. Создать аккаунт       → POST /api/users (role: employer)
2. Создать профиль       → POST /api/employers
3. Опубликовать вакансию → POST /api/vacancies
4. Посмотреть вакансии   → GET /api/employers/{id}/vacansies
5. Обновить вакансию     → PUT /api/vacancies/{id}
```

**Попробуйте в Swagger UI:**
http://localhost:8080/api/docs

## 💡 Полезные команды

```bash
# Показать все API ссылки
make api-docs

# Открыть Swagger UI
make api-docs-open

# Посмотреть логи
make docker-logs-app

# Подключиться к БД
make db-psql

# Перезапустить всё
make docker-rebuild
```

## 📝 Примеры curl

### Создать пользователя
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "tg_chat_id": "123456789",
    "tg_user_name": "john_doe",
    "role": "employee"
  }'
```

### Создать вакансию
```bash
curl -X POST http://localhost:8080/api/vacancies \
  -H "Content-Type: application/json" \
  -d '{
    "employer_id": "YOUR_EMPLOYER_ID",
    "tags": ["golang", "senior"],
    "title": "Senior Go Developer",
    "description": "Looking for Go expert",
    "location": "Remote",
    "salary": "200k - 300k руб"
  }'
```

### Получить все вакансии
```bash
curl http://localhost:8080/api/vacancies
```

## 🛠️ Что выбрать?

| Инструмент | Когда использовать |
|------------|-------------------|
| **Swagger UI** | 🏆 Первое знакомство с API, изучение документации, быстрое тестирование |
| **Postman** | Сложные сценарии, автоматизация, коллекции тестов |
| **curl** | Быстрые проверки, скрипты, CI/CD |
| **HTTPie** | Красивый вывод в терминале, интерактивная работа |

## 🎓 Обучение

### Если вы новичок в API:

1. Начните со **Swagger UI**: http://localhost:8080/api/docs
2. Изучите раздел **Health Check**
3. Попробуйте создать пользователя: `POST /api/users`
4. Получите пользователя: `GET /api/users/{id}`
5. Переходите к более сложным операциям

### Если вы опытный разработчик:

1. Импортируйте OpenAPI спецификацию в ваш инструмент
2. Изучите схемы данных
3. Используйте автоматическую генерацию клиентов
4. Интегрируйте в свой код

## 📖 Подробная документация

- **Swagger Guide**: [SWAGGER_GUIDE.md](SWAGGER_GUIDE.md)
- **API Documentation**: [api/README.md](api/README.md)
- **Postman Testing**: [POSTMAN_TESTING.md](POSTMAN_TESTING.md)
- **Main README**: [README.md](README.md)

## 🆘 Помощь

**API не работает?**
```bash
make dev-status    # Проверить статус
make docker-logs   # Посмотреть логи
```

**Swagger не открывается?**
```bash
make run           # Запустить приложение
make api-docs      # Показать ссылки
```

**Нужны примеры?**
- Swagger UI: http://localhost:8080/api/docs (встроенные примеры)
- Postman коллекция: [postman_collection.json](postman_collection.json)
- Тестовые данные: [migrations/test_data.sql](migrations/test_data.sql)

---

**Начните прямо сейчас!**

```bash
make api-docs-open
```

🎉 **Удачи!**

