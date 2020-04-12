package models

import (
	"github.com/jmoiron/sqlx"
)

type TagSong struct {
	TagSongID int `json:"tag_song_id"`
	Tag int`json:"tag"`
	Song int `json:"song"`
}

func (tagSong *TagSong) CreateTagSong(db *sqlx.DB) int {
	var id int
	db.QueryRowx(`SELECT id from tag_song where map_tag = $1 and map_song = $2`).Scan(&id)
	if id==0{
		db.QueryRow(
			`INSERT INTO tag_song (map_tag, map_song) VALUES ($1, $2) RETURNING id`,
			tagSong.Tag, tagSong.Song,
		).Scan(&id)
	}
	return id
}
