package models

type Album struct {
	AlbumID string `json:"album_id"`
	AlbumName string `json:"album_name"`
	AlbumArtist Artist `json:"artist"`
	AlbumCoverUrl string `json:"cover_url"`
	AlbumCoverThumbnailUrl string `json:"cover_thumbnail_url"`
}