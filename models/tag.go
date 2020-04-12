package models

import (
	"github.com/jmoiron/sqlx"
)

type Tag struct {
	TagID int `json:"tag_id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name"`
}

func (tag *Tag) CreateTag(db *sqlx.DB) (int, bool) {
	createdNew := false
	var id int
	db.QueryRow(
		`SELECT id from tag where tag_name = $1`, tag.TagName,
	).Scan(&id)
	if id==0{
		db.QueryRow(
			`INSERT INTO tag (tag_name) VALUES ($1) RETURNING id`,
			tag.TagName,
		).Scan(&id)
		createdNew = true
	}
	return id, createdNew
}

func GetTag(db *sqlx.DB) ( []Tag, error){
	rows, err := db.Queryx(
		`SELECT * FROM tag`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []Tag

	for rows.Next() {
		var tag Tag
		rows.StructScan(&tag)
		tags = append(tags, tag)
	}
	return tags, nil
}

func GetTagsOfSong(db *sqlx.DB, songID int) ( []Tag, error){
	rows, err := db.Queryx(
		`SELECT tag.id, tag.tag_name from tag join tag_song on tag_song.map_tag=tag.id where tag_song.map_song= $1`, songID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []Tag

	for rows.Next() {
		var tag Tag
		rows.StructScan(&tag)
		tags = append(tags, tag)
	}
	return tags, nil
}
