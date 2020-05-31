package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FailedSong struct {
	Id int `db:"id"`
	BoxId int `db:"box_id"`
	ErrorMessage string `db:"error_message"`
	Update int `db:"update"`
}

func SaveFailedSong(db *sqlx.DB, boxID string, error string, update int) (int, error){
	var id int
	err := db.QueryRow(
		`INSERT INTO failed_song (box_id, error_message, update) VALUES ($1, $2, $3) RETURNING id`,
		boxID,  error, update,
	).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return id, err
}