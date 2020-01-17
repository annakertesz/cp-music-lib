package main

import (
	"database/sql"
	"fmt"
	"github.com/annakertesz/cp-music-lib/controller"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	url, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatalln("$DATABASE_URL is required")
	}

	db, err = connect(url)

	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	log.Println("Started")
	if err := http.ListenAndServe(":"+port, controller.Routes(db)); err != nil {
		log.Fatal("Could not start HTTP server", err.Error())
	}
}

func connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id       SERIAL,
      username VARCHAR(64) NOT NULL UNIQUE,
      CHECK (CHAR_LENGTH(TRIM(username)) > 0)
    );
  `)
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS album (
      id       SERIAL,
      album_name VARCHAR(64) NOT NULL UNIQUE,
      
    );
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}

