package main

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/controller"
	"github.com/annakertesz/cp-music-lib/updater"
	"github.com/carlescere/scheduler"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "anna"
	password = "anna"
	dbname   = "centralp"
	developerToken = "3luIrU57KaYyTb1eD3kP0iL3yjU90zwr"
	testFolder = "114476926207"
	cpFolder = "11059102688"
	)


func main() {
	songFolderStr, ok := os.LookupEnv("SONG_FOLDER")
	if !ok {
		songFolderStr = cpFolder
	}
	clientID, ok := os.LookupEnv("CLIENT_ID")
	if !ok {
		panic("need box credentials: CLIENT_ID")
	}
	clientSecret, ok := os.LookupEnv("CLIENT_SECRET")
	if !ok {
		panic("need box credentials: CLIENT_SECRET")
	}
	privateKey, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok {
		privateKeyData, err := ioutil.ReadFile("private.key")
		if err != nil {
			fmt.Println(err)
			panic("need box credentials: PRIVATE_KEY")
		}
		privateKey = string(privateKeyData)

	}
	songFolder, err := strconv.Atoi(songFolderStr)
	if err != nil {
		panic("songFolder var should be numeric")
	}
	coverFolderStr, ok := os.LookupEnv("COVER_FOLDER")
	if !ok {
		coverFolderStr=testFolder
	}
	coverFolder, err := strconv.Atoi(coverFolderStr)
	if err != nil {
		panic("coverFolder var should be numeric")
	}
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
	server := controller.NewServer(db, clientID, clientSecret, privateKey, songFolder, coverFolder)
	updater.Update(songFolder, coverFolder, server.Token, db)
	if !ok {
		port = "8080"
	}
	updater := func() {
		err = updater.Update(songFolder, coverFolder, server.Token, db)
		if err != nil {
			server.GetBoxToken()
			err = updater.Update(songFolder, coverFolder, server.Token, db)
			if err != nil {
				fmt.Printf("couldnt update %v", err.Error())
			}
		}
	}
	scheduler.Every(1).Day().Run(updater)
	log.Println("Started")
	if err := http.ListenAndServe(":"+port, server.Routes()); err != nil {
		log.Fatal("Could not start HTTP server", err.Error())
	}
}

func connect(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbURL)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
//
//	_, err = db.Exec(`
//drop table if exists users;
//
//drop table if exists tag_song;
//
//drop table if exists song;
//
//drop table if exists tag;
//
//drop table if exists album;
//
//drop table if exists artist;
//`)
//	if err != nil {
//		return nil, err
//	}
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

CREATE TABLE IF NOT EXISTS cp_update
(
    id           SERIAL NOT NULL,
    ud_date         DATE,
    found_songs   INTEGER,
    created_songs INTEGER,
    failed_songs  INTEGER,
    deleted_songs INTEGER,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS failed_song
(
    id           SERIAL NOT NULL,
    box_id        varchar(500),
    error_message varchar(500),
    cp_update       INTEGER REFERENCES cp_update (id),
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS sessions
(
    id           SERIAL NOT NULL,
    session_id        varchar(500),
    cp_user INTEGER REFERENCES cp_user (id),
    expiration       timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS playlist
(
    id           SERIAL NOT NULL,
    title        varchar(500),
    cp_user INTEGER REFERENCES cp_user (id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS playlist_song
(
    id       SERIAL NOT NULL,
    map_playlist  INTEGER REFERENCES playlist (id),
    map_song INTEGER REFERENCES song (id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cp_user
(
    id            SERIAL NOT NULL,
    username      varchar(150),
    first_name    varchar(150),
    last_name     varchar(150),
    email         varchar(150),
    password_hash varchar(150),
    phone         varchar(500),
    user_status   int,
    PRIMARY KEY (id)
);
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
