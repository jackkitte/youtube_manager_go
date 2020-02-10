package main

import (
	"github.com/jackkitte/youtube_manager_go/middlewares"
	"github.com/jackkitte/youtube_manager_go/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	heroku := os.Getenv("HEROKU")
	if heroku != "true" {
		err := godotenv.Load("env/.env")

		if err != nil {
			logrus.Fatal("[Echo] Error loading .env")
		}
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	port := os.Getenv("PORT")
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YouTubeService())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())

	//Routes
	routes.Init(e)

	e.Logger.Fatal(e.Start(":" + port))
}
