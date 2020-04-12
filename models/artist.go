package models

import (
	"github.com/jmoiron/sqlx"
)

type Artist struct {
	ArtistID int `json:"artist_id" db:"id"`
	ArtistName string `json:"artist_name" db:"artist_name"`
}

func (artist *Artist) CreateArtist(db *sqlx.DB) (int, error) {
	var id int
	err := db.QueryRow(
		`SELECT id FROM artist where artist_name = $1`, artist.ArtistName,
	).Scan(&id)
	if id==0{
		err = db.QueryRow(
			`INSERT INTO artist (artist_name) VALUES ($1) RETURNING id`,
			artist.ArtistName,
		).Scan(&id)
	}
	return id, err
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