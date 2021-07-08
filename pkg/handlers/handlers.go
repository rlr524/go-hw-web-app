/*
Package handlers implements all route handlers and uses the template package to render templates.
*/
package handlers

import (
	"github.com/rlr524/go-hw-web-app/pkg/template"
	"net/http"
)

// Home is the / page handler. All web handler functions need to take in, as params,
// the ResponseWriter(w) method and a pointer to the Request(r) method
func Home(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "home.page.gohtml")
}

// About is the /about page handler
func About(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "about.page.gohtml")
}

// Madison is the /madison page handler
func Madison(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "madison.page.gohtml")
}
