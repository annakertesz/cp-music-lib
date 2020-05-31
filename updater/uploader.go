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

func UploadSong(token string, coverFolder int, fileBytes []byte, songBoxID int, db *sqlx.DB) error {
	reader := bytes.NewReader(fileBytes)
	metadata, err := tag.ReadFrom(reader)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	artist := models.Artist{
		ArtistName: metadata.Artist(),
	}
	artistID, err := artist.CreateArtist(db)
	if artistID == 0 || err != nil {
		fmt.Println("Couldnt create artist")
		if err != nil {
			return err
		}
		return errors.New("unexpected error while creating artist")
	}
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
			boxID, err := box_lib.UploadFile(token, coverFolder, albumID, metadata.Picture().Data)
			if err != nil {
				fmt.Println("couldnt upload cover to box")
			}
			album.SaveAlbumImageID(db, boxID)
		}
	}
	song := models.NewSong(metadata.Title(), albumID, songBoxID)

	songID, _, err := song.CreateSong(db)
	if songID == 0 || err != nil {
		fmt.Println("Couldnt create song")
		if err != nil {
			return err
		}
		return errors.New("unexpected error while creating song")
	}
	tags := strings.Split(metadata.Genre(), "/")
	for _, tag := range tags {
		tagObj := models.Tag{TagName: strings.TrimSpace(tag)}
		tagID, _ := tagObj.CreateTag(db)
		if tagID == 0 {
			fmt.Println("failed to save tag")
		}
		tagSong := models.TagSong{
			Tag:  tagID,
			Song: songID,
		}
		tagSongId := tagSong.CreateTagSong(db)
		if tagSongId == 0 {
			fmt.Println("failed to save tag_song")
		}
	}
	return nil
}

//func checkIfSongExists(title string, album string) bool {
//
//}
