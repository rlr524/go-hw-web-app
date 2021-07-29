package middleware

import (
	"github.com/justinas/nosurf"
	"github.com/rlr524/go-hw-web-app"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   hello_world_web_app.SetEnvironment(),
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
