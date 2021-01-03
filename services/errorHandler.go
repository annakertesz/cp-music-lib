package services

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func HandleError(db *sqlx.DB, err models.ErrorModel) int{
	log.Printf("ERROR from %v:  %v %v", err.Service, err.Err, err.Message)
	fmt.Printf("ERROR from %v:  %v %v", err.Service, err.Err, err.Message)
	if err.Sev>1 {
		id, _ := models.CreateLog(db, err.Service, err.Err.Error(), err.Message)
		return id
	}
	return 0
}

