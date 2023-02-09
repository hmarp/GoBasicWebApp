package render

import (
	"fmt"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders a specified template
func RenderTemplate(w http.ResponseWriter, t string) {
	_, inMap := templateCache[t]
	
	if !inMap {
		fmt.Println("Creating template and adding to cache")
		err := createTemplateCache(t)

		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Using cached template")
	}

	tmpl := templateCache[t]
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	templateCache[t] = tmpl

	return nil
}
