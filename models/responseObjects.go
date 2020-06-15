package models

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
	ID int `json:"id"`
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

type UserReqObj struct {
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email string `json:"email" db:"email"`
	Password string `db:"password"`
	Phone string `json:"phone" db:"phone"`
}

type PlaylistReqObj struct {
	Title string `json:"title"`
}

type UserRO struct {
	ID       int  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
}

type UserValidationObj struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}