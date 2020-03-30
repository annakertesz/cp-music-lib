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

func GetArtistByID(id int, db *sqlx.DB) (*Artist, error){
	var artist Artist
	err := db.QueryRowx(
		`SELECT * FROM artist WHERE id = $1` , id,
	).StructScan(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func GetArtist(db *sqlx.DB) ([]Artist, error){
	rows, err := db.Queryx(
		`SELECT * FROM artist`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var artists []Artist

	for rows.Next() {
		var artist Artist
		rows.StructScan(&artist)
		artists = append(artists, artist)
	}
	return artists, nil
}