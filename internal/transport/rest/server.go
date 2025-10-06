package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"jobot/internal/api"
)

const readHeaderTimeoutSeconds = 5

type ConfigHTTPServer struct {
	Port string
	Host string
}

type HandlersConfig struct {
	UserController     api.UserController
	EmployeeController api.EmployeeController
	ResumeController   api.ResumeController
	EmployerController api.EmployerController
	VacancyController  api.VacancyController
	ReactionController api.ReactionController
}

// CreateHTTPServerWithChi creates HTTP server with Chi router
func CreateHTTPServerWithChi(ctx context.Context, cfg *ConfigHTTPServer, handlersConfig *HandlersConfig) *http.Server {
	r := chi.NewRouter()

	// Basic middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{
			"status":  "ok",
			"service": "jobot",
		})
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		// User routes
		r.Route("/user", func(r chi.Router) {
			r.Post("/", handlersConfig.UserController.CreateUser)
			r.Put("/", handlersConfig.UserController.UpdateUser)

			r.Route("/{UserID}", func(r chi.Router) {
				r.Get("/", handlersConfig.UserController.GetUser)
				r.Delete("/", handlersConfig.UserController.DeleteUser)
			})
		})

		// Employee routes
		r.Route("/employee", func(r chi.Router) {
			r.Post("/", handlersConfig.EmployeeController.CreateEmployee)

			r.Route("/{EmployeeID}", func(r chi.Router) {
				r.Get("/reactions", handlersConfig.EmployeeController.GetEmployeeListReactions)
				r.Get("/", handlersConfig.EmployeeController.GetEmployee)
				r.Put("/", handlersConfig.EmployeeController.UpdateEmployee)
				r.Delete("/", handlersConfig.EmployeeController.DeleteEmployee)
			})
		})

		// Resume routes
		r.Route("/resumes", func(r chi.Router) {
			r.Post("/", handlersConfig.ResumeController.CreateResume)
			r.Get("/user/{EmployeeID}", handlersConfig.ResumeController.GetEmployeeListResumes)

			r.Route("/{ResumeID}", func(r chi.Router) {
				r.Get("/", handlersConfig.ResumeController.GetResume)
				r.Put("/", handlersConfig.ResumeController.UpdateResume)
				r.Delete("/", handlersConfig.ResumeController.DeleteResume)
			})
		})

		// Employer routes
		r.Route("/employer", func(r chi.Router) {
			r.Post("/", handlersConfig.EmployerController.CreateEmployer)

			r.Route("/{EmployerID}", func(r chi.Router) {
				r.Get("/", handlersConfig.EmployerController.GetEmployer)
				r.Get("/vacansies", handlersConfig.EmployerController.GetEmployerListVacansies)
				r.Put("/", handlersConfig.EmployerController.UpdateEmployer)
				r.Delete("/", handlersConfig.EmployerController.DeleteEmployer)
			})
		})

		// Job posting routes (vacansies)
		r.Route("/vacansies", func(r chi.Router) {
			r.Post("/", handlersConfig.VacancyController.CreateVacansy)
			r.Get("/", handlersConfig.VacancyController.GetVacansyList)

			r.Route("/{VacansieID}", func(r chi.Router) {
				r.Get("/", handlersConfig.VacancyController.GetVacansy)
				r.Put("/", handlersConfig.VacancyController.UpdateVacansy)
				r.Delete("/", handlersConfig.VacancyController.DeleteVacansy)
			})
		})

		// Reaction routes
		r.Route("/reactions", func(r chi.Router) {
			r.Post("/", handlersConfig.ReactionController.CreateReaction)
		})
	})

	// Print all registered routes (for debugging)
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		return nil
	})

	return &http.Server{
		ReadHeaderTimeout: readHeaderTimeoutSeconds * time.Second,
		Addr:              cfg.Host + ":" + cfg.Port,
		Handler:           r,
	}
}
