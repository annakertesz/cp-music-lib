package controller

import (
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func getAllArtist(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	artists, err := models.GetArtist(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	artistsJSON, err := json.Marshal(artistROListFromArtists(artists))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(artistsJSON)
	w.WriteHeader(http.StatusOK)
}

func getArtistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	param := chi.URLParam(r, "artistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)  //TODO: bad request to swagger
		return
	}
	artist, err := models.GetArtistByID(id, db)
	if artist == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	artistJSON, err := json.Marshal(artistROFromArtist(*artist))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(artistJSON)
	w.WriteHeader(http.StatusOK)
}

func artistROFromArtist(artist models.Artist) models.ArtistRO {
	return models.ArtistRO{
		ID:artist.ArtistID,
		Name:artist.ArtistName,
	}
}

func artistROListFromArtists(artists []models.Artist) []models.ArtistRO {
	var artistROs []models.ArtistRO
	for _, artist := range artists {
		artistROs = append(artistROs, artistROFromArtist(artist))
	}
	return artistROs
}
