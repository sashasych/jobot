# Services Layer

Бизнес-логика приложения разделена на отдельные сервисы по доменным сущностям.

## Структура

```
internal/service/
├── user/          # Управление пользователями
├── employee/      # Управление сотрудниками
├── employer/      # Управление работодателями
├── resume/        # Управление резюме
├── vacancy/       # Управление вакансиями
├── reaction/      # Управление реакциями на вакансии
└── models/        # Модели сервисного слоя
```

## Принципы

- Каждый сервис инкапсулирован в отдельный пакет
- Сервисы взаимодействуют с репозиториями для доступа к данным
- Вся бизнес-логика находится в сервисах
- Сервисы не знают о деталях HTTP/API слоя

## Созданные сервисы

### UserService
**Файл:** `internal/service/user/user.go`

**Методы:**
- `CreateUser(ctx, user)` - создание пользователя
- `GetUser(ctx, id)` - получение пользователя по ID
- `UpdateUser(ctx, req, id)` - обновление пользователя
- `DeleteUser(ctx, id)` - удаление пользователя

### EmployeeService
**Файл:** `internal/service/employee/employee.go`

**Методы:**
- `CreateEmployee(ctx, employee)` - создание сотрудника
- `GetEmployee(ctx, id)` - получение сотрудника по ID
- `UpdateEmployee(ctx, req, id)` - обновление сотрудника
- `DeleteEmployee(ctx, id)` - удаление сотрудника

**Особенности:**
- Генерирует UUID для employee_id
- Управляет тегами сотрудника

### EmployerService
**Файл:** `internal/service/employer/employer.go`

**Методы:**
- `CreateEmployer(ctx, employer)` - создание работодателя
- `GetEmployer(ctx, id)` - получение работодателя по ID
- `UpdateEmployer(ctx, req, id)` - обновление работодателя
- `DeleteEmployer(ctx, id)` - удаление работодателя

**Особенности:**
- Генерирует UUID для employer_id
- Управляет информацией о компании (название, описание, сайт, локация, размер)

### ResumeService
**Файл:** `internal/service/resume/resume.go`

**Методы:**
- `CreateResume(ctx, resume)` - создание резюме
- `GetResume(ctx, id)` - получение резюме по ID
- `UpdateResume(ctx, req, id)` - обновление резюме
- `DeleteResume(ctx, id)` - удаление резюме

**Особенности:**
- Генерирует UUID для resume_id
- Управляет Telegram file ID

### VacancyService
**Файл:** `internal/service/vacancy/vacancy.go`

**Методы:**
- `CreateVacancy(ctx, vacancy)` - создание вакансии
- `GetVacancy(ctx, id)` - получение вакансии по ID
- `GetVacanciesByEmployer(ctx, employerID)` - получение всех вакансий работодателя
- `UpdateVacancy(ctx, req, id)` - обновление вакансии
- `DeleteVacancy(ctx, id)` - удаление вакансии

**Особенности:**
- Генерирует UUID для vacancy_id
- Управляет тегами, названием, описанием, локацией и зарплатой
- Возвращает список вакансий по работодателю

### ReactionService
**Файл:** `internal/service/reaction/reaction.go`

**Методы:**
- `CreateReaction(ctx, reaction)` - создание реакции
- `GetReaction(ctx, id)` - получение реакции по ID
- `GetReactionsByEmployee(ctx, employeeID)` - получение всех реакций сотрудника
- `DeleteReaction(ctx, id)` - удаление реакции

**Особенности:**
- Генерирует UUID для reaction_id
- Возвращает список реакций по сотруднику
- Реакция имеет два типа: like/dislike

## Использование

### Пример создания сервиса

```go
// Создание репозитория
employeeRepo := employeerepository.NewEmployeeRepository(db)

// Создание сервиса
employeeService := employee.NewEmployeeService(employeeRepo)

// Использование сервиса
err := employeeService.CreateEmployee(ctx, &models.Employee{
    UserID: "user-uuid",
    Tags:   []string{"golang", "backend"},
})
```

### Пример обновления

```go
err := employeeService.UpdateEmployee(ctx, &models.EmployeeUpdateRequest{
    Tags: &[]string{"golang", "backend", "postgresql"},
}, employeeID)
```

## Зависимости

Все сервисы зависят от:
- `internal/repository` - интерфейсы репозиториев
- `internal/service/models` - модели сервисного слоя
- `context.Context` - для передачи контекста
- `github.com/google/uuid` - для генерации UUID

## Best Practices

1. **Инкапсуляция**: Каждый сервис работает только со своими данными
2. **Единая ответственность**: Один сервис - одна доменная сущность
3. **Обработка ошибок**: Все ошибки оборачиваются с контекстом через `fmt.Errorf`
4. **Timestamps**: `CreatedAt` и `UpdatedAt` устанавливаются автоматически
5. **UUID**: Генерация UUID происходит в сервисах, а не в контроллерах

## Следующие шаги

После создания сервисов необходимо:
1. ✅ Создать репозитории для каждого сервиса
2. ✅ Создать контроллеры API для каждого сервиса
3. ✅ Добавить валидацию входных данных
4. ✅ Написать unit-тесты для сервисов
5. ✅ Добавить логирование

