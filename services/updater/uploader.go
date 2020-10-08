package updater

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
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
			image, _, err := image.Decode(bytes.NewReader(metadata.Picture().Data))
			newImage := resize.Resize(160, 160, image, resize.Lanczos3)
			b := make([]byte, 0, 1024)
			buf := bytes.NewBuffer(b)

			err = jpeg.Encode(buf, newImage, nil)
			boxID, err := box_lib.UploadFile(token, coverFolder, albumID, b)

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
