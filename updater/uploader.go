package updater

import (
	"bytes"
	"errors"
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	"strings"
)

func UploadSong(token string, fileBytes []byte, songBoxID int, db *sqlx.DB) error {
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
	artistID, err := artist.CreateArtist(db)
	if artistID == 0 || err != nil {
		fmt.Println("Couldnt create artist")
		if err != nil {
			return err
		}
		return errors.New("unexpected error while creating artist")
	}

	fmt.Printf("\ncreate album %v", metadata.Album())

	album := models.Album{
		AlbumName:   metadata.Album(),
		AlbumArtist: artistID,
	}
	albumID, createdNewAlbum, err := album.CreateAlbum(db)
	if albumID == 0 || err != nil {
		fmt.Println("Couldnt create album")
		if err != nil {
			return err
		}
		return errors.New("unexpected error while creating album")
	}
	if createdNewAlbum {
		if metadata.Picture() == nil {
			fmt.Printf("couldn't find image for the album %v", album.AlbumName)
		} else {
			boxID, err := box_lib.UploadFile(token, 110166546915, albumID, metadata.Picture().Data)
			if err != nil {
				fmt.Println("couldnt upload cover to box")
			}
			album.SaveAlbumImageID(db, boxID)
		}
	}
	fmt.Printf("\ncreate song %v", metadata.Title())
	song := models.NewSong(metadata.Title(), albumID, songBoxID)

	songID, _, err := song.CreateSong(db)
	if songID == 0 || err != nil {
		fmt.Println("Couldnt create song")
		if err != nil {
			return err
		}
		return errors.New("unexpected error while creating song")
	}
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
