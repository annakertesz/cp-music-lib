package controller

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

func createPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){

}

func getPlaylistByUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getPlaylistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
