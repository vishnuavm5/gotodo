package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	router http.Handler
}

func New() *App {

	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start() error {
	fmt.Println("Hello world")
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	log.Printf("server is running on port :%v", port)

	err := server.ListenAndServe()

	if err != nil {

		return err
	}
	return err

}
