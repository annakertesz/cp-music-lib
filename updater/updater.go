package updater

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/jmoiron/sqlx"
)

func Update(folder int, date string, token string, db *sqlx.DB) error {
	idList, err := box_lib.GetFileIDsToUpload(token, folder, date)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Start upload the list items.")
	for i := range idList {
		fmt.Printf("\n%v\n",idList[i])
		fileBytes, _, err := box_lib.DownloadFileBytes(token, idList[i])
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//file, err := os.Open(fmt.Sprintf("sources/music/%v.mp3", idList[i]))
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return err
		//}
		err = UploadSong(token, fileBytes, idList[i], db)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	fmt.Println("done")
	return nil
}
