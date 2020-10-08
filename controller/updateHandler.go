package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/services/updater"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request, token string, coverFolder, musicFolder int){
	err := updater.Update(musicFolder, coverFolder, token, db)
	if err!= nil {
		fmt.Println("database update was unsuccessful")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
