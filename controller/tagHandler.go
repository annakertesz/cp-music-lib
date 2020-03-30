package controller

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getAllTag(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
