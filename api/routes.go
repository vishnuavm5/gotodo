package api

import (
	"hellocheck/handlers"
	"hellocheck/pkg"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func loadRoutes() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	router.Route("/user", userRoutes)
	router.Route("/todo", todoRoutes)

	return router

}

func userRoutes(router chi.Router) {
	userHandler := &handlers.Users{}

	router.Post("/create", userHandler.CreateNewUser)
	router.Post("/login", userHandler.Login)

}

func todoRoutes(router chi.Router) {
	todoHandler := &handlers.Todo{}

	router.Post("/create", pkg.MiddlewareAuth(todoHandler.CreateTodo))
	router.Get("/", pkg.MiddlewareAuth(todoHandler.GetAllTodos))
	router.Get("/{id}", pkg.MiddlewareAuth(todoHandler.GetTodoById))
	router.Put("/{id}", pkg.MiddlewareAuth(todoHandler.UpdateTodoById))
	router.Delete("/{id}", pkg.MiddlewareAuth(todoHandler.DeleteById))

}
