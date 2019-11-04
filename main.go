package main

import (
	"mailserver/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mailserver/server"
	"mailserver/libraries/middlewares"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	gin.SetMode(gin.DebugMode)
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}