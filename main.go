//package main
//
//import (
//	"fmt"
//	"github.com/annakertesz/cp-music-lib/transport"
//	"log"
//	"net/http"
//	"os"
//)
//
//func determineListenAddress() (string, error) {
//	port := os.Getenv("PORT")
//	if port == "" {
//		return "", fmt.Errorf("$PORT not set")
//	}
//	return ":" + port, nil
//}
//
//func main() {
//	addr, err := determineListenAddress()
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("Started")
//	if err := http.ListenAndServe(addr, transport.Routes()); err != nil {
//		log.Fatal("Could not start HTTP server", err.Error())
//	}
//}
//
//
package main

import (
	"database/sql"
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type e map[string]string

//type User struct {
//	Id       int64  `json:"id"`
//	Username string `json:"username"`
//}

func main() {
	var err error

	url, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatalln("$DATABASE_URL is required")
	}

	db, err = connect(url)

	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			switch r.Method {
			case "GET":
				users, err := models.Users(db)

				if err != nil {
					errorResponse(r, w, err)
				}

				respond(r, w, http.StatusOK, users)
			case "POST":
				user, err := models.UnmarshalUser(r)

				if err != nil {
					errorResponse(r, w, err)
					return
				}

				created, err := models.CreateUser(db, user.Username)

				if err != nil {
					errorResponse(r, w, err)
					return
				}

				respond(r, w, http.StatusOK, created)
			default:
				respond(r, w, http.StatusNotFound, nil)
			}
		} else {
			respond(r, w, http.StatusNotFound, nil)
		}
	})

	log.Printf("Starting server on port %s\n", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}

func connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id       SERIAL,
      username VARCHAR(64) NOT NULL UNIQUE,
      CHECK (CHAR_LENGTH(TRIM(username)) > 0)
    );
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func respond(r *http.Request, w http.ResponseWriter, status int, data interface{}) {
	if data != nil {
		bytes, err := json.Marshal(data)

		if err != nil {
			errorResponse(r, w, err)
			return
		}

		response(r, w, status, bytes)
	} else {
		response(r, w, status, nil)
	}
}

func errorResponse(r *http.Request, w http.ResponseWriter, err error) {
	bytes, _ := json.Marshal(e{"message": err.Error()})

	response(r, w, http.StatusInternalServerError, bytes)
}

func response(r *http.Request, w http.ResponseWriter, status int, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if bytes != nil {
		bytes = append(bytes, '\n')
		w.Write(bytes)
	}

	log.Printf(
		"\"%s %s %s\" %d %d\n",
		r.Method, r.URL.Path, r.Proto, status, len(bytes),
	)
}
