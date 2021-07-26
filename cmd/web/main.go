/*
Package main provides the application entry point. It spins up a server
and fires off all route handlers.
*/
package main

import (
	"fmt"
	"github.com/rlr524/go-hw-web-app/pkg/config"
	"github.com/rlr524/go-hw-web-app/pkg/handlers"
	"github.com/rlr524/go-hw-web-app/pkg/routes"
	"github.com/rlr524/go-hw-web-app/pkg/template"
	"log"
	"net/http"
)

// In Go, the colon is required when passing the port number in a variable
const portNumber = ":3030"

// main is the application entry point
func main() {
	// Instead of putting our page data into the main() function, here
	// we are still calling the HandleFunc function from the http package but
	// instead we're passing in our page handlers as closures and these
	// functions become our routes
	var app config.AppConfig
	tc, err := template.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// Create a repository, assign it to repo and pass it back to the NewHandlers function
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Calls the NewTemplates function from the template package and references the AppConfig
	// Need to reference app because we are setting app as a pointer to config.AppConfig above
	// in that it is a var named "app" that is of type config.AppConfig
	template.NewTemplates(&app)

	// Create a server using our routes and pat router in the routes package
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(&app),
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	err = srv.ListenAndServe()
	log.Fatal(err)
}
