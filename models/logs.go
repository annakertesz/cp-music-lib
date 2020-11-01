package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Logs struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at`
	Service   string    `db:"service"`
	Err       string    `db:"error"`
	Message   string    `db:"message"`
}

func CreateLog(db *sqlx.DB, service, error, message string)(int, error){
	var id int
	err := db.QueryRow(
		`INSERT INTO logs (service, error, message) VALUES ($1, $2, $3) RETURNING id`,
		service, error, message,
	).Scan(&id)
	return id, err
}
