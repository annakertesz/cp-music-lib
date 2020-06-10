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

func getAllSongs(db *sqlx.DB, w http.ResponseWriter, r *http.Request){

	songs, err := models.GetAllSongs(db)
	if err != nil {
		fmt.Printf("error in getAllSongs: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		fmt.Printf("error in getAllSongs: %v", err.Error())
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

func getSongByID(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param:= chi.URLParam(r, "songID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number", param)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Sprintf("id=%v (string: %v", id, param)
	song, err := models.GetSongByID(id, db)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	if song == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	album, err := models.GetAlbumByID(song.SongAlbum, db)
	if album == nil {
		fmt.Sprintf("\n couldnt find album %v for song %v", song.SongAlbum, song.SongID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	artist, err := models.GetArtistByID(album.AlbumArtist, db)
	if err != nil {
		fmt.Sprintf("\n couldnt find artist %v for album %v, for song %v", album.AlbumArtist, album.AlbumID, song.SongID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tags, err := models.GetTagsOfSong(db,song.SongID)
	if err != nil {
		fmt.Sprintf("\n couldnt find tag for song %v",  song.SongID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	songJSON, err := json.Marshal(songROFromSong(*song, *album, *artist, tags))
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
	fmt.Println(param)
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number", id)
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	fmt.Println(id)
	songs, err := models.GetSongByAlbum(id, db)
	if err != nil {
		fmt.Printf("error in getSongByAlbum: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		fmt.Printf("error in createSongRO: %v", err.Error())
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
		fmt.Printf("\nartist id %v isnt a number", id)
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	songs, err := models.GetSongByArtist(id, db)
	if err != nil {
		fmt.Printf("error in getSongByArtist: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		fmt.Printf("error in createSongRO: %v", err.Error())
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
	fmt.Println("1")
	param := r.URL.Query().Get("tagID")
	fmt.Println("2")
	id, err := strconv.Atoi(param)
	fmt.Println("3")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	fmt.Println("4")
	songs, err := models.GetSongByTag(id, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("5")
	SongROs, err := songROListFromSongs(songs, db)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("6")
	songsJSON, err := json.Marshal(SongROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("7")
	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
	w.WriteHeader(http.StatusOK)
}

func searchSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	getSongByTag(db, w, r)
}

func songROFromSong(song models.Song, album models.Album, artist models.Artist, tags []models.Tag) models.SongRO {
	return models.SongRO{
		ID:      song.SongID,
		Title:   song.SongName,
		Album:   albumROFromAlbum(album, artist),
		LqSong:  song.SongLqURL,
		HqSong:  song.SongHqURL,
		LqInstr: song.SongInstrumentalLqURL,
		HqInstr: song.SongInstrumentalHqURL,
		Tags:tagROListFromTag(tags),
	}
}

func songROListFromSongs(songs []models.Song, db *sqlx.DB) ([]models.SongRO, error) {
	songROs := make([]models.SongRO, 0)
	for _, song := range songs {
		album, err := models.GetAlbumByID(song.SongAlbum, db)
		if album == nil {
			fmt.Sprintf("\n couldnt find album %v for song %v", song.SongAlbum, song.SongID)
			return nil, err
		}
		artist, err := models.GetArtistByID(album.AlbumArtist, db)
		if err != nil {
			fmt.Sprintf("\n couldnt find artist %v for album %v, for song %v", album.AlbumArtist, album.AlbumID, song.SongID)
			return nil, err
		}
		tags, err := models.GetTagsOfSong(db, song.SongID)
		if err != nil {
			fmt.Sprintf("\n couldnt find tag for song %v",  song.SongID)
			return nil, err
		}
		songROs = append(songROs, songROFromSong(song, *album, *artist,  tags))
	}
	return songROs, nil
}
