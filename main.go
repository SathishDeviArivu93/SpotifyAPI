package main

import (
	"fmt"
	"log"
	"net/http"
	"spotify/spotify/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/tracks/{isrc}", handler.CreateTrack).Methods("POST")
	r.HandleFunc("/tracks/{isrc}", handler.GetTrackByISRC).Methods("GET")
	r.HandleFunc("/tracks/artist/{artist}", handler.GetTracksByArtist).Methods("GET")

	// Start server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
