package models

import (
	"github.com/jmoiron/sqlx"
)

type Artist struct {
	ArtistID int `json:"artist_id" db:"id"`
	ArtistName string `json:"artist_name" db:"artist_name"`
}

func (artist *Artist) CreateArtist(db *sqlx.DB) error {

	row := db.QueryRow(
		`INSERT INTO artist (artist_name) VALUES ($1) RETURNING id`,
		artist.ArtistName,
	)

	err := row.Scan(&artist.ArtistID)

	if err != nil {
		return err
	}

	return nil
}

func GetArtistByID(id int, db *sqlx.DB) (Artist, error){
	rows, err := db.Queryx(
		`SELECT * FROM artist WHERE id = $1` , id,
	)
	if err != nil {
		return Artist{}, err
	}
	defer rows.Close()
	var artist Artist
	for rows.Next() {
		rows.StructScan(&artist)
	}
	return artist, nil
}