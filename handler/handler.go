package handler

import (
	"encoding/json"
	"net/http"
	"spotify/spotify/service"

	"github.com/gorilla/mux"
)

// Handler to create a track
func CreateTrack(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isrc := params["isrc"]

	// Fetch metadata from Spotify API
	err := service.FetchMetadataFromSpotify(isrc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()}) //"Failed to fetch metadata from Spotify API"})
		return
	}
	// fmt.Println("metadata", metadata)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Successfully created")
}

func GetTrackByISRC(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isrc := params["isrc"]

	// Fetch metadata from Spotify API
	trackData, err := service.FetchDataByISRC(isrc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackData)
}

func GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	artist := params["artist"]

	// Fetch metadata from Spotify API
	trackData, err := service.FetchDataByArtist(artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackData)
}
