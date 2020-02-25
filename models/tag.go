package models

import (
	"github.com/jmoiron/sqlx"
)

type Tag struct {
	TagID string `json:"tag_id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name"`
}

func (tag *Tag) CreateTag(db *sqlx.DB) error {

	row := db.QueryRow(
		`INSERT INTO tag (tag_name) VALUES ($1) RETURNING id`,
		tag.TagName,
	)

	err := row.Scan(&tag.TagID)

	if err != nil {
		return err
	}

	return nil
}