package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/updater"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func update(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	fmt.Println("Update:")
	token := r.URL.Query().Get("token")
	folder := r.URL.Query().Get("folderID")
	date := r.URL.Query().Get("date")
	fmt.Printf("\nfolder id: %v   date: %v")
	folderID, err := strconv.Atoi(folder)
	if err!= nil {
		fmt.Println("Need numeric folder id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = updater.Update(folderID, date, token, db)
	if err!= nil {
		fmt.Println("database update was unsuccessful")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
