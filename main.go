package main

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/config"
	"github.com/annakertesz/cp-music-lib/controller"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services/updater"
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
	testFolder = "114476926207"
	cpFolder = "11059102688"
	)


func main() {

	// set up server
	config := getConfig()
	db, err := connect(config.Url)
	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	server := controller.NewServer(db, config)

	//First update
	clearDB(true, db)
	updater.Update(config.SongFolder, config.CoverFolder, server.BoxConfig.Token, db)

	// schedule update for every day
	updaterfunc := func() {
		err = updater.Update(config.SongFolder, config.CoverFolder, server.BoxConfig.Token, db)
		if err != nil {
			server.GetBoxToken()
			err = updater.Update(config.SongFolder, config.CoverFolder, server.BoxConfig.Token, db)
			if err != nil {
				fmt.Printf("couldnt update %v", err.Error())
			}
		}
	}
	updater.Update(config.SongFolder, config.CoverFolder, server.BoxConfig.Token, db)
	scheduler.Every(1).Day().Run(updaterfunc)

	//start server
	log.Println("Started")
	if err := http.ListenAndServe(":"+port, server.Routes()); err != nil {
		log.Fatal("Could not start HTTP server", err.Error())
	}
}
func getConfig() config.Config {
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
	sengridAPIKey, ok := os.LookupEnv("SENGRID_API_KEY")
	if !ok {
		panic("need box credentials: SENGRID_API_KEY")
	}
	senderName, ok := os.LookupEnv("SENDER_NAME")
	if !ok {
		panic("need box credentials: SENDER_NAME")
	}
	senderEmail, ok := os.LookupEnv("SENDER_EMAIL")
	if !ok {
		panic("need box credentials: SENDER_EMAIL")
	}
	adminEmail, ok := os.LookupEnv("ADMIN_EMAIL")
	if !ok {
		panic("need box credentials: ADMIN_EMAIL")
	}
	developerEmail, ok := os.LookupEnv("DEVELOPER_EMAIL")
	if !ok {
		panic("need box credentials: DEVELOPER_EMAIL")
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
	return config.Config{
		BoxConfig:     config.BoxConfig{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			PrivateKey:   privateKey,
		},
		SengridConfig: config.SengridConfig{
			SengridAPIKey:  sengridAPIKey,
			SenderName:     senderName,
			SenderEmail:    senderEmail,
			AdminEmail:     adminEmail,
			DeveloperEmail: developerEmail,
		},
		SongFolder:    songFolder,
		CoverFolder:   coverFolder,
		PsqlInfo:      psqlInfo,
		Url:		   url,
	}
}

func clearDB(do bool, db *sqlx.DB) error {
	if do{
		err := models.ClearUpdates(db)
		err = models.ClearSong(db)
		err = models.ClearPlaylist(db)
		err = models.ClearTagSong(db)
		err = models.ClearTag(db)
		err = models.ClearAlbum(db)
		err = models.ClearArtist(db)
		if err != nil {
			return err
		}
	}
	return nil
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
   error_log_id INTEGER REFERENCES logs (id),
   cp_update       INTEGER REFERENCES cp_update (id),
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

CREATE TABLE IF NOT EXISTS token
(
   id       SERIAL NOT NULL,
   token  varchar(500)
);

CREATE TABLE IF NOT EXISTS playlist_song
(
   id       SERIAL NOT NULL,
   map_playlist  INTEGER REFERENCES playlist (id),
   map_song INTEGER REFERENCES song (id),
   PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS logs
(
	id       SERIAL NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
   service varchar(255),
   error varchar(500),
   message varchar(500)
   PRIMARY KEY (id)
);
 `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
