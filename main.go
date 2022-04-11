// Declare that this is the main package
package main

import (
	"github.com/gin-gonic/gin"

	"gin-test/controler"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Endpoint to get all albums
	router.GET("/albums", controler.GetAlbums)

	router.GET("/albums/:id", controler.GetAlbumById)

	router.GET("/test", controler.TestFunction)

	// Serve the application
	router.Run("localhost:8080")
}
