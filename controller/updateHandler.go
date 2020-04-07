package controller

import (
	"github.com/annakertesz/cp-music-lib/updater"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	token := chi.URLParam(r, "token")
	folder := chi.URLParam(r, "folderID")
	date := chi.URLParam(r, "date")
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
