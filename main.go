package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var count = 10

	fmt.Println("*** START ***")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.gohtml")
		tmpl.ExecuteTemplate(w, "index.gohtml", count)
	})

	r.Post("/increase", func(w http.ResponseWriter, r *http.Request) {
		count += 1
		w.Write([]byte(fmt.Sprintf("%v", count)))
	})

	r.Post("/decrease", func(w http.ResponseWriter, r *http.Request) {
		count -= 1
		w.Write([]byte(fmt.Sprintf("%v", count)))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
