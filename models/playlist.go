package models

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

type Playlist struct {
	Title string `json:"title" db:"title"`
	User  int    `json:"user" db:"cp_user"`
}

func CreatePlaylist(db *sqlx.DB, playlist Playlist) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO playlist (title, cp_user) VALUES ($1, $2) RETURNING id`,
		playlist.Title, playlist.User,
	).Scan(&id)

	return id, err
}


func UnmarshalPlaylist(r *http.Request) (*Playlist, error) {
	defer r.Body.Close()

	var playlist Playlist

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &playlist)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
