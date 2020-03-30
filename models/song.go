package models

import (
	"github.com/jmoiron/sqlx"
)

type Song struct {
	SongID int `json:"song_id" db:"id"`
	SongName string `json:"song_name" db:"song_name"`
	SongLqURL string `json:"song_lq_url" db:"song_lq_url"`
	SongHqURL string `json:"song_hq_url" db:"song_hq_url"`
	SongInstrumentalLqURL string `json:"song_instrumental_lq_url" db:"instrumental_lq_url"`
	SongInstrumentalHqURL string `json:"song_instrumental_hq_url" db:"instrumental_hq_url"`
	SongAlbum *Album `json:"song_album" db:"song_album"`
	SongArtist *Artist `json:"song_artist" db:"song_artist"`
	SongTags []*Tag `json:"song_tags" db:"song_tag"`
}

func (song *Song) CreateSong(db *sqlx.DB) error {

	row := db.QueryRow(
		`INSERT INTO song (song_name, song_artist, song_album, song_lq_url, song_hq_url, instrumental_lq_url, instrumental_hq_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		song.SongName, song.SongArtist.ArtistID, song.SongAlbum.AlbumID, song.SongLqURL, song.SongHqURL, song.SongInstrumentalLqURL, song.SongInstrumentalHqURL,
	)

	err := row.Scan(&song.SongID)

	if err != nil {
		return err
	}

	return nil
}

//func GetSongByTitleAndAlbum(title string, album string, db *sqlx.DB) ([]Song, error) {
//
//	//Get song object
//	rows, err := db.Queryx(
//		`SELECT id, song_name, song_artist, song_album, song_lq_url, song_hq_url, instrumental_lq_url, instrumental_hq_url FROM song WHERE song.song_name = $1` , title,
//	)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	songs := make([]Song, 0, 5)
//
//	for rows.Next() {
//		var id int
//		var song_name string
//		var song_artist int
//		var song_album int
//		var song_lq_url string
//		var song_hq_url string
//		var instrumental_lq_url string
//		var instrumental_hq_url string
//
//		err := rows.Scan(&id, &song_name, &song_artist, &song_album, &song_hq_url, &song_lq_url, &instrumental_hq_url, &instrumental_lq_url)
//		if err != nil {
//			fmt.Print(err.Error())
//			return nil, err
//		}
//		//artist, err := GetArtistByID(song_artist, db)
//		album, err := GetAlbumByID(song_album, db)
//		tags, err := GetTagsOfSong(id, db)
//		if err != nil {
//			fmt.Print(err.Error())
//			return nil, err
//		}
//
//		songs = append(songs, Song{
//			SongID:                id,
//			SongName:              song_name,
//			SongLqURL:             song_lq_url,
//			SongHqURL:             song_hq_url,
//			SongInstrumentalLqURL: instrumental_lq_url,
//			SongInstrumentalHqURL: instrumental_hq_url,
//			SongAlbum:             &album,
//			SongArtist:            album.AlbumArtist,
//			SongTags:              tags,
//		})
//	}
//
//	return songs, nil
//}