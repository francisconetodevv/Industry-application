package main

import (
	"fmt"
	"html/template"
	"industrialApplication/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func renderTemplate(templateName string, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	// This lib is responsible to create the routers that allows to navigate and sent the data
	router := mux.NewRouter()

	/*
		The next step is creat a template.ExecuteTemplate() to render the webPage to connect with the routers.
		If I pass *.html, the code will read all the files that use the extension .html
	*/

	templates := template.Must(template.ParseGlob("static/html/*.html"))

	/*

		I have transformed this func in the next code - renderTemplate

		router.HandleFunc("/CreateMachine", func(w http.ResponseWriter, r *http.Request) {
			err := templates.ExecuteTemplate(w, "CreateMachine.html", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}).Methods(http.MethodGet)

		router.HandleFunc("/CreateMachine", server.CreateMachine).Methods(http.MethodPost)

	*/

	// Routers
	router.HandleFunc("/CreateMachine", server.CreateMachine).Methods(http.MethodGet)
	router.HandleFunc("/UpdateMachine", renderTemplate("UpdateMachine.html", templates)).Methods(http.MethodPut)

	/*
	   log.Fatal is a function from the log package that prints a log message
	   and terminates the program with exit status 1 (indicating an error). When used
	   with http.ListenAndServe, it is commonly applied to:

	   1. Start the HTTP server on the specified port (5000 in this case)
	   2. Capture critical initialization errors, such as:
	      - Port already in use
	      - Permission failures
	      - Network issues

	   If http.ListenAndServe returns an error, log.Fatal immediately
	   halts execution. Otherwise, the server runs indefinitely, and this line
	   is only reached in case of a subsequent error.
	*/
	fmt.Println("Escutando na rota 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
