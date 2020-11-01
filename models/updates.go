package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Update struct {
	Id int `db:"id"`
	UdDate string `db:"ud_date"`
	FoundSongs int `db:"found_songs"`
	CreatedSongs int `db:"created_songs"`
	FailedSongs int `db:"failed_songs"`
	DeletedSongs int `db:"deleted_songs"`
}

func GetLatestUpdate(db *sqlx.DB) (string, *ErrorModel) {
	var upds int
	err := db.QueryRowx(
		`select count(1) where exists (select * from cp_update)`,
	).Scan(&upds)
	if err != nil {
		return "", &ErrorModel{
			Service: "updates_table",
			Err:     err,
			Message: "error while count updates",
			Sev:     3,
		}
	}
	if upds == 0 {
		return "2010-05-15T13:35:01-07:00", nil
	}
	var cTime time.Time
	err = db.QueryRow(
		`SELECT MAX(ud_date) FROM cp_update`,
		).Scan(&cTime)
	if err != nil {
		return "", &ErrorModel{
			Service: "updates_table",
			Err:     err,
			Message: "error while try ro find the last update",
			Sev:     3,
		}
	}
	date, month, day := cTime.Date()
	fmt.Println(date, month,day)
	return cTime.Format(time.RFC3339), nil
}

func NewUpdate(db *sqlx.DB)(int, *ErrorModel){
	var id int
	err := db.QueryRow(
		`INSERT INTO cp_update (ud_date) VALUES ($1) RETURNING id`,
		time.Now(),
	).Scan(&id)
	if err != nil {
	return id, &ErrorModel{
		Service: "UpdateLogWriter",
		Err:     err,
		Message: "There was a problem while insert new Update log (cp_update table)",
		Sev:     3,
	}}
	return id, nil
}

func SaveUpdateNumbers(db *sqlx.DB, id int, found int, created int, failed int, deleted int) error{
	rows, err := db.Query(
		`UPDATE cp_update SET found_songs = $1 ,created_songs = $2, failed_songs = $3, deleted_songs = $4 WHERE id = $5`,
		found,  created, failed, deleted, id,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func ClearUpdates(db *sqlx.DB) error {
	sqlStatement := `DELETE from cp_update`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
