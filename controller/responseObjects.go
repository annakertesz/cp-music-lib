package controller

type TagRO struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type SongRO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Album AlbumRO `json:"album"`
	LqSong string `json:"lq_song"`
	HqSong string `json:"hq_song"`
	LqInstr string `json:"lq_instr"`
	HqInstr string `json:"hq_instr"`
}

type AlbumRO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Artist ArtistRO `json:"artist"`
	Thumbnail string `json:"thumbnail"`
}

type ArtistRO struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type PlaylistRO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Songs []SongRO `json:"songs"`
}

type UserRo struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Phone string `json:"phone"`
	userStatus string `json:"user_status"`
}