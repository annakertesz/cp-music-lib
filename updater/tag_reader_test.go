package updater

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	Update(11056063660, "2017-05-15T13:35:01-07:00", "sddXCkiTVNS8YC42WAAtzd07YIioPDCt", db )
}

func TestUploadSong(t *testing.T) {
	db, err := connectToDB()
	if err!=nil {
		panic("couldnt connect to db")
	}
	file, _ := os.Open("../sources/music/Dorothy.mp3")
	UploadSong(file, 124324, db)
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
