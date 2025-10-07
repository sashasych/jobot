package application

import (
	"context"
	"errors"
	"fmt"
	"jobot/internal/api/controllers"
	"jobot/internal/transport/rest"
	"jobot/pkg/database"
	"jobot/pkg/logger"
	"net/http"
	"sync"
	"time"

	api "jobot/internal/api"
	employeeRepo "jobot/internal/repository/employee"
	employerRepo "jobot/internal/repository/employer"
	reactionRepo "jobot/internal/repository/reaction"
	resumeRepo "jobot/internal/repository/resume"
	userRepo "jobot/internal/repository/user"
	vacancyRepo "jobot/internal/repository/vacancy"
	employeeSrv "jobot/internal/service/employee"
	employerSrv "jobot/internal/service/employer"
	reactionSrv "jobot/internal/service/reaction"
	resumeSrv "jobot/internal/service/resume"
	userSrv "jobot/internal/service/user"
	vacancySrv "jobot/internal/service/vacancy"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

const (
	shutDownTimeout = 10 * time.Second
)

type Application struct {
	config     *Config
	logger     *logger.Logger
	serverHTTP *http.Server
	db         *pgxpool.Pool
	controller *api.Controller
}

// NewApplication создает новое приложение с загруженной конфигурацией
func NewApplication(log *logger.Logger) (*Application, error) {
	// Загружаем конфигурацию из переменных окружения
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration from env: %w", err)
	}

	// Создаем логгер
	// TODO: добавить в logger.Logger
	log.SetLevel(cfg.App.GetLogLevel())

	// Создаем пул соединений с БД
	db, err := database.NewPostgresPool(context.Background(), cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Application{
		config: cfg,
		logger: log,
		db:     db,
	}, nil
}

// Initialize инициализирует приложение (создает контроллеры, сервисы и т.д.)
func (app *Application) Initialize(ctx context.Context) error {
	app.logger.Info("Initializing application",
		zap.String("app_name", app.config.App.Name),
		zap.String("version", app.config.App.Version),
		zap.String("environment", app.config.App.Environment),
		zap.Bool("debug", app.config.App.Debug),
	)

	// TODO: Создаем репозитории, сервисы и контроллеры
	app.InitializeControllers()

	/*
		handlersConfig := &rest.HandlersConfig{
			UserController:     userController,
			EmployeeController: employeeController,
			ResumeController:   resumeController,
			EmployerController: employerController,
			VacancyController:  vacancyController,
			ReactionController: reactionController,
		}
	*/

	// Создаем HTTP сервер
	cfgHTTP := &rest.ConfigHTTPServer{
		Port: app.config.HTTP.Port,
		Host: app.config.HTTP.Host,
		// ReadTimeout:  app.config.HTTP.ReadTimeout,  // TODO: добавить в rest.ConfigHTTPServer
		// WriteTimeout: app.config.HTTP.WriteTimeout, // TODO: добавить в rest.ConfigHTTPServer
		// IdleTimeout:  app.config.HTTP.IdleTimeout,  // TODO: добавить в rest.ConfigHTTPServer
	}

	app.serverHTTP = rest.CreateHTTPServerWithChi(ctx, cfgHTTP, app.controller)

	app.logger.Info("Application initialized successfully")
	return nil
}

func (app *Application) InitializeControllers() error {
	userRepository := userRepo.NewUserRepository(app.db)
	employeeRepository := employeeRepo.NewEmployeeRepository(app.db)
	resumeRepository := resumeRepo.NewResumeRepository(app.db)
	employerRepository := employerRepo.NewEmployerRepository(app.db)
	vacancyRepository := vacancyRepo.NewVacancyRepository(app.db)
	reactionRepository := reactionRepo.NewReactionRepository(app.db)

	userService := userSrv.NewUserService(userRepository)
	employeeService := employeeSrv.NewEmployeeService(employeeRepository)
	resumeService := resumeSrv.NewResumeService(resumeRepository)
	employerService := employerSrv.NewEmployerService(employerRepository)
	vacancyService := vacancySrv.NewVacancyService(vacancyRepository)
	reactionService := reactionSrv.NewReactionService(reactionRepository)

	userController := controllers.NewUserController(userService)
	employeeController := controllers.NewEmployeeController(employeeService)
	resumeController := controllers.NewResumeController(resumeService)
	employerController := controllers.NewEmployerController(employerService)
	vacancyController := controllers.NewVacancyController(vacancyService)
	reactionController := controllers.NewReactionController(reactionService)

	app.controller = &api.Controller{
		UserController:     userController,
		EmployeeController: employeeController,
		ResumeController:   resumeController,
		EmployerController: employerController,
		VacancyController:  vacancyController,
		ReactionController: reactionController,
	}

	return nil
}

// Start запускает приложение
func (app *Application) Start(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc) {
	app.logger.Info("Starting application",
		zap.String("http_address", app.config.HTTP.GetAddress()),
	)

	wg.Add(1)
	go app.startHTTPServer(wg, cancel)

	wg.Add(1)
	go app.gracefulStop(ctx, wg)
}

// startHTTPServer запускает HTTP сервер
func (app *Application) startHTTPServer(wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()

	app.logger.Info("HTTP server starting",
		zap.String("address", app.config.HTTP.GetAddress()),
	)

	err := app.serverHTTP.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.logger.Error("HTTP server failed to start",
			zap.Error(err),
		)
		cancel()
	}
}

// gracefulStop корректно останавливает приложение
func (app *Application) gracefulStop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ctx.Done()

	app.logger.Info("Application shutting down")

	wgShutDown := sync.WaitGroup{}
	wgShutDown.Add(1)

	go func() {
		defer wgShutDown.Done()

		ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), shutDownTimeout)
		defer cancelShutDown()

		app.logger.Info("Shutting down HTTP server")
		if err := app.serverHTTP.Shutdown(ctxShutDown); err != nil {
			app.logger.Error("HTTP server shutdown failed",
				zap.Error(err),
			)
		}

		app.logger.Info("Closing database connections")
		if app.db != nil {
			app.db.Close()
		}
	}()

	wgShutDown.Wait()
	app.logger.Info("Application stopped")
}

// GetConfig возвращает конфигурацию приложения
func (app *Application) GetConfig() *Config {
	return app.config
}

// GetLogger возвращает логгер приложения
func (app *Application) GetLogger() *logger.Logger {
	return app.logger
}

// GetDB возвращает пул соединений с БД
func (app *Application) GetDB() *pgxpool.Pool {
	return app.db
}
