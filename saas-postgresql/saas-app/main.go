package main

import (
	"net/http"

	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		HomePage().Render(context.Background(), w)
	})

	r.Get("/databases", func(w http.ResponseWriter, r *http.Request) {
		dbs := []Database{
			{name: "ranzer"},
			{name: "idh_plot"},
			{name: "twicovers"},
		}

		ListPage(dbs).Render(context.Background(), w)
	})

	r.Get("/databases", func(w http.ResponseWriter, r *http.Request) {
		dbs := []Database{
			{name: "ranzer"},
			{name: "idh_plot"},
			{name: "twicovers"},
		}

		ListPage(dbs).Render(context.Background(), w)
	})

	r.Get("/databases/{database}", func(w http.ResponseWriter, r *http.Request) {
		dbs := []Database{
			{name: "ranzer"},
			{name: "idh_plot"},
			{name: "twicovers"},
		}

		ListPage(dbs).Render(context.Background(), w)
	})

	http.ListenAndServe(":3000", r)
}

type Database struct {
	name string
}
