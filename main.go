package main

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/controller"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var db *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "anna"
	password = "anna"
	dbname   = "centralp"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println("sdsd " + psqlInfo)

	db, err = connect(psqlInfo)

	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}
	server := controller.NewServer(db)

	log.Println("Started")
	if err := http.ListenAndServe(":"+port, server.Routes()); err != nil {
		log.Fatal("Could not start HTTP server", err.Error())
	}
}

func connect(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbURL)

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

	if err != nil {
		return nil, err
	}

	return db, nil
}

