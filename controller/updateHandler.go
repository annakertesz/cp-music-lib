package controller

import (
	"github.com/annakertesz/cp-music-lib/updater"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	token := r.URL.Query().Get("token")
	folder := r.URL.Query().Get("folderID")
	date := r.URL.Query().Get("date")
	folderID, err := strconv.Atoi(folder)
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = updater.Update(folderID, date, token, db)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
