package controller

import (
	"encoding/json"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func createPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if len(title) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
	if err != nil {
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

func deletePlaylistByID(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param:= chi.URLParam(r, "playlistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nplaylist id %v isnt a number")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.DeletePlaylist(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func addSongToPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param:= chi.URLParam(r, "songID")
	songID, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	plID := r.URL.Query().Get("playlistID")
	playlistID, err := strconv.Atoi(plID)
	if err != nil {
		fmt.Printf("playlist id %v isnt a number")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.AddSongToPlalist(db, songID, playlistID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getAllPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	playlists, err := models.GetAllPLaylist(db, userID)
	playlistROs, err := playlistROListFromPlaylists(playlists, db)
	playlistJSON, err := json.Marshal(playlistROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlistJSON)
	w.WriteHeader(http.StatusOK)
}

func getPlaylistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param:= chi.URLParam(r, "playlistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	playlist, err := models.GetPlaylistByID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	playlistRO, err := playlistROFromPlaylist(playlist, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	playlistJSON, err := json.Marshal(playlistRO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlistJSON)
	w.WriteHeader(http.StatusOK)
}

func playlistROFromPlaylist(playlist models.Playlist, db *sqlx.DB) (*PlaylistRO, error) {
	user, err := models.GetUserByID(db, playlist.User)
	if err != nil {
		return nil, err
	}
	songs, err := models.GetSongsByPlaylist(db, playlist.ID)
	if err != nil {
		return nil, err
	}
	songROList, err := songROListFromSongs(songs, db)
	if err != nil {
		return nil, err
	}
	return &PlaylistRO{
		Title: playlist.Title,
		User:  UserROFromUser(user),
		Songs: songROList,
	}, nil
}

func playlistROListFromPlaylists(playlists []models.Playlist, db *sqlx.DB) ([]PlaylistRO, error) {
	playlistROs := make([]PlaylistRO, 0)
	for _, playlist := range playlists {
		pl, err := playlistROFromPlaylist(playlist, db)
		if err != nil {
			return nil, err
		}
		playlistROs = append(playlistROs, *pl)
	}
	return playlistROs, nil
}
