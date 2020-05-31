package controller

import (
	"fmt"
	"github.com/go-chi/chi"
"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Server struct {
	db *sqlx.DB
	token string
	musicFolder int
	coverFolder int
}

type e map[string]string

func NewServer(db *sqlx.DB, token string, musicFolder, coverFolder int) *Server {
	return &Server{db:db, token:token, musicFolder:musicFolder, coverFolder:coverFolder}
}

func (server *Server) Routes() chi.Router {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getIndex")
		fmt.Fprint(w, "cp")
	})

	//Songs
	r.Get("/song", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getAllSong")
		getAllSongs(server.db, w, r)
	})
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
	//TODO
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
		update(server.db, w, r, server.token, server.coverFolder, server.musicFolder)
	})

	r.Get("/download/{boxID}", func(w http.ResponseWriter, r *http.Request) {
		download(server.db, server.token, w, r)
	})


	//Playlist
	//TODO
	r.Get("/playlist/findByUser", func(w http.ResponseWriter, r *http.Request) {
		getPlaylistByUser(server.db, w, r)
	})
	//TODO
	r.Get("/playlist/{playlistID}", func(w http.ResponseWriter, r *http.Request) {
		getPlaylistById(server.db, w, r)
	})
	return r
}

