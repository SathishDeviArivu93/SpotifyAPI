package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"spotify/spotify/dal"
	"spotify/spotify/model"
	"spotify/spotify/utils"
	"strings"
)

func FindMaxValue(itm []model.Item, propertyName string) (item model.Item) {
	if len(itm) == 0 {
		return model.Item{}
	}
	for i := 0; i < len(itm); i++ {
		max := itm[i].Popularity
		for _, obj := range itm {
			switch propertyName {
			case "Popularity":
				if obj.Popularity > max {
					max = obj.Popularity
					item = obj
				}
			}
		}
	}
	return
}

// GetAccessToken retrieves the access token using the Client Credentials Flow
func GetAccessToken(credentials model.ClientCredentials) (string, error) {
	// Form data
	body := strings.NewReader("grant_type=client_credentials&client_id=95cfd3f92b0c416491ad68741884cd22&client_secret=51c919eff66d46e2b0f53d4dd299a43a")
	req, err := http.NewRequest("POST", utils.SpotifyTokenURL, body)
	if err != nil {
		return "", err
	}

	// Set content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response
	var tokenResponse model.TokenResponse
	if err := parseJSONResponse(resp, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

// Helper function to parse JSON response
func parseJSONResponse(resp *http.Response, target interface{}) error {
	return json.NewDecoder(resp.Body).Decode(target)
}

// Function to fetch metadata from Spotify API
func FetchMetadataFromSpotify(isrc string) error {

	var respon model.TrackRes
	var trk model.Track
	var trkart model.TrackArtist
	var spcItem model.Item

	// Replace these with your own client credentials
	credentials := model.ClientCredentials{
		ClientID:     "95cfd3f92b0c416491ad68741884cd22",
		ClientSecret: "51c919eff66d46e2b0f53d4dd299a43a",
	}

	// Obtain the access token
	accessToken, err := GetAccessToken(credentials)
	if err != nil {
		fmt.Println("Error obtaining access token:", err)
		return errors.New(err.Error())
	}
	// Construct the Spotify API request URL for track search by ISRC
	spotifyURL := fmt.Sprintf("%s?q=isrc:%s&type=track", utils.SpotifySearchEndpoint, isrc)

	req, err := http.NewRequest("GET", spotifyURL, nil)
	if err != nil {
		return errors.New(err.Error())
	}
	// Set the authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respon)
	if err != nil {
		fmt.Println("Error while decode", err)
		return errors.New(err.Error())
	}
	err = json.NewDecoder(resp.Body).Decode(respon)
	trk.ISRC = isrc

	// Finding the ISRC which has more popularity
	if len(respon.Tracks.Items) > 0 {
		spcItem = FindMaxValue(respon.Tracks.Items, "Popularity")
	}
	trk.Title = spcItem.Name
	if len(spcItem.Album.Images) > 0 {
		trk.SpotifyImage = spcItem.Album.Images[0].URI
	}

	// for i := 0; i < len(spcItem.Artists); i++ {
	// 	trk.ArtistNames = append(trk.ArtistNames, spcItem.Artists[i].Name)
	// }
	trk.ArtistNames = spcItem.Artists[0].Name
	trk.Popularity = spcItem.Popularity

	trkart.Name = spcItem.Name
	trkart.URI = spcItem.URI
	trkart.Href = spcItem.Href
	trkart.AlbumHref = spcItem.Album.Href
	trkart.AlbumName = spcItem.Album.Name
	trkart.AlbumType = spcItem.Album.Type
	trkart.AlbumURI = spcItem.Album.URI
	trkart.AlbumID = spcItem.Album.ID
	trkart.TotalTrack = spcItem.Album.TotalTracks
	trkart.ReleaseDate = spcItem.Album.ReleaseDate

	fmt.Println("Track", trk)
	fmt.Println("TrackArtist", trkart)

	err = dal.InsertionData(&trk, &trkart)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func FetchDataByISRC(isrc string) (model.Track, error) {
	return dal.FetchDataByISRC(isrc)
}

func FetchDataByArtist(artist string) ([]model.Track, error) {
	return dal.FetchDataByArtist(artist)
}
