package model

// Track represents the model for storing track metadata
type Track struct {
	ID           int    //`gorm:"primaryKey"`
	ISRC         string //`gorm:"uniqueIndex"`
	SpotifyImage string
	Title        string
	ArtistNames  string
	Popularity   int
}

type TrackArtist struct {
	ID          int
	Name        string
	URI         string
	Href        string
	AlbumHref   string
	AlbumName   string
	AlbumType   string
	AlbumURI    string
	AlbumID     string
	TotalTrack  int
	ReleaseDate string
}

type TrackRes struct {
	Tracks Response `json:"tracks,omitempty"`
}

type Response struct {
	Href     string `json:"href,omitempty"`
	Items    []Item `json:"items,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Next     *int   `json:"next,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Previous *int   `json:"previous,omitempty"`
	Total    int    `json:"total,omitempty"`
}

type Item struct {
	Album            Album        `json:"album,omitempty"`
	Artists          []Artist     `json:"artists,omitempty"`
	AvailableMarkets []string     `json:"available_markets,omitempty"`
	DiscNumber       int          `json:"disc_number,omitempty"`
	DurationMS       int          `json:"duration_ms,omitempty"`
	Explicit         bool         `json:"explicit,omitempty"`
	ExternalIds      ExternalID   `json:"external_ids,omitempty"`
	ExternalURLs     ExternalURLs `json:"external_urls,omitempty"`
	Href             string       `json:"href,omitempty"`
	ID               string       `json:"id,omitempty"`
	IsLocal          bool         `json:"is_local,omitempty"`
	Name             string       `json:"name,omitempty"`
	Popularity       int          `json:"popularity,omitempty"`
	PreviewURL       string       `json:"preview_url,omitempty"`
	TrackNumber      int          `json:"track_number,omitempty"`
	Type             string       `json:"type,omitempty"`
	URI              string       `json:"uri,omitempty"`
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls,omitempty"`
	Href         string       `json:"href,omitempty"`
	ID           string       `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Type         string       `json:"type,omitempty"`
	URI          string       `json:"uri,omitempty"`
}

type ExternalID struct {
	Isrc string `json:"isrc,omitempty"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify,omitempty"`
}

type Album struct {
	Artist               []Artist     `json:"artist,omitempty"`
	AlbumType            string       `json:"album_type,omitempty"`
	AvailableMarkets     []string     `json:"available_markets,omitempty"`
	ExternalURLs         ExternalURLs `json:"external_urls,omitempty"`
	Href                 string       `json:"href,omitempty"`
	ID                   string       `json:"id,omitempty"`
	Images               []Image      `json:"images,omitempty"`
	Name                 string       `json:"name,omitempty"`
	ReleaseDate          string       `json:"release_date,omitempty"`
	ReleaseDatePrecision string       `json:"release_date_precision,omitempty"`
	TotalTracks          int          `json:"total_tracks,omitempty"`
	Type                 string       `json:"type,omitempty"`
	URI                  string       `json:"uri,omitempty"`
}

type Image struct {
	Height int    `json:"height,omitempty"`
	URI    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
}

// TokenResponse holds the Spotify API access token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// ClientCredentials holds the client ID and client secret
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}
