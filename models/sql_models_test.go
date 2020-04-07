package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

func TestArtist_CreateArtist(t *testing.T) {
	db, err := connectToDB()
		if err!=nil {
			panic(err.Error())
		}
	artist := Artist{
		ArtistName: "name",
	}
	artist.CreateArtist(db)
}

//func TestAlbum_CreateAlbum(t *testing.T) {
//	db, err := connectToDB()
//	if err!=nil {
//		fmt.Println(err)
//	}
//	artist := Artist{
//		ArtistName: "name",
//	}
//	album := Album{
//		AlbumName:              "album_name",
//		AlbumArtist:            &artist,
//		AlbumCoverUrl:          "coverbigurl",
//		AlbumCoverThumbnailUrl: "coversmallurl",
//	}
//	fmt.Println(db)
//	artist.CreateArtist(db)
//	album.CreateAlbum(db)
//}

func connectToDB()(*sqlx.DB, error){
	url := "host=localhost port=5432 user=anna password=anna dbname=centralp sslmode=disable"
	db, err := sqlx.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}