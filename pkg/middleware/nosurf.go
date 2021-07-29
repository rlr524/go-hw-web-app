package middleware

import (
	"github.com/justinas/nosurf"
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   config.SetEnvironment(),
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
