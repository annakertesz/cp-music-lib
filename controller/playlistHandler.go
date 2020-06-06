package controller

import (
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func createPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	title := r.URL.Query().Get("title")
	if len(title)<1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
	if err!= nil {
		http.Error(w, err.Error(), 500)
		return
	}
	id, err := models.CreatePlaylist(db, userID, title)
	if err != nil || id == 0 {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
}
//
//func getAllPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
//	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
//	if err!= nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//}
//
//func getPlaylistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
//	w.WriteHeader(http.StatusOK)
//}

//func playlistROFromPlaylist(playlist models.Playlist,  db *sqlx.DB) PlaylistRO{
//	//user, err := models.GetUserByID(db, playlist.User)
//	//songROList, err := songROListFromSongs(songs, db)
//	//return PlaylistRO{
//	//	Title: playlist.Title,
//	//	User:  models.UserROFromUser(user),
//	//	Songs: nil,
//	//}
//}
//
//func playlistROListFromPlaylists(songs []models.Song, db *sqlx.DB) ([]SongRO, error) {
//
//}error