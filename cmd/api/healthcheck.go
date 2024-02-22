package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
	fmt.Fprintln(w, "environment: %s\n", app.config.env)
	fmt.Fprintln(w, "version: %s\n", app.config.version)	
}