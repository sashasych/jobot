package application

import (
	"time"

	"jobot/pkg/database"
	"jobot/pkg/logger"
)

// Config - основная конфигурация приложения
type Config struct {
	// HTTP Server конфигурация
	HTTP HTTPConfig `env:", prefix=HTTP_"`

	// Database конфигурация
	Database database.Config `env:", prefix=DB_"`

	// Logger конфигурация
	Logger logger.Config `env:", prefix=LOG_"`

	// Application конфигурация
	App AppConfig `env:", prefix=APP_"`

	// JWT конфигурация (для будущей аутентификации)
	JWT JWTConfig `env:", prefix=JWT_"`
}

// HTTPConfig - конфигурация HTTP сервера
type HTTPConfig struct {
	Host         string        `env:"HOST" default:"0.0.0.0"`
	Port         string        `env:"PORT" default:"8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"30s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"30s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" default:"120s"`
}

// AppConfig - конфигурация приложения
type AppConfig struct {
	Name        string `env:"NAME" default:"jobot"`
	Version     string `env:"VERSION" default:"1.0.0"`
	Environment string `env:"ENV" default:"development"`
	Debug       bool   `env:"DEBUG" default:"false"`
}

// JWTConfig - конфигурация JWT токенов
type JWTConfig struct {
	Secret     string        `env:"SECRET" required:"true"`
	Expiration time.Duration `env:"EXPIRATION" default:"24h"`
	Issuer     string        `env:"ISSUER" default:"jobot"`
}

// GetAddress возвращает полный адрес HTTP сервера
func (c *HTTPConfig) GetAddress() string {
	return c.Host + ":" + c.Port
}

// IsDevelopment проверяет, является ли окружение development
func (c *AppConfig) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction проверяет, является ли окружение production
func (c *AppConfig) IsProduction() bool {
	return c.Environment == "production"
}

// GetLogLevel возвращает уровень логирования в зависимости от окружения
func (c *AppConfig) GetLogLevel() string {
	if c.IsDevelopment() || c.Debug {
		return "debug"
	}
	return "info"
}
