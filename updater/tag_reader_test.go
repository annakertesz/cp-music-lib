package updater

import (
	"bytes"
	"fmt"
	"github.com/dhowden/tag"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"testing"
)

func TestTimestamp(t *testing.T){
	fmt.Println(hash("blablabla"))
	fmt.Println(hash("blablabla"))
	fmt.Println(hash("blablabla"))
}

func hash(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}


func TestUpdate(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	Update(11056063660, 345345, "NBHZNGZKqsIIQnNQOQhlUsOMqa1msDId", db )
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

//func TestUploadSong(t *testing.T) {
//	db, err := connectToDB()
//	if err!=nil {
//		panic("couldnt connect to db")
//	}
//	readFile, err := ioutil.ReadFile("../sources/instr.mp3")
//	UploadSong(readFile, 124324, db)
//}

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
