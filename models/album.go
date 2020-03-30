package models

import (
	"github.com/jmoiron/sqlx"
)

type Album struct {
	AlbumID int `json:"album_id" db:"id"`
	AlbumName string `json:"album_name" db:"album_name"`
	AlbumArtist int `json:"artist" db:"album_artist"`
	AlbumCoverUrl string `json:"cover_url" db:"cover_url"`
	AlbumCoverThumbnailUrl string `json:"cover_thumbnail_url" db:"cover_thumbnail_url"`
}

func (album *Album) CreateAlbum(db *sqlx.DB) error {

	row := db.QueryRow(
		`INSERT INTO album (album_name, album_artist, cover_url, cover_thumbnail_url) VALUES ($1, $2, $3, $4) RETURNING id`,
		album.AlbumName, album.AlbumArtist, album.AlbumCoverUrl, album.AlbumCoverThumbnailUrl,
	)

	err := row.Scan(&album.AlbumID)

	if err != nil {
		return err
	}

	return nil
}

func GetAlbum(db *sqlx.DB) ([]Album, error){

	rows, err := db.Queryx(
		`SELECT * FROM album`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albums := make([]Album, 0)
	for rows.Next() {
		var album Album
		rows.StructScan(&album)
		albums = append(albums, album)
	}
	return albums, nil
}

func GetAlbumByID(id int, db *sqlx.DB) (*Album, error){
	var album Album
	err := db.QueryRowx(
		`SELECT * FROM album WHERE id = $1` , id,
	).StructScan(&album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func GetAlbumByArtist(artistID int, db *sqlx.DB) ([]Album, error){

	rows, err := db.Queryx(
		`SELECT * FROM album WHERE album_artist = $1` , artistID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albums := make([]Album, 0)
	for rows.Next() {
		var album Album
		rows.StructScan(&album)
		albums = append(albums, album)
	}
	return albums, nil
}