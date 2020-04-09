package updater

import (
	"bytes"
	"fmt"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"testing"
)

func TestUpdate(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	Update(11056063660, "2017-05-15T13:35:01-07:00", "NBHZNGZKqsIIQnNQOQhlUsOMqa1msDId", db )
}

func TestTagReader(t *testing.T){
	readFile, err := ioutil.ReadFile("../sources/instr.mp3")
	if err!=nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(readFile)
	from, err := tag.ReadFrom(reader)
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(from)
}

func TestUploadSong(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	readFile, err := ioutil.ReadFile("../sources/instr.mp3")
	UploadSong(readFile, 124324, db)
}

func connectToDB()(*sqlx.DB, error) {
	url := "host=localhost port=5432 user=anna password=gfd dbname=centralp sslmode=disable"
	db, err := sqlx.Open("postgres", url)

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
