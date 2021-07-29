package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"github.com/rlr524/go-hw-web-app/pkg/handlers"
	middleware2 "github.com/rlr524/go-hw-web-app/pkg/middleware"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware2.NoSurf)
	mux.Use(middleware2.SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/madison", handlers.Repo.Madison)

	return mux
}
