package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kadekchrisna/nothoax-blog-api/routes"
)

// Serve for Serving App
func Serve() {
	gin.SetMode(gin.DebugMode)

	err := godotenv.Load()
	if err != nil {
		panic("Error when loading env.")
	}
	router := routes.Routes()
	router.Run(os.Getenv("PORT"))
}
