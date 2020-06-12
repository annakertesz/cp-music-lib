package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

type Playlist struct {
	ID int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	User  int    `json:"user" db:"cp_user"`
}

func UnmarshalPlaylist(r *http.Request) (*PlaylistReqObj, error) {
	defer r.Body.Close()

	var playlist PlaylistReqObj

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

func CreatePlaylist(db *sqlx.DB, userID int, title string) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO playlist (title, cp_user) VALUES ($1, $2) RETURNING id`,
		title,userID,
	).Scan(&id)

	return id, err
}

func DeletePlaylist(db *sqlx.DB, playlistID int) error {
	sqlStatement := `DELETE from playlist_song WHERE map_playlist =$1`
	_, err := db.Exec(sqlStatement, playlistID)
	if err != nil {
		return err
	}
	sqlStatement = `DELETE from playlist WHERE id = $1`
	_, err = db.Exec(sqlStatement, playlistID)
	if err != nil {
		return err
	}
	return nil
}

func AddSongToPlayist(db *sqlx.DB, songID int, playlistID int) error{
	sqlStatement := `INSERT INTO playlist_song (map_playlist, map_song) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, playlistID, songID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveSongFromPlayist(db *sqlx.DB, songID int, playlistID int) error{
	sqlStatement := `DELETE from playlist_song WHERE map_song =$1 AND map_playlist = $2`
	_, err := db.Exec(sqlStatement, songID, playlistID)
	if err != nil {
		return err
	}
	return nil
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

