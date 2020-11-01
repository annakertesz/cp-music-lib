package updater

import (
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func Update(songFolder int, coverFolder int, token string, db *sqlx.DB){
	log.Printf("start update at %v", time.Now())
	latestUpdate, err := models.GetLatestUpdate(db)
	if err != nil {
		services.HandleError(db, *err)
	}
	updateID, err := models.NewUpdate(db)
	if err != nil {
		services.HandleError(db, *err)
	}
	idList, err := box_lib.GetFileIDsToUpload(token, songFolder, latestUpdate)
	if err != nil {
		services.HandleError(db, *err)
	}
	foundSongs := len(idList)
	createdSongs := 0
	failedSongs := 0
	deletedSongs := 0
	for i := range idList {
		fileBytes, _, err := box_lib.DownloadFileBytes(token, idList[i])
		if err != nil {
			logID := services.HandleError(db, *err)
			failedSongs++
			models.SaveFailedSong(db, string(idList[i]), logID, updateID)
			continue
		}
		err = UploadSong(token, coverFolder, fileBytes, idList[i], db)
		if err != nil {
			logID := services.HandleError(db, *err)
			failedSongs++
			models.SaveFailedSong(db, string(idList[i]), logID, updateID)
			continue
		}
		createdSongs++
	}
	models.SaveUpdateNumbers(db, updateID, foundSongs, createdSongs, failedSongs, deletedSongs)
}
