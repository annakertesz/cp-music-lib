package models

import (
	"github.com/jmoiron/sqlx"
)

type Album struct {
	AlbumID int `json:"album_id" db:"id"`
	AlbumName string `json:"album_name" db:"album_name"`
	AlbumArtist *Artist `json:"artist" db:"album_artist"`
	AlbumCoverUrl string `json:"cover_url" db:"cover_url"`
	AlbumCoverThumbnailUrl string `json:"cover_thumbnail_url" db:"cover_thumbnail_url"`
}

func (album *Album) CreateAlbum(db *sqlx.DB) error {

	row := db.QueryRow(
		`INSERT INTO album (album_name, album_artist, cover_url, cover_thumbnail_url) VALUES ($1, $2, $3, $4) RETURNING id`,
		album.AlbumName, album.AlbumArtist.ArtistID, album.AlbumCoverUrl, album.AlbumCoverThumbnailUrl,
	)

	err := row.Scan(&album.AlbumID)

	if err != nil {
		return err
	}

	return nil
}

//func GetAlbum(db *sqlx.DB, w http.ResponseWriter, r *http.Request) ([]Album, error){
//
//	rows, err := db.Queryx(
//		`SELECT * FROM album WHERE id = $1` , id,
//	)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		var album_name string
//		var album_artist int
//		var cover_url string
//		var cover_thumbnail_url string
//		rows.Scan(&id, &album_name, &album_artist, &cover_thumbnail_url, &cover_url)
//		artist, err := GetArtistByID(album_artist, db)
//		if err != nil {
//			return nil, err
//		}
//		return Album{
//			AlbumID:                id,
//			AlbumName:              album_name,
//			AlbumArtist:            &artist,
//			AlbumCoverUrl:          cover_url,
//			AlbumCoverThumbnailUrl: cover_thumbnail_url,
//		}, nil
//	}
//	return Album{}, errors.New("there is no album with this id")
//}