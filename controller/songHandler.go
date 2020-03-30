package controller

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getSongByID(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getSongByAlbum(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getSongByArtist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getSongByTag(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func searchSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

//func getSongs(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
//	r.URL.Query().Get("id")
//	users, err := models.GetSongs(db)
//	b, err := json.Marshal(users)
//
//	if err != nil {
//		http.Error(w, err.Error(), 422)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Write(b)
//}
