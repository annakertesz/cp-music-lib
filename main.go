package main

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/library"
	"log"
	"net/http"
	"os"
)
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	x := library.Function(5,6)
	fmt.Fprintf(w, "library.Function(4,7) %v", x)
}
func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", hello)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}