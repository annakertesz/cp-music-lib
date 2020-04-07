package models

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type Album struct {
	AlbumID     int    `json:"album_id" db:"id"`
	AlbumName   string `json:"album_name" db:"album_name"`
	AlbumArtist int    `json:"artist" db:"album_artist"`
}

func (album *Album) CreateAlbum(db *sqlx.DB) (int, bool) {
	if strings.Contains(album.AlbumName, "nstrumental") {
		if strings.Index(album.AlbumName, "(") > 0 {
			album.AlbumName = strings.TrimSpace(album.AlbumName[:strings.Index(album.AlbumName, "(")])
		}
	}
	var id int
	createdNew := false
	db.QueryRow(
		`SELECT id from album where album_name = $1`, album.AlbumName,
	).Scan(&id)
	if id == 0 {
		db.QueryRow(
			`INSERT INTO album (album_name, album_artist) VALUES ($1, $2) RETURNING id`,
			album.AlbumName, album.AlbumArtist,
		).Scan(&id)
		createdNew = true
	}

	return id, createdNew
}

func GetAlbum(db *sqlx.DB) ([]Album, error) {

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

func GetAlbumByID(id int, db *sqlx.DB) (*Album, error) {
	var album Album
	err := db.QueryRowx(
		`SELECT * FROM album WHERE id = $1`, id,
	).StructScan(&album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func GetAlbumByArtist(artistID int, db *sqlx.DB) ([]Album, error) {

	rows, err := db.Queryx(
		`SELECT * FROM album WHERE album_artist = $1`, artistID,
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
