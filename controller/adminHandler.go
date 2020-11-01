package controller

import (
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services/updater"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request, token string, coverFolder, musicFolder int){
	updater.Update(musicFolder, coverFolder, token, db)
}

func cleardb(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	log.Printf("clear db")
		err := models.ClearUpdates(db)
		err = models.ClearSong(db)
		err = models.ClearPlaylist(db)
		err = models.ClearTagSong(db)
		err = models.ClearTag(db)
		err = models.ClearAlbum(db)
		err = models.ClearArtist(db)
		err = models.ClearFailedSongs(db)
		err = models.ClearLogs(db)
		if err != nil {
			w.WriteHeader(500)
		}
}
