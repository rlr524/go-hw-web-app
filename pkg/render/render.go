/*
Package render implements a simple means by which to render templates.
*/
package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// UseTemplate renders templates using html/template
func UseTemplate(w http.ResponseWriter, tmpl string) {
	_, err := UseTemplateTest(w)
	if err != nil {
		_ = fmt.Errorf("error getting template cache: #{err}")
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		_ = fmt.Errorf("error parsing template: %v", err)
		return
	}
}

func UseTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)

		// Create the template set (ts) from the page templates (*.page.gohtml)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("/templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
