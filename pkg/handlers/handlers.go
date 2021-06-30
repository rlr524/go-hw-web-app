package handlers

import (
	"github.com/rlr524/go-hw-web-app/pkg/render"
	"net/http"
)

// Home is the / page handler. All web handler functions need to take in, as params,
// the ResponseWriter method and a pointer to the Request method
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.gohtml")
}

// About is the /about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}

// Madison is the /madison page handler
func Madison(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "madison.gohtml")
}
