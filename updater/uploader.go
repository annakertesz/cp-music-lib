package updater

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

func UploadSong(file *os.File, songID string, db *sqlx.DB) {
	metadata, _ := tag.ReadFrom(file)

	artist := models.Artist{
		ArtistName: metadata.Artist(),
	}
	err := artist.CreateArtist(db)
	if err != nil {
		fmt.Print("artist: ")
		fmt.Printf(err.Error())
	}
	album := models.Album{
		AlbumID:                0,
		AlbumName:              metadata.Title(),
		AlbumArtist:            &artist,
		AlbumCoverUrl:          songID,
		AlbumCoverThumbnailUrl: songID,
	}
	err = album.CreateAlbum(db)
	if err != nil {
		fmt.Print("album: ")
		fmt.Printf(err.Error())
	}
	song := models.Song{
		SongName:              metadata.Title(),
		SongLqURL:             songID,
		SongHqURL:             songID,
		SongInstrumentalLqURL: songID,
		SongInstrumentalHqURL: songID,
		SongAlbum:             &album,
		SongArtist:            &artist,
		SongTags:              nil,
	}
	err = song.CreateSong(db)
	if err != nil {
		fmt.Print("song: ")
		fmt.Printf(err.Error())
	}
	tags := strings.Split(metadata.Genre(), "/")
	for _, tag := range tags {
		tag := strings.Trim(tag, " ")
		tagObj := models.Tag{TagName: strings.Trim(tag, " ")}
		err = tagObj.CreateTag(db)
		if err != nil {
			fmt.Print("tag: ")
			fmt.Printf(err.Error())
		}
		tagSong := models.TagSong{
			Tag:  &tagObj,
			Song: &song,
		}
		err = tagSong.CreateTagSong(db)
		if err != nil {
			fmt.Print("tag-song: ")
			fmt.Printf(err.Error())
		}
	}
}

//func checkIfSongExists(title string, album string) bool {
//
//}
