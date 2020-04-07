package controller

import (
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getUsers(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	users, err := models.GetUsers(db)
	b, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func createUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
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
}
