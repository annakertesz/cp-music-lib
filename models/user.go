package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email string `json:"email" db:"email"`
	PasswordHash string `db:"password_hash"`
	Phone string `json:"phone" db:"phone"`
	UserStatus string `json:"user_status" db:"user_status"`
}

type UserReqObj struct {
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email string `json:"email" db:"email"`
	Password string `db:"password"`
	Phone string `json:"phone" db:"phone"`
}

type UserRespObj struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
}

type UserValidationObj struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func GetUsers(db *sqlx.DB) ([]User, error) {
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

		err = rows.Scan(&u.ID, &u.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func UnmarshalUserValidation(r *http.Request) (*UserValidationObj, error){
	defer r.Body.Close()

	var user UserValidationObj

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


func UnmarshalUser(r *http.Request) (*UserReqObj, error) {
	defer r.Body.Close()

	var user UserReqObj

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

func CreateUser(db *sqlx.DB, user *UserReqObj) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO cp_user (username, first_name, last_name, email, password_hash, phone, user_status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		user.Username,  user.FirstName, user.LastName, user.Email, hash(user.Password), user.Phone, 1,
	).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdateUserStatus(db *sqlx.DB, id int) error {
	rows, err := db.Query(`UPDATE cp_user SET user_status = 2 WHERE id=$1`, id)

	if err != nil {
		fmt.Println("error set user status")
		return err
	}
	defer rows.Close()
	return nil
}

func CheckUserCredentials(db *sqlx.DB, username string, password string) int{
	var dbPasswordHash string
	var user_status int
	var id int
	err := db.QueryRow(
		`SELECT id, password_hash, user_status from cp_user where username = $1`, username,
	).Scan(&id, &dbPasswordHash, &user_status)
	fmt.Println(dbPasswordHash)
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password)))
	if (err ==nil) && (bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password)) == nil) && (user_status == 2){
		return id
	}
	return -1
}

func CreateSession(db *sqlx.DB, userID int, uuid string) error{
	var id int
	expiration := time.Now().AddDate(0,0,7).Format("2006-01-02 15:04:05")
	err := db.QueryRow(
		`INSERT INTO sessions (session_id, cp_user, expiration) VALUES ($1, $2, $3) RETURNING id`,
		uuid, userID, expiration,
	).Scan(&id)
	if err != nil || id == 0 {
		return errors.New("couldnt create session")
	}
	return nil
}

func hash(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}