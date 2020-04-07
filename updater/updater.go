package updater

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/jmoiron/sqlx"
	"os"
)

func Update(folder int, date string, token string, db *sqlx.DB) error {
	idList, err := box_lib.GetFileIDsToUpload(token, folder, date)
	if err != nil {
		return err
	}
	fmt.Println("progressBar")
	for i := range idList {
		fmt.Println(idList[i])
		err := box_lib.DownloadFile(token, idList[i])
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		file, err := os.Create(fmt.Sprintf("../sources/music/%v.mp3", idList[i]))
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		err = UploadSong(file, idList[i], db)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		err = os.Remove(fmt.Sprintf("../sources/music/%v.mp3", idList[i]))
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
