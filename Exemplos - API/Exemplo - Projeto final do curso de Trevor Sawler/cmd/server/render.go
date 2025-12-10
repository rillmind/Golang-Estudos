package main

import (
	"fmt"
	"net/http"
	"time"
)

var pathToTemplate = ".cmd/server/templates/"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticaded int
	Now           time.Time
	// User *data.User
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplate),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplate),
	}

	templateSlice := []string{}
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplate, t))

	// TODO
}
