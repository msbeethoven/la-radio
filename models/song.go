package models

import (
	"fmt"
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

func NewSong(songID int, title, artist string, year int, album, genre, audioPathFile string, duration time.Duration) *Song {
	if year < 0 {
		year = 0
	}
	

	return &Song{
		SongId:        songID,
		Title:         title,
		Artist:        artist,
		Year:          year,
		Album:         album,
		Genre:         genre,
		Duration:      time.Duration,
		AudioPathFile: audioPathFile,
	}
}




type Playlist struct {
    Name String
	CurrentSong Song
	LastPlayed Song
	SongList   []*Song
}

func NewPlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}


func (p *Playlist) AddSongToPlaylist(song *Song) {
	p.SongList = append(p.SongList, song)
}


func (p *Playlist) UpdateCurrentSong(song *Song) {
	p.LastPlayed = CurrentSong
	p.CurrentSong = song
}

