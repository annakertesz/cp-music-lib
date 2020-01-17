package controller

import (
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
)

type e map[string]string

func Routes(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getUsers(db, w, r)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request){
		createUser(db, w, r)
	})

	return r
}
