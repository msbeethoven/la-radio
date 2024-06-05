package models

import (
	"time"
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
