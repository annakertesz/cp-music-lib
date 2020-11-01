package controller

import (
	"github.com/annakertesz/cp-music-lib/services/updater"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request, token string, coverFolder, musicFolder int){
	updater.Update(musicFolder, coverFolder, token, db)
}
