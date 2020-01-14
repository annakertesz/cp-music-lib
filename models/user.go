package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func Users(db *sql.DB) ([]User, error) {
	rows, err := db.Query(
		`SELECT id, username FROM users ORDER BY username`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0, 5)

	for rows.Next() {
		u := User{}

		err = rows.Scan(&u.Id, &u.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}


func UnmarshalUser(r *http.Request) (*User, error) {
	defer r.Body.Close()

	var user User

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(db *sql.DB, username string) (*User, error) {
	created := User{}

	row := db.QueryRow(
		`INSERT INTO users (username) VALUES ($1) RETURNING id, username`,
		username+"new",
	)

	err := row.Scan(&created.Id, &created.Username)

	if err != nil {
		return nil, err
	}

	return &created, nil
}