package updater

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestUploadSong(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	file, _ := os.Open("../sources/The Somersault Boy_ Hate Love Hate (Instrumental) (1).mp3")
	UploadSong(file, "be1423", db)
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
