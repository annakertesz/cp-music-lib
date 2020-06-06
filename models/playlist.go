package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Playlist struct {
	Title string `json:"title" db:"title"`
	User  int    `json:"user" db:"cp_user"`
}

func CreatePlaylist(db *sqlx.DB, userID int, title string) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO playlist (title, cp_user) VALUES ($1, $2) RETURNING id`,
		title,userID,
	).Scan(&id)

	return id, err
}

func GetPlaylistByID(db *sqlx.DB, playlistID int) (Playlist, error){
	var playlist Playlist
	err := db.QueryRowx(
		`SELECT * FROM playlist WHERE id = $1`, playlistID,
	).StructScan(&playlist)

	return playlist, err
}

func GetAllPLaylist(db *sqlx.DB, userID int) ([]Playlist, error){
	rows, err := db.Queryx(
		`SELECT * FROM playlist WHERE cp_user = $1`, userID,
	)
	if err != nil {
		fmt.Println("error in query")
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	playlists := make([]Playlist, 0)
	for rows.Next() {
		var playlist Playlist
		err := rows.StructScan(&playlist)
		if err != nil {
			fmt.Println("error in scan playlist")
			fmt.Println(err.Error())
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

