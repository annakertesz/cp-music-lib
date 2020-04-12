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

func GetTagsOfSong(id int, db *sqlx.DB) ([]Tag, error){
		rows, err := db.Queryx(
			`SELECT t.id, t.tag_name FROM tag_song join tag t on tag_song.map_tag = t.id WHERE map_song = $1` , id,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		tags := make([]Tag, 0, 5)
		for rows.Next() {
			var tag Tag
			err := rows.StructScan(&tag)
			if err!=nil{
				return nil, err
			}
			tags = append(tags, tag)
		}
		return tags, nil
}
