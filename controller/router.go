package controller

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Server struct {
	db           *sqlx.DB
	Token        string
	musicFolder  int
	coverFolder  int
	clientID     string
	clientSecret string
	privateKey   string
}

type e map[string]string

func NewServer(db *sqlx.DB, clientID string, clientSecret string, privateKey string, musicFolder, coverFolder int) *Server {
	token := box_lib.AuthOfBox(clientID, clientSecret, privateKey)
	return &Server{db: db, Token: token, musicFolder: musicFolder, coverFolder: coverFolder, clientID: clientID, clientSecret: clientSecret, privateKey: privateKey}
}

func (s *Server) GetBoxToken() {
	s.Token = box_lib.AuthOfBox(s.clientID, s.clientSecret, s.privateKey)
}

func (server *Server) Routes() chi.Router {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Referer", "session", "Session"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             300, // Maximum value not ignored by any of major browsers
		OptionsPassthrough: false,
		Debug:              true,
	})
	r.Use(cors.Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getIndex")
		fmt.Fprint(w, "cp")
	})

	//Songs
	r.Get("/song", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAllSongs(server.db, w, r)
		}
	})
	r.Get("/song/{songID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getSongByID(server.db, w, r)
		}
	})
	r.Get("/song/findByAlbum", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getSongByAlbum(server.db, w, r)
		}
	})
	r.Get("/song/findByArtist", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getSongByArtist(server.db, w, r)
		}
	})
	r.Get("/song/findByTag", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getSongByTag(server.db, w, r)
		}
	})
	r.Post("/song/{songID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			addSongToPlaylist(server.db, w, r)
		}
	})
	r.Delete("/song/{songID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			removeSongFromPlaylist(server.db, w, r)
		}
	})
	//TODO
	r.Get("/song/findByFreeSearch", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			searchSong(server.db, w, r)
		}
	})

	//Albums
	r.Get("/album", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("album request")
		if authenticated(server.db, w, r) {
			getAllAlbum(server.db, w, r)
		}

	})
	r.Get("/album/findByArtist", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAlbumsByArtist(server.db, w, r)
		}
	})
	r.Get("/album/{albumID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAlbumsById(server.db, w, r)
		}
	})

	//Artists
	r.Get("/artist", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAllArtist(server.db, w, r)
		}
	})
	r.Get("/artist/{artistID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getArtistById(server.db, w, r)
		}
	})

	//Tags
	r.Get("/tag", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAllTag(server.db, w, r)
		}
	})

	r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			update(server.db, w, r, server.Token, server.coverFolder, server.musicFolder)
		}
	})

	r.Get("/download/{boxID}", func(w http.ResponseWriter, r *http.Request) {
		err := download(server.db, server.Token, w, r)
		if err != nil {
			server.GetBoxToken()
			err := download(server.db, server.Token, w, r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	//Playlist
	r.Post("/playlist", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			createPlaylist(server.db, w, r)
		}
	})
	r.Get("/playlist", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAllPlaylist(server.db, w, r)
		}
	})
	////TODO
	r.Get("/playlist/{playlistID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getPlaylistById(server.db, w, r)
		}
	})
	//r.Get("/playlist/{playlistID}/download", func(w http.ResponseWriter, r *http.Request) {
	//	if authenticated(server.db, w, r) {
	//		getPlaylistById(server.db, w, r)
	//	}
	//})
	r.Delete("/playlist/{playlistID}", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			deletePlaylistByID(server.db, w, r)
		}
	})

	//User
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		createUser(server.db, w, r)
	})
	r.Post("/user/{userID}/validate", func(w http.ResponseWriter, r *http.Request) {
		validateUser(server.db, w, r)
	})
	r.Post("/user/login", func(w http.ResponseWriter, r *http.Request) {
		loginUser(server.db, w, r)
	})
	return r
}

func authenticated(db *sqlx.DB, w http.ResponseWriter, r *http.Request) bool {
	sessionID := r.Header.Get("session")
	userID, _ := models.ValidateSessionID(db, sessionID)
	if userID <= 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
