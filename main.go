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
	url, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		url = psqlInfo
	}

	db, err = connect(url)

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
drop table if exists users;

drop table if exists tag_song;

drop table if exists song;

drop table if exists tag;

drop table if exists album;

drop table if exists artist;
`)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS artist
(
    id          SERIAL NOT NULL,
    artist_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag
(
    id       SERIAL NOT NULL,
    tag_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS album
(
    id                  SERIAL NOT NULL,
    album_name          varchar(150),
    album_artist        SERIAL REFERENCES artist (id),
    album_cover varchar(500),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS song
(
    id                  SERIAL NOT NULL,
    song_name           varchar(150),
    song_album          INTEGER REFERENCES album (id),
    song_tag            INTEGER REFERENCES tag (id),
    song_lq_url         varchar(500),
    song_hq_url         varchar(500),
    instrumental_lq_url varchar(500),
    instrumental_hq_url varchar(500),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag_song
(
    id       SERIAL NOT NULL,
    map_tag  INTEGER REFERENCES tag (id),
    map_song INTEGER REFERENCES song (id),
    PRIMARY KEY (id)
);
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
