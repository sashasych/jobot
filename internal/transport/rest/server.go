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

// CreateHTTPServerWithChi creates HTTP server with Chi router
func CreateHTTPServerWithChi(ctx context.Context, cfg *ConfigHTTPServer, usersController api.UserController) *http.Server {
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

	// Initialize controllers
	//usersController := controllers.NewUsersController(ctx, userService)

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
			r.Post("/", usersController.CreateUser)
			//r.Get("/", usersController.GetListUser)
			r.Put("/", usersController.UpdateUser)

			r.Route("/{UserID}", func(r chi.Router) {
				r.Get("/", usersController.GetUser)
				r.Delete("/", usersController.DeleteUser)
			})
		})

		/*
			// Employee routes
			r.Route("/employee", func(r chi.Router) {
				r.Post("/", usersController.CreateEmployee)
				r.Get("/", usersController.GetListEmployee)

				r.Route("/{EmployeeID}", func(r chi.Router) {
					r.Get("/", usersController.GetEmployee)
					r.Put("/", usersController.UpdateEmployee)
					r.Delete("/", usersController.DeleteEmployee)
				})
			})

			// Resume routes
			r.Route("/resumes", func(r chi.Router) {
				r.Post("/", usersController.CreateResume)
				r.Get("/user/{EmployeeID}", usersController.GetEmployeeListResume)

				r.Route("/{ResumeID}", func(r chi.Router) {
					r.Get("/", usersController.GetResume)
					r.Put("/", usersController.UpdateResume)
					r.Delete("/", usersController.DeleteResume)
				})
			})

			// Employer routes
			r.Route("/employer", func(r chi.Router) {
				r.Post("/", usersController.CreateEmployer)
				r.Get("/", usersController.GetListEmployer)

				r.Route("/{EmployerID}", func(r chi.Router) {
					r.Get("/", usersController.GetEmployer)
					r.Put("/", usersController.UpdateEmployer)
					r.Delete("/", usersController.DeleteEmployer)
				})
			})

			// Job posting routes (vacansies)
			r.Route("/vacansies", func(r chi.Router) {
				r.Post("/", usersController.CreateVacansie)
				r.Get("/employer/{EmployerID}", usersController.GetEmployerListVacansie)

				r.Route("/{VacansieID}", func(r chi.Router) {
					r.Get("/", usersController.GetVacansie)
					r.Put("/", usersController.UpdateVacansie)
					r.Delete("/", usersController.DeleteVacansie)
				})
			})
		*/
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
