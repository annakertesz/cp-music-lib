package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

func TestGetSongByTitleAndAlbum(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		fmt.Println(err)
	}
	songs, err := GetSongByTitleAndAlbum("Hate Love Hate (Instrumental)", "Hate Love Hate (Instrumental)", db)
	fmt.Print(songs)
}

func TestAlbum_CreateAlbum(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		fmt.Println(err)
	}
	artist := Artist{
		ArtistName: "name",
	}
	album := Album{
		AlbumName:              "album_name",
		AlbumArtist:            &artist,
		AlbumCoverUrl:          "coverbigurl",
		AlbumCoverThumbnailUrl: "coversmallurl",
	}
	fmt.Println(db)
	artist.CreateArtist(db)
	album.CreateAlbum(db)
}

func connectToDB()(*sqlx.DB, error){
	url := "host=localhost port=5432 user=postgres password= dbname=postgres sslmode=disable"
	db, err := sqlx.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}