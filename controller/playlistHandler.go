package controller

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func createPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	playlist, err := models.UnmarshalPlaylist(r)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	id, err := models.CreatePlaylist(db, userID, playlist.Title)
	if err != nil || id == 0 {
		http.Error(w, err.Error(), 422)
		return
	}
	fmt.Fprint(w, id)
	w.WriteHeader(http.StatusOK)
}

func deletePlaylistByID(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "playlistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nplaylist id %v isnt a number", param)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.DeletePlaylist(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func addSongToPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "songID")
	songID, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number", param)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	plID := r.URL.Query().Get("playlistID")
	playlistID, err := strconv.Atoi(plID)
	if err != nil {
		fmt.Printf("playlist id %v isnt a number", plID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.AddSongToPlayist(db, songID, playlistID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func removeSongFromPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "songID")
	songID, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number", param)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	plID := r.URL.Query().Get("playlistID")
	playlistID, err := strconv.Atoi(plID)
	if err != nil {
		fmt.Printf("playlist id %v isnt a number", plID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = models.RemoveSongFromPlayist(db, songID, playlistID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getAllPlaylist(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	userID, err := models.ValidateSessionID(db, r.Header.Get("session"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	playlists, err := models.GetAllPLaylist(db, userID)
	playlistROs, err := playlistROListFromPlaylists(playlists, db)
	playlistJSON, err := json.Marshal(playlistROs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlistJSON)
	w.WriteHeader(http.StatusOK)
}

func getPlaylistById(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "playlistID")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("\nsong id %v isnt a number", param)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	playlist, err := models.GetPlaylistByID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	playlistRO, err := playlistROFromPlaylist(playlist, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	playlistJSON, err := json.Marshal(playlistRO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlistJSON)
	w.WriteHeader(http.StatusOK)
}

func playlistROFromPlaylist(playlist models.Playlist, db *sqlx.DB) (*models.PlaylistRO, error) {
	user, err := models.GetUserByID(db, playlist.User)
	if err != nil {
		return nil, err
	}
	songs, err := models.GetSongsByPlaylist(db, playlist.ID)
	if err != nil {
		return nil, err
	}
	songROList, err := songROListFromSongs(songs, db)
	if err != nil {
		return nil, err
	}
	return &models.PlaylistRO{
		ID:    playlist.ID,
		Title: playlist.Title,
		User:  UserROFromUser(user),
		Songs: songROList,
	}, nil
}

func playlistROListFromPlaylists(playlists []models.Playlist, db *sqlx.DB) ([]models.PlaylistRO, error) {
	playlistROs := make([]models.PlaylistRO, 0)
	for _, playlist := range playlists {
		pl, err := playlistROFromPlaylist(playlist, db)
		if err != nil {
			return nil, err
		}
		playlistROs = append(playlistROs, *pl)
	}
	return playlistROs, nil
}

func createPlaylistZip(db *sqlx.DB, w http.ResponseWriter, r *http.Request, token string, partially bool) error {

		param := chi.URLParam(r, "playlistID")
		id, err := strconv.Atoi(param)
		if err != nil {
			fmt.Printf("\nsong id %v isnt a number", param)
			w.WriteHeader(http.StatusBadRequest)
			return err
		}
		playlist, err := models.GetPlaylistByID(db, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		playlistRO, err := playlistROFromPlaylist(playlist, db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		songs := playlistRO.Songs
	// Create a buffer to write our archive to.
		buf := new(bytes.Buffer)

		// Create a new zip archive.
		zw := zip.NewWriter(buf)

		// Add some files to the archive.
		for _, song := range songs {
			f, err := zw.Create(toAlphaNumeric(song.Title) + ".mp3")
			if err != nil {
				log.Fatal(err)
			}
			if song.LqSong != nil {
				id, err := strconv.Atoi(*song.LqSong)
				resp, _, errModel := box_lib.DownloadFileBytes(token, id)
				if errModel != nil {
					if partially {
						services.HandleError(db, *errModel)
						log.Printf("ERROR: failed download song for playlist - playlistID: %v, songID: %v", id, song.ID)
						continue
					} else {
						return errModel.Err
					}
				}
				_, err = f.Write([]byte(resp))
				if err != nil {
					if partially {
						services.HandleError(db, *errModel)
						log.Printf("ERROR: failed download song for playlist - playlistID: %v, songID: %v", id, song.ID)
						continue
					} else {
						return errModel.Err
					}
				}
			}
			if song.LqInstr != nil {
				id, err := strconv.Atoi(*song.LqInstr)
				resp, _, errModel := box_lib.DownloadFileBytes(token, id)
				if errModel != nil {
					if partially {
						services.HandleError(db, *errModel)
						log.Printf("ERROR: failed download song for playlist - playlistID: %v, songID: %v", id, song.ID)
						continue
					} else {
						return errModel.Err
					}
				}
				_, err = f.Write([]byte(resp))
				if err != nil {
					if partially {
						services.HandleError(db, *errModel)
						log.Printf("ERROR: failed download song for playlist - playlistID: %v, songID: %v", id, song.ID)
						continue
					} else {
						return errModel.Err
					}
				}
			}
		}

		// Make sure to check the error on Close.
		err = zw.Close()
		if err != nil {
			log.Fatal(err)
		}
		w.Write(buf.Bytes())
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment")
		return nil
}

func toAlphaNumeric(s string) string{
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")

	return processedString
}
