package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
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
	AudioPathFile string        `json:"audio_path_file"`
}

var db *sql.DB

func main() {
	fmt.Println(blah.Okay) //you had this package in the blah.go file called "bleh" first and it needed to be bleh.Okay

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/radio?sslmode=disable")
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

	router.Run("localhost:8080")
}

func getSongs(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	//Query database
	rows, err := db.Query("SELECT song_id, title, artist, album, year FROM songs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var s Song
		//var durationString string

		err := rows.Scan(&s.SongId, &s.Title, &s.Artist, &s.Album, &s.Year)
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

// import (
// 	//"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"
// 	//"github.com/jinzhu/gorm/dialects/postgres"
// 	_ "github.com/lib/pq"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func GetAllSongs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	db, err := gorm.Open("postgres", "host=localhost port=8080 user=your-username dbname=radio sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	var songs []Song
// 	db.Find(&songs) //finding all, everything in db, hence plural

// 	for _, song := range songs {
// 		fmt.Printf("SongId: %d, Title: %s, Artist: %.2f\n", song.SongId, song.Title, song.Artist)
// 	}
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/songs", GetAllSongs)

// 	// router.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
// 	// 	json.NewEncoder(w).Encode("Hello songs")
// 	// })

// 	log.Println("Radio is starting!")
// 	http.ListenAndServe(":8080", router)
// }

// type Song struct {
// 	SongId        int           `json:"song_id"`
// 	Title         string        `json:"title"`
// 	Artist        string        `json:"artist"`
// 	Year          int           `json:"year"`
// 	Album         string        `json:"album"`
// 	Genre         string        `json:"genre"`
// 	Duration      time.Duration `json:"duration"`
// 	AudioPathFile string        `json:"audio_path_file"`
// }

// func main() {
// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})
// }

// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, World!"))
// 	})

// 	http.ListenAndServe(":8080", nil)

// 	r := routes.SetupRoutes()
// 	r.Run(":8080")
// }
