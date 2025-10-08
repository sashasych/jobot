# 🛣️ API Routes Reference - Jobot

Справочник всех API путей приложения Jobot.

## ⚠️ Важно: Используется множественное число (plural)

Все пути API используют **множественное число** в соответствии с REST API best practices:
- ✅ `/api/users` (не `/api/user`)
- ✅ `/api/employees` (не `/api/employee`)
- ✅ `/api/employers` (не `/api/employer`)
- ✅ `/api/resumes` (не `/api/resume`)
- ✅ `/api/vacancies` (не `/api/vacancy`)
- ✅ `/api/reactions` (не `/api/reaction`)

---

## 📋 Полный список endpoints

### 🏥 Health Check (1 endpoint)

```
GET  /health
```

---

### 👤 Users - Пользователи (4 endpoints)

```
POST   /api/users               # Создать пользователя
GET    /api/users/{UserID}      # Получить пользователя по ID
PUT    /api/users/{UserID}      # Обновить пользователя
DELETE /api/users/{UserID}      # Удалить пользователя
```

**Параметры пути:**
- `{UserID}` - UUID пользователя

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/users -d '{"tg_chat_id":"123","role":"employee"}'
curl http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440001
```

---

### 👨‍💼 Employees - Сотрудники (5 endpoints)

```
POST   /api/employees                       # Создать профиль сотрудника
GET    /api/employees/{EmployeeID}          # Получить сотрудника по ID
GET    /api/employees/{EmployeeID}/reactions # Получить реакции сотрудника
PUT    /api/employees/{EmployeeID}          # Обновить сотрудника
DELETE /api/employees/{EmployeeID}          # Удалить сотрудника
```

**Параметры пути:**
- `{EmployeeID}` - UUID сотрудника (или UserID для получения по user_id)

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/employees -d '{"user_id":"...","tags":["golang"]}'
curl http://localhost:8080/api/employees/660e8400-e29b-41d4-a716-446655440001
curl http://localhost:8080/api/employees/660e8400-e29b-41d4-a716-446655440001/reactions
```

---

### 🏢 Employers - Работодатели (5 endpoints)

```
POST   /api/employers                        # Создать профиль работодателя
GET    /api/employers/{EmployerID}           # Получить работодателя по ID
GET    /api/employers/{EmployerID}/vacansies # Получить вакансии работодателя
PUT    /api/employers/{EmployerID}           # Обновить работодателя
DELETE /api/employers/{EmployerID}           # Удалить работодателя
```

**Параметры пути:**
- `{EmployerID}` - UUID работодателя (или UserID для получения по user_id)

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/employers -d '{"user_id":"...","company_name":"TechCorp",...}'
curl http://localhost:8080/api/employers/770e8400-e29b-41d4-a716-446655440001
curl http://localhost:8080/api/employers/770e8400-e29b-41d4-a716-446655440001/vacansies
```

---

### 📄 Resumes - Резюме (4 endpoints)

```
POST   /api/resumes                 # Загрузить резюме
GET    /api/resumes/{ResumeID}      # Получить резюме по ID
PUT    /api/resumes/{ResumeID}      # Обновить резюме
DELETE /api/resumes/{ResumeID}      # Удалить резюме
```

**Параметры пути:**
- `{ResumeID}` - UUID резюме (или EmployeeID для получения по employee_id)

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/resumes -d '{"employee_id":"...","tg_file_id":"BAADAgAD..."}'
curl http://localhost:8080/api/resumes/880e8400-e29b-41d4-a716-446655440001
```

---

### 💼 Vacancies - Вакансии (5 endpoints)

```
POST   /api/vacancies                # Создать вакансию
GET    /api/vacancies                # Получить список всех вакансий
GET    /api/vacancies/{VacancyID}    # Получить вакансию по ID
PUT    /api/vacancies/{VacancyID}    # Обновить вакансию
DELETE /api/vacancies/{VacancyID}    # Удалить вакансию
```

**Параметры пути:**
- `{VacancyID}` - UUID вакансии

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/vacancies -d '{"employer_id":"...","title":"Dev","tags":["go"],...}'
curl http://localhost:8080/api/vacancies
curl http://localhost:8080/api/vacancies/990e8400-e29b-41d4-a716-446655440001
```

---

### 👍 Reactions - Реакции (2 endpoints)

```
POST   /api/reactions                            # Создать реакцию (лайк на вакансию)
GET    /api/employees/{EmployeeID}/reactions    # Получить реакции сотрудника
```

**Параметры пути:**
- `{EmployeeID}` - UUID сотрудника

**Примеры:**
```bash
curl -X POST http://localhost:8080/api/reactions -d '{"employee_id":"...","vacansie_id":"...","reaction":"like"}'
curl http://localhost:8080/api/employees/660e8400-e29b-41d4-a716-446655440001/reactions
```

---

## 📊 Итоговая статистика

- **Всего endpoints**: 26
- **Health check**: 1
- **Users**: 4
- **Employees**: 5 (включая вложенный /reactions)
- **Employers**: 5 (включая вложенный /vacansies)
- **Resumes**: 4
- **Vacancies**: 5
- **Reactions**: 2

---

## 🎯 Вложенные ресурсы (Nested Resources)

### Реакции сотрудника
```
GET /api/employees/{EmployeeID}/reactions
```
Вместо `/api/reactions/employee/{EmployeeID}` используется вложенный ресурс.

### Вакансии работодателя
```
GET /api/employers/{EmployerID}/vacansies
```
Вместо `/api/vacancies/employer/{EmployerID}` используется вложенный ресурс.

**Преимущества вложенных ресурсов:**
- ✅ Более понятная иерархия
- ✅ RESTful дизайн
- ✅ Семантически правильно: "реакции принадлежат сотруднику"

---

## 📝 Примечания

### 1. Naming Convention

| ❌ Неправильно | ✅ Правильно |
|---------------|--------------|
| `/api/user` | `/api/users` |
| `/api/employee` | `/api/employees` |
| `/api/vacancy` | `/api/vacancies` |
| `/api/reaction` | `/api/reactions` |
| `/api/reactions/{id}` | `/api/employees/{id}/reactions` |

### 2. Опечатка: vacansies vs vacancies

**⚠️ Внимание:** В коде используется `vacansies` (с опечаткой) вместо `vacancies`:
```
GET /api/employers/{id}/vacansies  ← используется в коде
```

Правильное написание было бы `vacancies`, но для consistency с остальным кодом (поля `vacansie_id` и т.д.) используется `vacansies`.

### 3. Path Parameters

Все параметры пути используют CamelCase:
- `{UserID}` (не `{userId}` или `{user_id}`)
- `{EmployeeID}` (не `{employeeId}`)
- `{EmployerID}`
- `{ResumeID}`
- `{VacancyID}` (но параметр называется `VacansyID` в коде)

---

## 🔍 Как найти нужный endpoint

1. **Swagger UI**: http://localhost:8080/api/docs
   - Визуальный поиск по категориям
   - Фильтр endpoints

2. **OpenAPI спецификация**: 
   - YAML: http://localhost:8080/api/swagger.yaml
   - JSON: http://localhost:8080/api/swagger.json

3. **Этот файл**: быстрый справочник

---

## 🚀 Быстрые тесты

```bash
# Health check
curl http://localhost:8080/health

# Создать пользователя-сотрудника
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"tg_chat_id":"111","role":"employee"}'

# Создать пользователя-работодателя
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"tg_chat_id":"222","role":"employer"}'

# Получить все вакансии
curl http://localhost:8080/api/vacancies

# Создать реакцию
curl -X POST http://localhost:8080/api/reactions \
  -H "Content-Type: application/json" \
  -d '{"employee_id":"...","vacansie_id":"...","reaction":"like"}'
```

---

## 📖 Дополнительная документация

- [README.md](../README.md) - Основная документация
- [API_OVERVIEW.md](../API_OVERVIEW.md) - Обзор API
- [SWAGGER_GUIDE.md](../SWAGGER_GUIDE.md) - Руководство по Swagger
- [POSTMAN_TESTING.md](../POSTMAN_TESTING.md) - Тестирование через Postman

---

**Все пути используют множественное число (plural) ✅**

Это соответствует REST API best practices и используется всеми крупными API (GitHub, Google, Stripe, и т.д.).

