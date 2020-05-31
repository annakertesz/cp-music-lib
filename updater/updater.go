package updater

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
)

func Update(songFolder int, coverFolder int, token string, db *sqlx.DB) error {
	latestUpdate, err := models.GetLatestUpdate(db)
	if err != nil {
		return err
	}
	updateID, err := models.NewUpdate(db)
	if err != nil {
		return err
	}
	idList, err := box_lib.GetFileIDsToUpload(token, songFolder, latestUpdate)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	foundSongs := len(idList)
	createdSongs := 0
	failedSongs := 0
	deletedSongs := 0
	for i := range idList {
		fileBytes, _, err := box_lib.DownloadFileBytes(token, idList[i])
		if err != nil {
			fmt.Println(err.Error())
			failedSongs++
			models.SaveFailedSong(db, string(idList[i]), err.Error(), updateID)
			continue
		}
		err = UploadSong(token, coverFolder, fileBytes, idList[i], db)
		if err != nil {
			fmt.Println(err.Error())
			failedSongs++
			models.SaveFailedSong(db, string(idList[i]), err.Error(), updateID)
			continue
		}
		createdSongs++
	}
	models.SaveUpdateNumbers(db, updateID, foundSongs, createdSongs, failedSongs, deletedSongs)
	return nil
}
