package controller

import (
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func createPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	playlist, err := models.UnmarshalPlaylist(r)
	if err != nil || playlist == nil {
		http.Error(w, err.Error(), 404)
		return
	}
	id, err := models.CreatePlaylist(db, *playlist)
	if err != nil || id == 0 {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getPlaylistByUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getPlaylistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
