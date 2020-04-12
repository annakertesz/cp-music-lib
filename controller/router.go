package controller

import (
	"fmt"
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
		fmt.Fprint(w, "cp")
	})

	//Songs
	r.Get("/song/{songID}", func(w http.ResponseWriter, r *http.Request) {
		getSongByID(server.db, w, r)
	})
	r.Get("/song/findByAlbum", func(w http.ResponseWriter, r *http.Request) {
		getSongByAlbum(server.db, w, r)
	})
	r.Get("/song/findByArtist", func(w http.ResponseWriter, r *http.Request) {
		getSongByArtist(server.db, w, r)
	})
	r.Get("/song/findByTag", func(w http.ResponseWriter, r *http.Request) {
		getSongByTag(server.db, w, r)
	})
	r.Get("/song/findByFreeSearch", func(w http.ResponseWriter, r *http.Request) {
		searchSong(server.db, w, r)
	})

	//Albums
	r.Get("/album", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("album request")
		getAllAlbum(server.db, w, r)
	})
	r.Get("/album/findByArtist", func(w http.ResponseWriter, r *http.Request) {
		getAlbumsByArtist(server.db, w, r)
	})
	r.Get("/album/{albumID}", func(w http.ResponseWriter, r *http.Request) {
		getAlbumsById(server.db, w, r)
	})

	//Artists
	r.Get("/artist", func(w http.ResponseWriter, r *http.Request) {
		getAllArtist(server.db, w, r)
	})
	r.Get("/artist/{artistID}", func(w http.ResponseWriter, r *http.Request) {
		getArtistById(server.db, w, r)
	})
	
	//Tags
	r.Get("/tag", func(w http.ResponseWriter, r *http.Request) {
		getAllTag(server.db, w, r)
	})

	r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
		update(server.db, w, r)
	})

	r.Get("/download/{boxID}", func(w http.ResponseWriter, r *http.Request) {
		download(server.db, w, r)
	})


	//Playlist
	r.Get("/playlist/findByUser", func(w http.ResponseWriter, r *http.Request) {
		getPlaylistByUser(server.db, w, r)
	})
	r.Get("/playlist/{playlistID}", func(w http.ResponseWriter, r *http.Request) {
		getPlaylistById(server.db, w, r)
	})
	return r
}

