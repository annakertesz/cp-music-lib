

package transport

import (
	"fmt"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

type HTTP struct {
	logger *zap.Logger

}

func NewHTTP(logger *zap.Logger)HTTP{
	return HTTP{logger}
}

func (h *HTTP) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "cp-audio-lib")
	})

	/*
GET /songs
GET	/artists
GET	/albums
GET	/{album}/songs
GET	/{artist}/albums
GET	/songs/{songID}
POST/songs/{songID}
DEL /songs/{songID}
GET	/albums/{albumID}
POST/albums/{albumID}
DEL	/albums/{albumID}
GET	/artist/{artistID}
POST/artist/{artistID}
DEL	/artist/{artistID}



*/

	return r
}
