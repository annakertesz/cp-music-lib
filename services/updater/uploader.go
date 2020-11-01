package updater

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"strings"
)

func UploadSong(token string, coverFolder int, fileBytes []byte, songBoxID int, db *sqlx.DB) *models.ErrorModel {
	reader := bytes.NewReader(fileBytes)
	metadata, err := tag.ReadFrom(reader)
	if err != nil {
		return &models.ErrorModel{
			Service: "Updater",
			Err:     err,
			Message: fmt.Sprintf("error while trying to read metadata from item (boxID = %v)", songBoxID),
			Sev:     3,
		}
	}
	artist := models.Artist{
		ArtistName: metadata.Artist(),
	}
	artistID, err := artist.CreateArtist(db)
	if artistID == 0 || err != nil {
		if err != nil {
			return &models.ErrorModel{
				Service: "Updater",
				Err:     err,
				Message: fmt.Sprintf("error while trying to create new artist (boxID = %v, artist = %v)", songBoxID, metadata.Artist()),
				Sev:     3,
			}
		}
		return &models.ErrorModel{
			Service: "Updater",
			Err:     errors.New("unexpected error while creating album"),
			Message: fmt.Sprintf("error while trying to create new album (boxID = %v, album = %v)", songBoxID, metadata.Album()),
			Sev:     3,
		}
	}
	album := models.Album{
		AlbumName:   metadata.Album(),
		AlbumArtist: artistID,
	}
	albumID, createdNewAlbum, err := album.CreateAlbum(db)
	if albumID == 0 || err != nil {
		if err != nil {
			return &models.ErrorModel{
				Service: "Uploader",
				Err:     err,
				Message: fmt.Sprintf("error while trying to create new album (boxID = %v, album = %v)", songBoxID, metadata.Album()),
				Sev:     3,
			}
		}
		return &models.ErrorModel{
			Service: "Uloader",
			Err:     errors.New("unexpected error while creating album"),
			Message: fmt.Sprintf("error while trying to create new album (boxID = %v, album = %v)", songBoxID, metadata.Album()),
			Sev:     3,
		}
	}
	if createdNewAlbum {
		if metadata.Picture() == nil {
			services.HandleError(db, models.ErrorModel{
				Service: "Uploader",
				Err:     errors.New("missing picture from metadata"),
				Message: fmt.Sprintf("there is no picture for item (boxID = %v", songBoxID),
				Sev:     3,
			})
		} else {
			image, _, err := image.Decode(bytes.NewReader(metadata.Picture().Data))
			if err != nil {
				services.HandleError(db, models.ErrorModel{
					Service: "Uploader",
					Err:     err,
					Message: fmt.Sprintf("error while decode image from metadata for item (boxID = %v", songBoxID),
					Sev:     3,
				})
			} else {
				newImage := resize.Resize(160, 160, image, resize.Lanczos3)
				b := make([]byte, 0, 1024)
				buf := bytes.NewBuffer(b)

				err = jpeg.Encode(buf, newImage, nil)
				if err != nil {
					services.HandleError(db, models.ErrorModel{
						Service: "Uploader",
						Err:     err,
						Message: fmt.Sprintf("error while encode resized image for item (boxID = %v", songBoxID),
						Sev:     3,
					})
				} else {
					boxID, err := box_lib.UploadFile(token, coverFolder, albumID, buf.Bytes())

					if err != nil {
						services.HandleError(db, models.ErrorModel{
							Service: "Uploader",
							Err:     err,
							Message: fmt.Sprintf("Couldnt upload image to box (boxID = %v", songBoxID),
							Sev:     3,
						})
					}
					album.SaveAlbumImageID(db, boxID)
				}
			}
		}
	}
	song := models.NewSong(metadata.Title(), albumID, songBoxID)

	songID, _, err := song.CreateSong(db)
	if songID == 0 || err != nil {
		if err != nil {
			return &models.ErrorModel{
				Service: "Uploader",
				Err:     err,
				Message: fmt.Sprintf("error while trying to create new song (boxID = %v, title = %v)", songBoxID, metadata.Title()),
				Sev:     3,
			}
		}
		return &models.ErrorModel{
			Service: "Uloader",
			Err:     errors.New("unexpected error while creating song"),
			Message: fmt.Sprintf("error while trying to create new song (boxID = %v, title = %v)", songBoxID, metadata.Title()),
			Sev:     3,
		}
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
