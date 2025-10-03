# User Repository

Реализация репозитория для работы с пользователями в PostgreSQL.

## Структура

```
internel/repository/
├── repository.go          # Интерфейс UserRepository
├── models/
│   └── user.go           # Модели User и UserRole
└── user/
    └── repo.go           # Реализация UserRepository для PostgreSQL
```

## Использование

### 1. Подключение к базе данных

```go
import (
    "context"
    "jobot/pkg/database"
)

ctx := context.Background()

cfg := database.Config{
    Host:     "localhost",
    Port:     "5432",
    User:     "postgres",
    Password: "postgres",
    DBName:   "jobot",
    SSLMode:  "disable",
}

pool, err := database.NewPostgresPool(ctx, cfg)
if err != nil {
    log.Fatal(err)
}
defer database.Close(pool)
```

### 2. Создание репозитория

```go
import "jobot/internel/repository/user"

userRepo := user.NewUserRepository(pool)
```

### 3. Создание пользователя

```go
import (
    "jobot/internel/repository/models"
    "golang.org/x/crypto/bcrypt"
)

newUser := &models.User{
    Email:    "user@example.com",
    IsActive: true,
}

// Хешируем пароль
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
newUser.PasswordHash = string(hashedPassword)

err := userRepo.CreateUser(ctx, newUser)
```

### 4. Получение пользователя

```go
// По ID
user, err := userRepo.GetUserByID(ctx, userID)

// По Email
user, err := userRepo.GetUserByEmail(ctx, "user@example.com")
```

### 5. Обновление пользователя

```go
user.IsActive = false
err := userRepo.UpdateUser(ctx, user)
```

### 6. Удаление пользователя

```go
err := userRepo.DeleteUser(ctx, userID)
```

### 7. Работа с ролями

```go
// Создание роли
userRole := &models.UserRole{
    UserID: user.ID,
    Role:   "employee", // или "employer", "admin"
}
err := userRepo.CreateUserRole(ctx, userRole)

// Получение ролей пользователя
roles, err := userRepo.GetUserRoles(ctx, user.ID)
```

## Миграции

SQL-миграции находятся в `migrations/001_create_users_table.sql`

Для применения миграций можно использовать:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)
- Или любой другой инструмент для миграций

Пример с golang-migrate:

```bash
# Установка
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Применение миграций
migrate -database "postgres://postgres:postgres@localhost:5432/jobot?sslmode=disable" \
        -path migrations up
```

## Обработка ошибок

Репозиторий возвращает специфичные ошибки:

```go
import "jobot/internel/repository/user"

user, err := userRepo.GetUserByEmail(ctx, email)
if errors.Is(err, user.ErrUserNotFound) {
    // Пользователь не найден
}

err = userRepo.CreateUser(ctx, newUser)
if errors.Is(err, user.ErrUserAlreadyExists) {
    // Пользователь с таким email уже существует
}
```

## Архитектура

### Разделение ответственности

- **users** - базовая таблица для аутентификации
  - Email, пароль (hash), статус активности
  
- **user_roles** - роли пользователей (многие ко многим)
  - Один пользователь может иметь несколько ролей
  
- **employees** - профиль соискателя
  - Связан с users через user_id
  - Содержит персональную информацию
  
- **employers** - профиль работодателя
  - Связан с users через user_id
  - Связан с companies через company_id

### Преимущества такой архитектуры

1. **Единая точка аутентификации** - один механизм для всех ролей
2. **Гибкость ролей** - пользователь может быть и employee и employer одновременно
3. **Отсутствие дублирования** - общие данные хранятся в users
4. **Легкая масштабируемость** - легко добавить новые роли

## Пример полного использования

См. `examples/user_repository_example.go`

```bash
# Запуск примера
go run examples/user_repository_example.go
```

