package controller

import (
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func getSongByID(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param:= chi.URLParam(r, "songID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	song, err := models.GetSongByID(id, db)
	if song == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	album, err := models.GetAlbumByID(song.SongAlbum, db)
	if album == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	artist, err := models.GetArtistByID(album.AlbumArtist, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	songJSON, err := json.Marshal(songROFromSong(*song, *album, *artist))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(songJSON)
	w.WriteHeader(http.StatusOK)
}

func getSongByAlbum(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param := r.URL.Query().Get("albumID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	songs, err := models.GetSongByAlbum(id, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	songsJSON, err := json.Marshal(SongROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
	w.WriteHeader(http.StatusOK)
}

func getSongByArtist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param := r.URL.Query().Get("artistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	songs, err := models.GetSongByArtist(id, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	songsJSON, err := json.Marshal(SongROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
	w.WriteHeader(http.StatusOK)
}

func getSongByTag(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param := r.URL.Query().Get("tagID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	songs, err := models.GetSongByTag(id, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	songsJSON, err := json.Marshal(SongROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
	w.WriteHeader(http.StatusOK)
}

func searchSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func songROFromSong(song models.Song, album models.Album, artist models.Artist) SongRO{
	return SongRO{
		ID:      song.SongID,
		Title:   song.SongName,
		Album:   albumROFromAlbum(album, artist),
		LqSong:  song.SongLqURL,
		HqSong:  song.SongHqURL,
		LqInstr: song.SongInstrumentalLqURL,
		HqInstr: song.SongInstrumentalHqURL,
	}
}

func songROListFromSongs(songs []models.Song, db *sqlx.DB) ([]SongRO, error) {
	songROs := make([]SongRO, 0)
	for _, song := range songs {
		album, err := models.GetAlbumByID(song.SongAlbum, db)
		artist, err := models.GetArtistByID(album.AlbumArtist, db)
		if err != nil {
			return nil, err
		}
		songROs = append(songROs, songROFromSong(song, *album, *artist))
	}
	return songROs, nil
}
