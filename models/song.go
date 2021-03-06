package models

import (
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type Song struct {
	SongID                int     `json:"song_id" db:"id"`
	SongName              string  `json:"song_name" db:"song_name"`
	SongLqURL             *string `json:"song_lq_url" db:"song_lq_url"`
	SongHqURL             *string `json:"song_hq_url" db:"song_hq_url"`
	SongInstrumentalLqURL *string `json:"song_instrumental_lq_url" db:"instrumental_lq_url"`
	SongInstrumentalHqURL *string `json:"song_instrumental_hq_url" db:"instrumental_hq_url"`
	SongAlbum             int     `json:"song_album" db:"song_album"`
	boxID                 int
}

func NewSong(name string, album int, boxID int) Song {
	return Song{
		SongName:  name,
		SongAlbum: album,
		boxID:     boxID,
	}
}

func (song *Song) CreateSong(db *sqlx.DB) (int, bool, error) {
	isInstrumental := false
	createdNew := false
	var id int
	if strings.Contains(song.SongName, "nstrumental") {
		if strings.Index(song.SongName, "(") > 0 {
			song.SongName = strings.TrimSpace(song.SongName[:strings.Index(song.SongName, "(")])
			isInstrumental = true
		}
	}
	err := db.QueryRow(
		`SELECT id from song where song_name = $1`, song.SongName,
	).Scan(&id)
	if id != 0 {
		if isInstrumental {
			rows, err := db.Query(`UPDATE song SET instrumental_lq_url = $1 WHERE id=$2`, song.boxID, id)

			if err != nil {
				log.Println("error in update song query")
				return 0, false, err
			}
			defer rows.Close()
		} else {
			rows, err := db.Query(`UPDATE song SET song_lq_url = $1 WHERE id=$2`, song.boxID, id)

			if err != nil {
				log.Println("error in update song query")
				return 0, false, err
			}
			defer rows.Close()
		}
	} else {
		if isInstrumental {
			err = db.QueryRow(
				`INSERT INTO song (song_name, song_album,instrumental_lq_url) VALUES ($1, $2, $3) RETURNING id`,
				song.SongName, song.SongAlbum, song.boxID,
			).Scan(&id)
		} else {
			err = db.QueryRow(
				`INSERT INTO song (song_name, song_album,song_lq_url) VALUES ($1, $2, $3) RETURNING id`,
				song.SongName, song.SongAlbum, song.boxID,
			).Scan(&id)
		}
		createdNew = true
	}

	return id, createdNew, err
}

func GetSongByID(id int, db *sqlx.DB) (*Song, error) {
	var song Song
	err := db.QueryRowx(`SELECT * FROM song WHERE id = $1`, id,
	).StructScan(&song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func GetSongByArtist(id int, db *sqlx.DB) ([]Song, error) {
	rows, err := db.Queryx(
		`SELECT song.id, song.instrumental_hq_url, song.instrumental_lq_url, song.song_album, song.song_hq_url, song.song_lq_url, song.song_name from song join album on song.song_album = album.id where album_artist = $1`, id,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in scan songs: %v", err.Error())
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetAllSongs(db *sqlx.DB) ([]Song, error) {
	rows, err := db.Queryx(
		`SELECT song.id, song.instrumental_hq_url, song.instrumental_lq_url, song.song_album, song.song_hq_url, song.song_lq_url, song.song_name from song join album on song.song_album = album.id`,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in scan songs %v", err.Error())
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSongsByPlaylist(db *sqlx.DB, playlistID int) ([]Song, error) {
	rows, err := db.Queryx(
		`SELECT 
				song.id, 
				song.instrumental_hq_url, 
				song.instrumental_lq_url, 
				song.song_album, 
				song.song_hq_url, 
				song.song_lq_url, 
				song.song_name
			FROM song JOIN playlist_song ON song.id = playlist_song.map_song 
			WHERE playlist_song.map_playlist = $1`, playlistID,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in scan songs: %v ", err.Error())
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSongByAlbum(id int, db *sqlx.DB) ([]Song, error) {
	rows, err := db.Queryx(
		`SELECT * FROM song WHERE song_album = $1`, id,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in query: %v", err.Error())
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSongByTag(id int, db *sqlx.DB) ([]Song, error) {
	rows, err := db.Queryx(
		`select song.id, song.instrumental_hq_url, song.instrumental_lq_url, song.song_album, song.song_hq_url, song.song_lq_url, song.song_name from tag_song join tag on tag.id = tag_song.map_tag join song on song.id = tag_song.map_song where tag.id = $1`, id,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in query: %v", err.Error())
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSongByEverything(keyword string, db *sqlx.DB)([]Song, error) {
	wildchart := "%" + keyword + "%"
	rows, err := db.Queryx(
		`select 
					song.id, 
					song.instrumental_hq_url, 
					song.instrumental_lq_url, 
					song.song_album, 
					song.song_hq_url, 
					song.song_lq_url, 
					song.song_name
				from tag_song 
				join tag on tag.id = tag_song.map_tag join song on song.id = tag_song.map_song 
				join album on song.song_album = album.id
				join artist on album.album_artist = artist.id
				where tag.tag_name ilike $1 or artist.artist_name ilike $1 or album.album_name ilike $1 or song.song_name ilike $1 
				group by song.id`, wildchart,
	)
	if err != nil {
		log.Printf("error in query: %v", err.Error())
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.StructScan(&song)
		if err != nil {
			log.Printf("error in query: %v", err.Error())
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func ClearSong(db *sqlx.DB) error {
	sqlStatement := `DELETE from song`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
