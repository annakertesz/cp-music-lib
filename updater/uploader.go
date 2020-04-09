package updater

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

func UploadSong(fileBytes []byte, songBoxID int, db *sqlx.DB) error {
	reader := bytes.NewReader(fileBytes)
	fmt.Println("upload song")
	fmt.Print("read metadata from file")
	metadata, err := tag.ReadFrom(reader)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	artist := models.Artist{
		ArtistName: metadata.Artist(),
	}
	fmt.Printf("\ncreate artist %v", metadata.Artist())
	artistID := artist.CreateArtist(db)
	if artistID == 0 {
		panic("artist")
	}

	fmt.Printf("\ncreate album %v", metadata.Album())

	album := models.Album{
		AlbumName:   metadata.Album(),
		AlbumArtist: artistID,
	}
	albumID, createdNewAlbum := album.CreateAlbum(db)
	if albumID == 0 {
		panic("album")
	}
	if createdNewAlbum {
		out, err := os.Create(fmt.Sprintf("../sources/imgs/%v.jpg", albumID))
		if metadata.Picture() == nil {
			fmt.Printf("couldn't find image for the album %v", album.AlbumName)
		} else {

			err = binary.Write(out, binary.BigEndian, metadata.Picture().Data)
			if err != nil {
				return err
			}
		}
	}
	if albumID == 0 {
		panic("album")
	}
	fmt.Printf("\ncreate song %v", metadata.Title())
	song := models.NewSong(metadata.Title(), albumID, songBoxID)

	songID, _ := song.CreateSong(db)
	fmt.Printf("\ncreate tags %v", metadata.Genre())
	tags := strings.Split(metadata.Genre(), "/")
	for _, tag := range tags {
		tagObj := models.Tag{TagName: strings.TrimSpace(tag)}
		tagID, _ := tagObj.CreateTag(db)
		tagSong := models.TagSong{
			Tag:  tagID,
			Song: songID,
		}
		tagSong.CreateTagSong(db)
	}
	return nil
}

//func checkIfSongExists(title string, album string) bool {
//
//}
