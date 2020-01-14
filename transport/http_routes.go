package transport

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "cp-audio-lib")
	})

	return r
}