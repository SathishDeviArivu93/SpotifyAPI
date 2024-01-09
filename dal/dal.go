package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"spotify/spotify/model"

	_ "github.com/go-sql-driver/mysql"
)

// Initialize DB and Migrate
func InitDB() (*sql.DB, error) {
	dsn := "root:pass123@tcp(127.0.0.1:3306)/spotifyschema"

	// Connect to MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("err", err)
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func InsertionData(trk *model.Track, trkart *model.TrackArtist) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println("InitDB error", err)
		return errors.New(err.Error())
	}
	defer db.Close()

	trackData, err := FetchDataByISRC(trk.ISRC)
	if err != nil {
		return errors.New(err.Error())
	}
	if trackData.ISRC != "" {
		return errors.New("Already record created with this ISRC")
	}

	stmt, err := db.Prepare("INSERT INTO track (ISRC, SpotifyImage, Title, ArtistName, Popularity) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New(err.Error())
	}

	defer stmt.Close()

	// Execute the SQL statement with struct data
	rslt, err := stmt.Exec(trk.ISRC, trk.SpotifyImage, trk.Title, trk.ArtistNames, trk.Popularity)
	if err != nil {
		fmt.Println("err in running exec", err)
		return err
	}

	lastid, _ := rslt.LastInsertId()
	fmt.Println("lastid", lastid)

	stmt1, err1 := db.Prepare("INSERT INTO trackartist (ID, Name, URI, Href, AlbumHref, AlbumName, AlbumType, AlbumURI, AlbumID, TotalTrack, ReleaseDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err1 != nil {
		fmt.Println("err in running stmt1", err1)
	}

	defer stmt1.Close()

	// Execute the SQL statement with struct data
	_, err = stmt1.Exec(int(lastid), trkart.Name, trkart.URI, trkart.Href, trkart.AlbumHref, trkart.AlbumName, trkart.AlbumType, trkart.AlbumURI, trkart.AlbumID, trkart.TotalTrack, trkart.ReleaseDate)
	if err != nil {
		fmt.Println("err in running exec1", err)
		return err
	}
	return nil
}

func FetchDataByISRC(isrc string) (model.Track, error) {

	query := "SELECT * FROM track WHERE isrc = ?"

	db, _ := InitDB()
	var trk model.Track

	// Query the database
	rows, err := db.Query(query, isrc)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		if err := rows.Scan(&trk.ID, &trk.ISRC, &trk.SpotifyImage, &trk.Title, &trk.ArtistNames, &trk.Popularity); err != nil {
			log.Fatal(err)
		}

		// Process the retrieved data, e.g., print it
		fmt.Printf("ISRC: %s\n", trk.ISRC)
		// Add other fields as needed
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return trk, err
}

func FetchDataByArtist(artistname string) ([]model.Track, error) {
	query := "SELECT * FROM track WHERE ArtistName = ?"

	db, _ := InitDB()
	var trks []model.Track

	// Query the database
	rows, err := db.Query(query, artistname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var trk model.Track
		if err := rows.Scan(&trk.ID, &trk.ISRC, &trk.SpotifyImage, &trk.Title, &trk.ArtistNames, &trk.Popularity); err != nil {
			return nil, err
		}
		trks = append(trks, trk)
	}
	return trks, err
}
