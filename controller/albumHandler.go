package controller

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getAllAlbum(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}


func getAlbumsById(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}


func getAlbumsByArtist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}