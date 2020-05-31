package models

import (
	"fmt"
	"testing"
)

func TestGetLatestUpdate(t *testing.T) {
	db, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}
	//location, err := time.LoadLocation("Europe/Budapest")
	//if err != nil {
	//	panic(err.Error())
	//}
	//var id int
	//err = db.QueryRow(
	//	`INSERT INTO update (ud_date,found_songs,created_songs, failed_songs, deleted_songs) VALUES ($1, 1, 2, 3, 4) RETURNING id`,
	//	time.Date(2017, 02, 01, 00,00,00,00, location),
	//).Scan(&id)
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = db.QueryRow(
	//	`INSERT INTO update (ud_date,found_songs,created_songs, failed_songs, deleted_songs) VALUES ($1, 1, 2, 3, 4) RETURNING id`,
	//	time.Date(2017, 02, 01, 00,00,00,00, location),
	//).Scan(&id)
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = db.QueryRow(
	//	`INSERT INTO update (ud_date,found_songs,created_songs, failed_songs, deleted_songs) VALUES ($1, 16, 25, 36, 43) RETURNING id`,
	//	time.Date(2018, 02, 01, 00,00,00,00, location),
	//).Scan(&id)
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = db.QueryRow(
	//	`INSERT INTO update (ud_date,found_songs,created_songs, failed_songs, deleted_songs) VALUES ($1, 15, 24, 37, 43) RETURNING id`,
	//	time.Date(2019, 02, 01, 00,00,00,00, location),
	//).Scan(&id)
	//if err != nil {
	//	panic(err.Error())
	//}
	update, err := GetLatestUpdate(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(update)
}

func TestASaveFailedSong(t *testing.T){
	db, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}
	SaveFailedSong(db, "12345", "errormessage", 1)
}

func TestNewUpdate(t *testing.T) {
	db, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}
	update, err := NewUpdate(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(update)
	latestUpdate, _ := GetLatestUpdate(db)
	fmt.Println(latestUpdate)
}

func TestSaveUpdateNumbers(t *testing.T) {
	db, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}
	err = SaveUpdateNumbers(db, 1, 100, 100, 100, 100)
	if err != nil {
		panic(err.Error())
	}
}