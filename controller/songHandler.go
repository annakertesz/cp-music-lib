package controller

import (
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getSongs(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	r.URL.Query().Get("id")
	users, err := models.GetSongs(db)
	b, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
