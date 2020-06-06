package controller

import (
	"encoding/json"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
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
		http.Error(w, err.Error(), 404)
		return
	}
	_, err = models.CreateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param:= chi.URLParam(r, "userID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nuser id %v isnt a number")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.UpdateUserStatus(db, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func loginUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	user, err := models.UnmarshalUserValidation(r)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	userID := models.CheckUserCredentials(db, user.Username, user.Password)
	if (userID >0){
		uuid := uuid.New()
		err := models.CreateSession(db, userID, uuid.String())
		if err == nil {
			fmt.Fprint(w, uuid)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

