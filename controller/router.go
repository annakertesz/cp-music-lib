package controller

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Server struct {
	db *sqlx.DB
}

type e map[string]string

func NewServer(db *sqlx.DB) *Server {
	return &Server{db:db}
}

func (server *Server) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getUsers(server.db, w, r)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request){
		createUser(server.db, w, r)
	})

	r.Get("/getSongs", func(w http.ResponseWriter, r *http.Request) {
		//params: album, artist, tag, id, playlist

	})

	r.Get("/getArtists", func(w http.ResponseWriter, r *http.Request) {
		//params: id
	})

	r.Get("/getAlbum", func(w http.ResponseWriter, r *http.Request) {
		//params: id, artist
	})

	r.Get("/getTags", func(w http.ResponseWriter, r *http.Request) {
		//params: id
	})

	r.Get("/search/", func(w http.ResponseWriter, r *http.Request) {
		//params: keyword
	})

	r.Get("/getPlaylist", func(w http.ResponseWriter, r *http.Request) {
		//params: user
	})

	r.Get("/getPlaylist", func(w http.ResponseWriter, r *http.Request) {
		//params: user
	})

	

	return r
}

