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

/*
type HandlersConfig struct {
	UserController     api.UserController
	EmployeeController api.EmployeeController
	ResumeController   api.ResumeController
	EmployerController api.EmployerController
	VacancyController  api.VacancyController
	ReactionController api.ReactionController
}
*/

// CreateHTTPServerWithChi creates HTTP server with Chi router
func CreateHTTPServerWithChi(ctx context.Context, cfg *ConfigHTTPServer, controller *api.Controller) *http.Server {
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

	// Swagger documentation
	r.Get("/api/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger-ui.html")
	})
	r.Get("/api/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		http.ServeFile(w, r, "api/swagger.yaml")
	})
	r.Get("/api/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "api/swagger.json")
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		// User routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", controller.UserController.CreateUser)
			r.Put("/", controller.UserController.UpdateUser)

			r.Route("/{UserID}", func(r chi.Router) {
				r.Get("/", controller.UserController.GetUser)
				r.Delete("/", controller.UserController.DeleteUser)
			})
		})

		// Employee routes
		r.Route("/employees", func(r chi.Router) {
			r.Post("/", controller.EmployeeController.CreateEmployee)

			r.Route("/{EmployeeID}", func(r chi.Router) {
				r.Get("/reactions", controller.ReactionController.GetEmployeeReactions)
				r.Get("/", controller.EmployeeController.GetEmployee)
				r.Put("/", controller.EmployeeController.UpdateEmployee)
				r.Delete("/", controller.EmployeeController.DeleteEmployee)
			})
		})

		// Resume routes
		r.Route("/resumes", func(r chi.Router) {
			r.Post("/", controller.ResumeController.CreateResume)

			r.Route("/{ResumeID}", func(r chi.Router) {
				r.Get("/", controller.ResumeController.GetResume)
				r.Put("/", controller.ResumeController.UpdateResume)
				r.Delete("/", controller.ResumeController.DeleteResume)
			})
		})

		// Employer routes
		r.Route("/employers", func(r chi.Router) {
			r.Post("/", controller.EmployerController.CreateEmployer)

			r.Route("/{EmployerID}", func(r chi.Router) {
				r.Get("/", controller.EmployerController.GetEmployer)
				r.Get("/vacansies", controller.VacancyController.GetEmployerVacancies)
				r.Put("/", controller.EmployerController.UpdateEmployer)
				r.Delete("/", controller.EmployerController.DeleteEmployer)
			})
		})

		// Job posting routes (vacansies)
		r.Route("/vacansies", func(r chi.Router) {
			r.Post("/", controller.VacancyController.CreateVacancy)
			r.Get("/", controller.VacancyController.GetVacancyList)

			r.Route("/{VacansyID}", func(r chi.Router) {
				r.Get("/", controller.VacancyController.GetVacancy)
				r.Put("/", controller.VacancyController.UpdateVacancy)
				r.Delete("/", controller.VacancyController.DeleteVacancy)
			})
		})

		// Reaction routes
		r.Route("/reactions", func(r chi.Router) {
			r.Post("/", controller.ReactionController.CreateReaction)
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
