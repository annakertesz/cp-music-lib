package updater

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/jmoiron/sqlx"
	"os"
)

func Update(folder int, date string, token string, db *sqlx.DB) {
	idList, _ := box_lib.GetFileIDsToUpload(token, folder, date)
	fmt.Println("progressBar")
	for i := range idList {
			fmt.Println(idList[i])
			box_lib.DownloadFile(token, idList[i])
			file, _ := os.Open(fmt.Sprintf("../sources/music/%v.mp3", idList[i]))
			UploadSong(file, idList[i], db)
			os.Remove(fmt.Sprintf("../sources/music/%v.mp3", idList[i]))

	}
}
