package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"net/http"
)

type e map[string]string

func Routes(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := models.Users(db)
		b, err := json.Marshal(users)

		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request){
		user, err := models.UnmarshalUser(r)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		created, err := models.CreateUser(db, user.Username)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		b, err := json.Marshal(created)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	return r
}
