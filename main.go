package main

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/transport"
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

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started")
	if err := http.ListenAndServe(addr, transport.Routes()); err != nil {
		log.Fatal("Could not start HTTP server", err.Error())
	}
}


