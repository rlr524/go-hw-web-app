/*
Package handlers implements all route handlers and uses the template package to render templates.
The package uses the repository pattern
*/
package handlers

import (
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"github.com/rlr524/go-hw-web-app/pkg/models"
	"github.com/rlr524/go-hw-web-app/pkg/template"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(rp *Repository) {
	Repo = rp
}

// Home is the / page handler. All web handler functions need to take in, as params,
// the ResponseWriter(w) method and a pointer to the Request(r) method
// All of the handlers also have a receiver via a pointer (m) to the repository and have access to the repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the /about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some business logic in which we define some data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// Send the data to the template
	template.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Madison is the /madison page handler
func (m *Repository) Madison(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["check"] = "Madison, are you behaving yourself?"

	template.RenderTemplate(w, "madison.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
