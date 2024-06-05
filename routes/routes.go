package routes

import (
	"radio/api"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	songAPI := r.Group("/api/songs")
	{
		//songAPI.POST("/upload", handlers.UploadSong)
		songAPI.GET("/:song_id", handlers.GetSongMetadata)
	}

	return r
}
