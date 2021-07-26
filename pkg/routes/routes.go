package routes

import (
	"github.com/bmizerany/pat"
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"github.com/rlr524/go-hw-web-app/pkg/handlers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/madison", http.HandlerFunc(handlers.Repo.Madison))

	return mux
}
