// Declare that this is the main package
package main

import (
	"fmt"
	"os"

	// import custom packages
	"gin-test/stationary"

	// Gin provides the framework for building web applications
	"github.com/gin-gonic/gin"

	// import dotenv for environment variables
	"github.com/joho/godotenv"
)

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// group of static routes looks like this: localhost:8080/static/...
	static := router.Group("/static")
	{
		// Endpoint to get all albums
		static.GET("/albums", stationary.GetAlbums)

		static.GET("/albums/:id", stationary.GetAlbumById)

		static.GET("/test", stationary.TestFunction)
	}

	serverLocation := makeServerLocation()

	// Serve the application
	router.Run(serverLocation)
}

// Helper function to read environment variables
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

// using environment variables, consume the variables
func makeServerLocation() string {
	location := goDotEnvVariable("ROOT_URL")
	port := goDotEnvVariable("PORT")

	return fmt.Sprintf("%s:%s", location, port)
}
