package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	//	"timeconversion"
	"fmt"
	"radio/blah"
)

type Song struct {
	SongId        int           `json:"song_id"`
	Title         string        `json:"title"`
	Artist        string        `json:"artist"`
	Year          int           `json:"year"`
	Album         string        `json:"album"`
	Genre         string        `json:"genre"`
	Duration      time.Duration `json:"duration"`
	AudioPathFile string        `json:"audio_file_path"`
}

var db *sql.DB

func main() {
	fmt.Println(blah.Okay) //you had this package in the blah.go file called "bleh" first and it needed to be bleh.Okay

	var err error
	// db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/radio?sslmode=disable")
	db, err = sql.Open("postgres", "host=localhost port=5431 user=postgres dbname=radio sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	//CORS Exceptions
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // You can set this to the origins you want to allow, or use "*" to allow all origins.
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// Router requests
	router.GET("/songs", getSongs)
	//router.POST("/songs", createTrack)

	// Getting the songid to stream
	router.GET("/stream/id/:id", streamSongByID)

	router.Run("localhost:8081")
}

func getSongs(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	//Query database
	rows, err := db.Query("SELECT song_id, title, artist, album, year , audio_file_path FROM songs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var s Song
		//var durationString string

		err := rows.Scan(&s.SongId, &s.Title, &s.Artist, &s.Album, &s.Year, &s.AudioPathFile)
		if err != nil {
			log.Fatal(err)
		}
		// s.Duration, err = handlers.parseDurationFromHHMMSS(durationString)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		songs = append(songs, s)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, songs)
}

func streamSongByID(c *gin.Context) {
	id := c.Param("id")
	var audioPath string

	err := db.QueryRow("SELECT audio_file_path FROM songs WHERE song_id = $1", id).Scan(&audioPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.File(audioPath) // This streams the file to the browser
}
