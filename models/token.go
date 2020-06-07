package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateToken(db *sqlx.DB) (string, error){
	token := uuid.New()
	sqlStatement := `INSERT INTO token (token) VALUES ($1)`
	_, err := db.Exec(sqlStatement, token.String())
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func GetToken(db *sqlx.DB, token string) (int, error) {
	var id int
	err := db.QueryRowx(`SELECT id FROM token WHERE token = $1`, token,
	).Scan(&id)
	return id, err
}

func DeleteToken(db *sqlx.DB, id int) error {
	sqlStatement := `DELETE FROM token WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	return err
}