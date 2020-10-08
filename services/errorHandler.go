package services

import (
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func HandleError(db *sqlx.DB, service string, err error, message string, sev int){
	log.Printf("ERROR from %v:  %v %v", service, err, message)
	if sev>1 {
		models.CreateLog(db, service, err.Error(), message)
	}
}

