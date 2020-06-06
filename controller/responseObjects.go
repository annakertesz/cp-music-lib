package controller

type TagRO struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type SongRO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Album AlbumRO `json:"album"`
	LqSong *string `json:"lq_song"`
	HqSong *string `json:"hq_song"`
	LqInstr *string `json:"lq_instr"`
	HqInstr *string `json:"hq_instr"`
	Tags []TagRO
}

type PlaylistRO struct {
	Title string `json:"title" db:"title"`
	User  UserRO `json:"user" db:"cp_user"`
	Songs []SongRO
}

type AlbumRO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Artist ArtistRO `json:"artist"`
	Cover *string `json:"cover"`
}

type ArtistRO struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type UserRO struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Phone string `json:"phone"`
	userStatus string `json:"user_status"`
}