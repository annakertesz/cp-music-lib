package models

type Song struct {
	SongID string `json:"song_id"`
	SongName string `json:"song_name"`
	SongLqURL string `json:"song_lq_url"`
	SongHqURL string `json:"song_hq_url"`
	SongInstrumentalLqURL string `json:"song_instrumental_lq_url"`
	SongInstrumentalHqURL string `json:"song_instrumental_hq_url"`
	SongAlbum Album `json:"song_album"`
	SongArtist Artist `json:"song_artist""`
	SongTags []Tag `json:"song_tags"`
}
