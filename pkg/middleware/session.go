package middleware

import (
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"net/http"
)

var app *config.AppConfig

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
