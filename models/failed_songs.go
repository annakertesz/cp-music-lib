package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FailedSong struct {
	Id int `db:"id"`
	BoxId int `db:"box_id"`
	ErrorLogID int `db:"error_log_id"`
	Update int `db:"update"`
}

func SaveFailedSong(db *sqlx.DB, boxID string, logId int, update int) (int, error){
	var id int
	err := db.QueryRow(
		`INSERT INTO failed_song (box_id, ErrorLogID, update) VALUES ($1, $2, $3) RETURNING id`,
		boxID,  logId, update,
	).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return id, err
}