package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/config"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/annakertesz/cp-music-lib/services/updater"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Server struct {
	db           *sqlx.DB
	musicFolder  int
	coverFolder  int
	defaultPicture int
	BoxConfig    config.BoxConfig
	EmailSender  services.EmailSender
}

type e map[string]string

func NewServer(db *sqlx.DB, cfg config.Config) *Server {
	token := box_lib.AuthOfBox(cfg.BoxConfig.ClientID, cfg.BoxConfig.ClientSecret, cfg.BoxConfig.PrivateKey)
	boxCfg := cfg.BoxConfig
	boxCfg.Token = token
	emailSender := services.NewEmailSender(cfg.SengridConfig)
	return &Server{
		db:           db,
		musicFolder:  cfg.SongFolder,
		coverFolder:  cfg.CoverFolder,
		defaultPicture: cfg.DefaultPicture,
		EmailSender:  emailSender,
		BoxConfig:    boxCfg,
	}
}

func (s *Server) GetBoxToken() {
	s.BoxConfig.Token = box_lib.AuthOfBox(s.BoxConfig.ClientID, s.BoxConfig.ClientSecret, s.BoxConfig.PrivateKey)
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
		fmt.Fprint(w, "cp")
	})

	//Songs
	r.Get("/song", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			getAllSongs(server.db, w, r)
		}
	})
	r.Get("/song/{songID}", func(w http.ResponseWriter, r *http.Request) {
		//if authenticated(server.db, w, r) {
			getSongByID(server.db, w, r)
		//}
	})
	r.Post("/song/buy", func(w http.ResponseWriter, r *http.Request) {
		if authenticated(server.db, w, r) {
			buySong(server.db, server.EmailSender, w, r)
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
		//if authenticated(server.db, w, r) {
			searchSong(server.db, w, r)
		//}
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
			getAllTag(server.db, w, r)
	})

	r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
			update(server.db, w, r, server.BoxConfig.Token, server.coverFolder, server.musicFolder)
	})

	r.Get("/cleardb", func(w http.ResponseWriter, r *http.Request) {
		cleardb(server.db, w, r)
	})

	r.Get("/download/", func(w http.ResponseWriter, r *http.Request) {
		err := download(server.db, server.BoxConfig.Token, server.defaultPicture, w, r)
		if err != nil {
			server.GetBoxToken()
			err := download(server.db, server.BoxConfig.Token,server.defaultPicture, w, r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	r.Get("/download/{boxID}", func(w http.ResponseWriter, r *http.Request) {
		err := download(server.db, server.BoxConfig.Token, server.defaultPicture, w, r)
		if err != nil {
			server.GetBoxToken()
			err := download(server.db, server.BoxConfig.Token, server.defaultPicture, w, r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	//Playlist
	r.Post("/playlist", func(w http.ResponseWriter, r *http.Request) {
		//if authenticated(server.db, w, r) {
			createPlaylist(server.db, w, r)
		//}
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
		createUser(server.db, server.EmailSender, w, r)
	})
	r.Get("/user/{userID}/validate/{token}", func(w http.ResponseWriter, r *http.Request) {
		validateUser(server.db, w, r)
	})
	r.Get("/updateBdgr83rgsdf", func(w http.ResponseWriter, r *http.Request) {
		server.GetBoxToken()
		go updater.Update(server.musicFolder, server.coverFolder, server.BoxConfig.Token, server.db)
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
