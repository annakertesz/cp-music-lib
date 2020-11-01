package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

const default_id  = 712601893142

type Album struct {
	AlbumID     int    `json:"album_id" db:"id"`
	AlbumName   string `json:"album_name" db:"album_name"`
	AlbumArtist int    `json:"artist" db:"album_artist"`
	AlbumCover *string `json:"album_cover" db:"album_cover"`
}

func (album *Album) CreateAlbum(db *sqlx.DB) (int, bool, error) {
	if strings.Contains(album.AlbumName, "nstrumental") {
		if strings.Index(album.AlbumName, "(") > 0 {
			album.AlbumName = strings.TrimSpace(album.AlbumName[:strings.Index(album.AlbumName, "(")])
		}
	}
	var id int
	createdNew := false
	err := db.QueryRow(
		`SELECT id from album where UPPER(album_name) = UPPER($1)`, album.AlbumName,
	).Scan(&id)
	if id == 0 {
		err = db.QueryRow(
			`INSERT INTO album (album_name, album_artist) VALUES ($1, $2) RETURNING id`,
			album.AlbumName, album.AlbumArtist,
		).Scan(&id)
		createdNew = true
	}
	album.AlbumID=id
	return id, createdNew, err
}

func (album *Album) SaveAlbumImageID(db *sqlx.DB, id string){
	rows, err := db.Queryx(`UPDATE album SET album_cover = $1 WHERE id=$2`, id, album.AlbumID)
	if err != nil {
		rows, err = db.Queryx(`UPDATE album SET album_cover = $1 WHERE id=$2`, default_id, album.AlbumID)
	}
	defer rows.Close()
}

func GetAlbum(db *sqlx.DB) ([]Album, error) {

	rows, err := db.Queryx(
		`SELECT * FROM album order by album_name`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albums := make([]Album, 0)
	for rows.Next() {
		var album Album
		err := rows.StructScan(&album)
		if err != nil {
			fmt.Println(err)
		}
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

func ClearAlbum(db *sqlx.DB) error {
	sqlStatement := `DELETE from album`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
